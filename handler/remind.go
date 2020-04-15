package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	error2 "github.com/2020-LonelyPlanet-backend/miniProject/pkg/error"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 查询是否显示小红点
// @Description 根据返回值来判断是否显示提醒的小红点
// @Tags remind
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.RemindEx "{"msg":"success", "existence":"true/false"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /remind/day/remindbox/status/ [get]
func DayRemindExistence(c *gin.Context) {
	uid := c.GetString("uid")
	result, err := model.ConfirmRemindExist(uid)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}

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

// @Summary 回复提醒
// @Description 查看回复提醒
// @Tags remind
// @Accept json
// @Produce json
// @Param limit query string true "每页数量"
// @Param page query string true "当前请求页数，从0开始"
// @Param token header string true "token"
// @Success 200 {object} model.RemindBox "{"msg":"success", "num":数量, "content":数组，其中包含每个的id， 其中confirm是用来判断显示的内容是否带有小眼睛图标， 2为接受，3为拒绝，其中red_point字段是用来表示是否未读，即单条信息是否显示小红点}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /remind/day/remindbox/ [get]
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
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}

	var offset int
	offset = page * limit

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

// @Summary 更新申请人阅读状态
// @Description 更新申请人阅读状态, 需要在用户在回复提醒中点击需求或者点击小眼睛的同时, 通过请求这条来更新状态, 如果可以希望可以根据是否已读来判断是否进行此次请求以减少请求次数
// @Tags remind
// @Accept json
// @Produce json
// @Param application_id path string true "用户已查看的申请id, 在别的api中给出"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"msg":"success"} 成功"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /remind/day/remindbox/done/:application_id/ [post]
func UpdateRemindStatus1(c *gin.Context) {
	uid := c.GetString("uid")
	if len(c.Param("application_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	applicationId, err := strconv.Atoi(c.Param("application_id"))
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
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

// @Summary 更新收件人阅读状态
// @Description 更新申收件人阅读状态, 需要在用户在申请提醒中点击需求或者点击小眼睛或者直接处理请求的同时, 通过请求这条来更新状态, 如果可以希望可以根据是否已读来判断是否进行此次请求以减少请求次数
// @Tags application
// @Accept json
// @Produce json
// @Param application_id path string true "用户已查看的申请id, 在别的api中给出"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"msg":"success"} 成功"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /application/done/:application_id/ [post]
func UpdateRemindStatus2(c *gin.Context) {
	uid := c.GetString("uid")
	if len(c.Param("application_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	applicationId, err := strconv.Atoi(c.Param("application_id"))
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}

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
