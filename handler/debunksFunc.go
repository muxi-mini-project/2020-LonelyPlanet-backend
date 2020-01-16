package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	"github.com/gin-gonic/gin"
	//"strconv"
)

func DebunksCreate(c *gin.Context) {
	uid := c.GetString("uid")
	var data model.DebunkInfo
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	debunk := model.Debunks{
		SenderSid: uid,
		Content:   data.Content,
		Colour:    data.Colour,
		SendTime:  data.SendTime,
	}
	secretid, err := model.CraeteDebunk(debunk)
	if err != nil {
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
	secretid := c.Query("secretId")
	err := model.DeleteDebunk(secretid)
	if err != nil {
		c.JSON(400,gin.H{
			"message":"删除失败",
		})
		return
	}
	c.JSON(200,gin.H{
		"message": "success",
	})
}

func DebunksHistory(c *gin.Context) {
	uid := c.GetString("uid")
	history,err := model.HistoryDebunk(uid)
	if err != nil {
		c.JSON(400,gin.H{
			"message":"请求失败",
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"请求成功",
		"history": history,
	})
	return
}