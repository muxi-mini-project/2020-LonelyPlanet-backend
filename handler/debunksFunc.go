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
	sendtime := model.NowTime()
	debunk := model.Debunk{
		//Debunkid:  0,
		SenderSid: uid,
		Content:   data.Content,
		Colour:    data.Colour,
		SendTime:  sendtime,
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

func GetSecret(c *gin.Context) {
	secretid, _ := strconv.Atoi(c.Query("secretId"))
	secret, err := model.GetDebunk(secretid)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "获取秘密失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "请求成功",
		"secret":  secret,
	})
}

func DebunksDelete(c *gin.Context) {
	var err1 error
	secretid, _ := strconv.Atoi(c.Query("secretId"))
	//fmt.Println(secretid)
	if !model.CheckDebunk(secretid) {
		c.JSON(400, gin.H{
			"message": "该秘密不存在",
		})
		return
	}
	err := model.DeleteDebunk(secretid)
	commentHistory, _ := model.HistoryComment1(secretid)
	for _, data := range commentHistory {
		err1 = model.DeleteComment(data.CommentId)
	}
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "删除秘密失败",
		})
		return
	}
	if err1 != nil {
		log.Println(err1)
		c.JSON(400, gin.H{
			"message": "删除秘密失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func DebunksHistory(c *gin.Context) {
	limit := 5
	page, _ := strconv.Atoi(c.Query("page"))
	uid := c.GetString("uid")
	history, err := model.HistoryDebunk(uid, page, limit)
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
	var secret model.Debunk
	limit := 1
	uid := c.GetString("uid")
	page, _ := strconv.Atoi(c.Query("page"))
	secret, err := model.SquareDebunk(uid, page, limit)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "请求错误",
		})
		return
	}
	i := model.RandNum(10)
	c.JSON(200, gin.H{
		"message": "请求成功",
		"secret":  secret,
		"number":  i,
	})
	return
}
