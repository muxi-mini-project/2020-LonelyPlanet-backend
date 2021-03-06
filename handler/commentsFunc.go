package handler

import (
	"fmt"
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CommentCreate(c *gin.Context) {
	uid := c.GetString("uid")

	black, err := model.ConfirmBlackList(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if black {
		c.JSON(401, gin.H{
			"message": "此用户在黑名单中",
		})
		return
	}

	var data model.Night_comment
	secretid, _ := strconv.Atoi(c.Query("secretId"))
	receiverid := c.Query("receiver_sid")
	if !model.CheckDebunk(secretid) {
		c.JSON(400, gin.H{
			"message": "该秘密不存在",
		})
		return
	}
	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	data.ReceiverSid = receiverid
	data.SecretId = secretid
	data.CommentTime = model.NowTime()
	fmt.Println(data)
	err = model.CreateComment(data)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "发布失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "发布成功",
	})
}

func CommentHistory(c *gin.Context) {
	secretid, _ := strconv.Atoi(c.Query("secretId"))
	page, _ := strconv.Atoi(c.Query("page"))
	limit := 5
	history, err := model.HistoryComment(secretid, page, limit)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "请求失败",
		})
		return
	}
	comment := model.GetCommentHistory(history)
	c.JSON(200, gin.H{
		"message": "请求成功",
		"history": comment,
	})
}

func CommentDelete(c *gin.Context) {
	commentid, _ := strconv.Atoi(c.Query("commentId"))
	if !model.CheckComment(commentid) {
		c.JSON(400, gin.H{
			"message": "该评论不存在",
		})
		return
	}
	err := model.DeleteComment(commentid)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}
