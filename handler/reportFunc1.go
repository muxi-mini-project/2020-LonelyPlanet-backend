package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	"github.com/2020-LonelyPlanet-backend/miniProject/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func NightSecretReport(c *gin.Context) {
	uid := c.GetString("uid")
	secretid, _ := strconv.Atoi(c.Query("secretId"))

	var newReport model.ReportInformation
	if err := c.BindJSON(&newReport); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "Bad Request!",
		})
		return
	}

	secret, err := model.GetDebunk(secretid)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "举报失败",
		})
		return
	}

	err = util.SendMail(2, secret.Content, newReport.Reason, uid, secret.SenderSid, newReport.Addition)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "举报失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "举报成功",
	})
}

func NightNightReport(c *gin.Context) {
	uid := c.GetString("uid")
	commentid, _ := strconv.Atoi(c.Query("commentId"))

	var newReport model.ReportInformation
	if err := c.BindJSON(&newReport); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "Bad Request!",
		})
		return
	}

	comment, err := model.GetComment(commentid)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "举报失败",
		})
		return
	}

	a := "匿名"
	err = util.SendMail(2, comment.Comment, newReport.Reason, uid, a, newReport.Addition)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "举报失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "举报成功",
	})
}
