package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func DebunksCreate(c *gin.Context) {
	uid := c.GetString("uid")
	var data model.DebunkInfo
	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	debunk := model.Debunk{
		SenderSid: uid,
		Content:   data.Content,
		Colour:    data.Colour,
		SendTime:  data.SendTime,
	}
	secretid, err := model.CraeteDebunk(debunk)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "Success",
		"secretId": secretid,
	})
}

func DebunksDelete(c *gin.Context) {
	secretid, _ := strconv.Atoi(c.Query("secretId"))
	err := model.DeleteDebunk(secretid)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func DebunksHistory(c *gin.Context) {
	uid := c.GetString("uid")
	history, err := model.HistoryDebunk(uid)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "请求成功",
		"history": history,
	})
	return
}
func DebunksSquare(c *gin.Context) {
	limit := 1
	var secret []model.Debunk
	page, _ := strconv.Atoi(c.Query("page"))
	secret, err := model.SquareDebunk(page, limit)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "请求错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "请求成功",
		"secrets": secret,
	})
	return
}
