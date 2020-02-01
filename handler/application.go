package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	error2 "github.com/2020-LonelyPlanet-backend/miniProject/pkg/error"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SolveApplication(c *gin.Context) {
	status, _ := strconv.Atoi(c.Query("status"))  //2->接受　3->拒绝
	if len(c.Param("application_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	applicationId, _ := strconv.Atoi(c.Param("application_id"))

	err := model.SolveApplication(applicationId, status)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg": "success",
	})
	return
}

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
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	var offset int
	offset = page* limit

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
