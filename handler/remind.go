package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	error2 "github.com/2020-LonelyPlanet-backend/miniProject/pkg/error"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DayInformationExistence(c *gin.Context) {
	uid := c.GetString("uid")
	result := model.ConfirmRemindExist(uid)
	c.JSON(200, gin.H{
		"msg":       "success",
		"existence": result,
	})
	return
}

/*
func NightInformationExistence(c *gin.Context) {
	uid := c.GetString("uid")
	result := model.ConfirmRemindExist(uid,2) //2＝黑夜
	c.JSON(200,gin.H{
		"msg":"success",
		"existence":result,
	})
	return
}
*/

func ReminderBox(c *gin.Context) {
	uid := c.GetString("uid")
	if len(c.Query("limit")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	if len(c.Query("page")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	var offset int
	offset = page* limit

	result, err := model.ReminderBox(uid, limit, offset)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg":     "success",
		"num":     len(result),
		"content": result,
	})
}

func UpdateRemindStatus1(c *gin.Context) {
	uid := c.GetString("uid")
	if len(c.Param("application_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	applicationId, _ := strconv.Atoi(c.Param("application_id"))
	err, status := model.ReminderChangeStatus(applicationId, uid, 2)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	if !status {
		ErrUnauthorized(c, error2.Unauthorized)
		return
	}
	c.JSON(200, gin.H{
		"msg": "success",
	})
	return
}

func UpdateRemindStatus2(c *gin.Context) {
	uid := c.GetString("uid")
	if len(c.Param("application_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	applicationId, _ := strconv.Atoi(c.Param("application_id"))
	err, status := model.ReminderChangeStatus(applicationId, uid, 1)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	if !status {
		ErrUnauthorized(c, error2.Unauthorized)
		return
	}
	c.JSON(200, gin.H{
		"msg": "success",
	})
	return
}
