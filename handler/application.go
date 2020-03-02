package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	error2 "github.com/2020-LonelyPlanet-backend/miniProject/pkg/error"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 处理请求
// @Description 根据请求Id来处理请求，id别的api会给出，通过状态status来处理请求，2为接受，3为拒绝
// @Tags application
// @Accept json
// @Produce json
// @Param status query string true "状态"
// @Param application_id path string true "请求id"
// @Param token header string true "token"
// @Param AcceptApplication body model.AcceptApplication true "联系方式和附加信息"
// @Success 200 {object} model.Res "{"msg":"success"}/{"msg":"需求已经被删除了!"}/{"msg":"已经处理过了!"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /application/solve/:application_id/ [put]
func SolveApplication(c *gin.Context) {
	uid := c.GetString("uid")
	status, err := strconv.Atoi(c.Query("status")) //2->接受　3->拒绝
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}

	if status != 2 && status != 3 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	if len(c.Param("application_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	applicationId, err := strconv.Atoi(c.Param("application_id"))
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}

	var tmpContent model.AcceptApplication
	if status == 2 {
		err := c.BindJSON(&tmpContent)
		if err != nil || len(tmpContent.ContactWay) != 2 {
			ErrBadRequest(c, error2.ParamBadRequest)
			return
		}
	}

	err, flag := model.SolveApplication(applicationId, status, uid, tmpContent)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	if flag == 4 {
		c.JSON(200, gin.H{
			"msg": "需求已经被删除了!",
		})
		return
	}

	if flag == 3 {
		ErrUnauthorized(c, error2.Unauthorized)
		return
	}
	if flag == 2 {
		c.JSON(200, gin.H{
			"msg": "已经处理过了!",
		})
		return
	}

	if flag == 1 {
		c.JSON(200, gin.H{
			"msg": "success",
		})
	}

	return
}

// @Summary 申请提醒
// @Description 点击申请提醒　查看所有待确认和接受的申请
// @Tags application
// @Accept json
// @Produce json
// @Param limit query string true "每页数量"
// @Param page query string true "当前请求页数，从0开始"
// @Param token header string true "token"
// @Success 200 {object} model.ApplicationView "{"msg":"success", "num":数量, "applications":数组，其中包含每个的id}，其中red_point字段是用来表示是否未读，即单条信息是否显示小红点"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /application/todo/ [get]
func ViewAllApplicationRemind(c *gin.Context) {
	uid := c.GetString("uid")
	if len(c.Query("limit")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	if len(c.Query("page")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	limit, err := strconv.Atoi(c.Query("limit")) //商定limit
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

	result, err := model.ViewAllApplication(uid, offset, limit)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}

	c.JSON(200, gin.H{
		"msg":          "success",
		"num":          len(result),
		"applications": result,
	})
	return
}

/*
func ApplicationResult(c *gin.Context) {
	uid := c.GetString("uid")

	result, err :=
	if err != nil {
		c.JSON(400,gin.H{
			"msg":"wrong",
		})
		return
	}

	c.JSON(200,gin.H{
		"msg":"success",
		"num":len(result),
		"applications":result,
	})
	return
}
*/
