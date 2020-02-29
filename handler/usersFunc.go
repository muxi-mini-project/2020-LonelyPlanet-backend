package handler

import (
	"fmt"
	//"github.com/2020-LonelyPlanet-backend/miniProject/middleware"
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	error2 "github.com/2020-LonelyPlanet-backend/miniProject/pkg/error"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
)

type verifyInfo struct {
	VerifyItem string `json:"verify_item"`
	VerifyInfo string `json:"verify_info"`
}

// @Summary 登录
// @Description Login
// @Tags user
// @Accept json
// @Produce json
// @Param loginInfo body model.LoginInfo true "学号和密码"
// @Success 200 {object} model.Res "{"msg":"success", "token": string}"
// @Failure 401 {object} error.Error "{"error_code":"20001", "message":"Password or account wrong."} 登录失败, {"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /login/ [post]
func UserLogin(c *gin.Context) {
	//a, _ := c.Get("token2")
	var tmpUser model.SuInfo
	var tmpLoginInfo model.LoginInfo
	if err := c.BindJSON(&tmpLoginInfo); err != nil {
		ErrBadRequest(c, error2.BadRequest)
		return
	}

	tmpUser, err := model.GetUserInfoFormOne(tmpLoginInfo.Sid, tmpLoginInfo.Pwd)
	if err != nil {
		ErrUnauthorized(c, error2.LoginError)
		return
	}

	//fmt.Println(tmpUser)

	user := model.UserInfo{
		Sid:      tmpUser.User.Usernumber,
		NickName: getRandomString(8),
		College:  tmpUser.User.DeptName,
		Gender:   getGender(tmpUser.User.Xb),
		Grade:    model.ChangeString(tmpUser.User.Usernumber, 1, 4),
		Portrait: getRandomPortrait(), //rand
	}

	err = model.CreatUser(user) //写入用户数据到数据库
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}

	//c.SetCookie("token", produceToken(user.Sid), 3600, "/", "mini", true, false)
	//c.Header("token", produceToken(user.Sid))
	c.JSON(200, gin.H{
		"msg":   "success",
		"token": produceToken(user.Sid),
	})
	return
}

func Test(c *gin.Context) {
	//a,b := c.Get("uid")
	//a := c.Query("qq")
	//fmt.Println(len(a))
	//token := c.Request.Header.Get("token")
	//uid := c.GetString("uid")
	//fmt.Println(uid)
	c.SetCookie("1", "2", 3600, "/", "l", false, true)
	return
}

func getGender(n string) string {
	if n == "1" {
		return "男"
	} else {
		return "女"
	}
}

func getRandomPortrait() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9)
}

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func Homepage(c *gin.Context) {
	uid := c.GetString("uid")
	tmpUser, err := model.FindUser(uid)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg":      "Success",
		"sid":      tmpUser.Sid,
		"nickName": tmpUser.NickName,
		"college":  tmpUser.College,
		"gender":   tmpUser.Gender,
		"grade":    tmpUser.Grade,
		"portrait": tmpUser.Portrait,
	})
	return
}

type jwtClaims struct {
	jwt.StandardClaims
	Uid string `json:"uid"`
}

var (
	key        = "miniProject" //salt
	ExpireTime = 3600          //token expire time
)

func produceToken(uid string) string {
	//id, _ := strconv.Atoi(uid)
	claims := &jwtClaims{
		Uid: uid,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	//fmt.Println(singedToken, err)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims jwtClaims) (string, error) {
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// @Summary 更改用户信息
// @Description VerifyItem传入修改的板块,当前只能更改昵称
// @Tags user
// @Accept json
// @Produce json
// @Param verifyInfo body handler.verifyInfo true "修改的板块和信息"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"msg":"success"} 成功"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/change_info/ [post]
func ChangeInfo(c *gin.Context) {
	uid := c.GetString("uid")
	var tmpInfo verifyInfo
	if err := c.BindJSON(&tmpInfo); err != nil {
		log.Println(err)
		ErrBadRequest(c, error2.BadRequest)
		return
	}
	if tmpInfo.VerifyItem != "Nickname" {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	err := model.VerifyInfo(uid, tmpInfo.VerifyItem, tmpInfo.VerifyInfo) //便于以后扩充
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg": "success",
	})
	return
}
