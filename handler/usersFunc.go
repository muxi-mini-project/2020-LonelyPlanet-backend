package handler

import (
	"fmt"
	"github.com/2020-LonelyPlanet-backend/miniProject/middleware"
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	error2 "github.com/2020-LonelyPlanet-backend/miniProject/pkg/error"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"log"
=======
	log "github.com/sirupsen/logrus"
>>>>>>> e29a1e1360672468b9761bcd5446b1855c1d3870
	"math/rand"
	"time"
)

type verifyInfo struct {
	VerifyItem string `json:"verify_item"`
	VerifyInfo string `json:"verify_info"`
}

func UserLogin(c *gin.Context) {
	//a, _ := c.Get("token2")
	var tmpUser model.SuInfo
	var tmpLoginInfo model.LoginInfo
	fmt.Println("121")
	if err := c.BindJSON(&tmpLoginInfo); err != nil {
		ErrBadRequest(c, error2.BadRequest)
		return
	}

	tmpUser, err := model.GetUserInfoFormOne(tmpLoginInfo.Sid, tmpLoginInfo.Pwd)
	if err != nil {
		ErrBadRequest(c, error2.LoginError)
		return
	}

	//fmt.Print(tmpUser)

	user := model.UserInfo{
		Sid:           tmpUser.User.Usernumber,
		NickName:      getRandomString(8),
		College:       tmpUser.User.DeptName,
		Gender:        getGender(tmpUser.User.Xb),
		Grade:         model.ChangeString(tmpUser.User.Usernumber, 1, 4),
		Portrait:      getRandomPortrait(),//rand
	}

	err = model.CreatUser(user)    //写入用户数据到数据库
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}

	c.Header("token", produceToken(user.Sid))
	c.JSON(200, gin.H{
		"msg": "success",
	})
	return
}

func Test(c *gin.Context) {
	//a,b := c.Get("uid")
	//a := c.Query("qq")
	//fmt.Println(len(a))
	//token := c.Request.Header.Get("token")
	uid := c.GetString("uid")
	fmt.Println(uid)
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
	Uid string	`json:"uid"`
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

func ChangeInfo(c *gin.Context) {
	uid := c.GetString("uid")
	var tmpInfo verifyInfo
	if err := c.BindJSON(&tmpInfo); err != nil {
		log.Println(err)
		ErrBadRequest(c, error2.BadRequest)
		return
	}
	err := model.VerifyInfo(uid, tmpInfo.VerifyItem, tmpInfo.VerifyInfo)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg": "Success",
	})
<<<<<<< HEAD
	return
=======
>>>>>>> e29a1e1360672468b9761bcd5446b1855c1d3870
}
