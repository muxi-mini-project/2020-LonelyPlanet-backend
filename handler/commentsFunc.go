package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CommentCreate(c *gin.Context) {
	var data model.Night_comment
	if err := c.BindJSON(&data) ; err != nil {
		log.Println(err)
		c.JSON(400,gin.H{
			"message" : "Bad Request!",
		})
		return
	}
	err := model.CreateComment(data)
	if err != nil {
		log.Println(err)
		c.JSON(400,gin.H{
			"message" : "发布失败",
		})
		return
	}
	c.JSON(200,gin.H{
		"message" : "发布成功",
	})
}

func CommentHistory(c *gin.Context) {
	secretid,_ := strconv.Atoi(c.Param("secret_id"))
	history,err := model.HistoryComment(secretid)
	if err != nil {
		log.Println(err)
		c.JSON(400,gin.H{
			"message" : "请求失败",
		})
		return
	}
	c.JSON(200,gin.H{
		"message" : "请求成功",
		"history" : history,
	})
}

func CommentDelete(c *gin.Context) {
	commentid,_ := strconv.Atoi(c.Query("comment_id"))
	err := model.DeleteComment(commentid)
	if err != nil {
		log.Println(err)
		c.JSON(400,gin.H{
			"message" : "删除失败",
		})
		return
	}
	c.JSON(200,gin.H{
		"message" : "删除成功",
	})
}


