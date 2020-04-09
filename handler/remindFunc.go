package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func RemindNightRemindboxView(c *gin.Context) {
	uid := c.GetString("uid")
	commentdata, err1 := model.GetCommentData(uid)
	if err1 != nil {
		log.Println(err1)
		c.JSON(400, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message":     "请求成功",
		"commentdata": commentdata,
	})
}

func NightRemindExistence(c *gin.Context) {
	uid := c.GetString("uid")
	status, err := model.CheckRemindExist(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "请求错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "请求成功",
		"status":  status,
	})
}

func UpdateNightRemindStatus(c *gin.Context) {
	commentid, _ := strconv.Atoi(c.Query("commentId"))
	status, err := model.CheckCommentIdExist(commentid)
	if status == 1 {
		c.JSON(400, gin.H{
			"message": "更新失败 该秘密不存在",
		})
	}
	err = model.ChangeStatus(commentid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "更新失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "更新成功",
	})
}
