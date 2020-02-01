package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	"github.com/gin-gonic/gin"
	"log"
)

func RemindNightRemindboxView(c *gin.Context){
	uid := c.GetString("uid")
	secretid,err := model.GetSecretid(uid)
	if err != nil {
		log.Println(err)
		c.JSON(400,gin.H{
			"message" : "请求错误",
		})
		return
	}
	commentdata,err1 := model.GetCommentData(secretid)
	if err1 != nil {
		log.Println(err1)
		c.JSON(400,gin.H{
			"message" : "请求失败",
		})
		return
	}
	c.JSON(200,gin.H{
		"message" : "请求成功",
		"commentdata" : commentdata,
	})
}