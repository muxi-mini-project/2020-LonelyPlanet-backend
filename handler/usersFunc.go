package handler

import (
	"fmt"
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	"github.com/gin-gonic/gin"
	"github.com/kocoler/GoExercises/miniProject/middleware"
	log "github.com/sirupsen/logrus"
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
	if err := c.BindJSON(&tmpLoginInfo); err != nil {
		c.JSON(400,gin.H{
			"message":"Bad Request!",
		})
		return
	}
	tmpUser, err := model.GetUserInfoFormOne(tmpLoginInfo.Sid, tmpLoginInfo.Pwd)
	if  err != nil {
		c.JSON(400,gin.H{
			"message":err,
		})
		return
	}
	fmt.Print(tmpUser)

	user := model.UserInfo{
		Sid:                   tmpUser.User.Usernumber,
		NickName:              getRandomString(8),
		College:               tmpUser.User.DeptName,
		Gender:                getGender(tmpUser.User.Xb),
		Grade:                 model.ChangeString(tmpUser.User.Usernumber,1,4),
		NightPortrait:         "",
		Requirements:          0,
		Debunks:               0,
	}

	err = model.CreatUser(user)
	if err != nil {
		c.JSON(400,gin.H{
			"message":err,
		})
		return
	}
	//写入用户数据到数据库
	c.Header("token",middleware.ProduceToken(user.Sid))
}

func Test(c *gin.Context) {
	a,b := c.Get("uid")
	fmt.Println(a,b)
}

func getGender(n string) string {
	if n == "1" {
		return "男"
	}else {
		return "女"
	}
}

func  getRandomString(l int) string {
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
		log.Println(err)
		c.JSON(400,gin.H{
			"message":err,
		})
		return
	}
	c.JSON(200,gin.H{
		"msg":"Success",
		"Sid":tmpUser.Sid,
		"NickName":tmpUser.NickName,
		"College":tmpUser.College,
		"Gender":tmpUser.Gender,
		"Grade":tmpUser.Grade,
	})
}

func ChangeInfo(c *gin.Context) {
	uid := c.GetString("uid")
	var tmpInfo verifyInfo
	if err := c.BindJSON(&tmpInfo); err != nil {
		log.Println(err)
		c.JSON(400,gin.H{
			"message":"Bad Request!",
		})
		return
	}
	err := model.VerifyInfo(uid,tmpInfo.VerifyItem,tmpInfo.VerifyInfo)
	if err != nil {
		log.Println(err)
		c.JSON(400,gin.H{
			"message":err,
		})
		return
	}
	c.JSON(200,gin.H{
		"msg":"Success",
	})
}