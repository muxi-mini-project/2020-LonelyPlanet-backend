package handler

import (
	"fmt"
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	error2 "github.com/2020-LonelyPlanet-backend/miniProject/pkg/error"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type condition struct {
	TypeOne  int
	Tag      string
	Place    string
	TimeFrom int
	TimeEnd  int
	Date     int
	Limit    int
	Page     int
}

type contractWay struct {
	ContractWayType string `json:"contract_way_type"`
	ContractWay     string `json:"contract_way"`
}

func Square(c *gin.Context) {
	var tmpCondition condition
	type1 := c.Query("type")
	if len(c.Query("type")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	if len(c.Query("limit")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	if len(c.Query("page")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	tmpCondition.TypeOne, _ = strconv.Atoi(type1)
	tmpCondition.Tag = c.Query("tag")
	tmpCondition.Place = c.Query("place")
	if len(c.Query("time_from")) != 0 {
		tmpCondition.TimeFrom, _ = strconv.Atoi(c.Query("time_from"))
	}
	if len(c.Query("time_from")) == 0 {
		tmpCondition.TimeFrom = 0
	}
	if len(c.Query("time_end")) != 0 {
		tmpCondition.TimeEnd, _ = strconv.Atoi(c.Query("time_end"))
	}
	if len(c.Query("time_end")) == 0 {
		tmpCondition.TimeEnd = 0
	}
	if len(c.Query("date")) != 0 {
		tmpCondition.Date = model.BinStr2Dec(c.Query("date"))
	}
	if len(c.Query("date")) == 0 {
		tmpCondition.Date = 0
	}

	tmpCondition.Limit, _ = strconv.Atoi(c.Query("limit"))
	tmpCondition.Page, _ = strconv.Atoi(c.Query("page"))

	tag := model.ConvertStringToIntSlice(tmpCondition.Tag)
	place := model.ConvertStringToIntSlice(tmpCondition.Place)
	var offset int
	offset = tmpCondition.Page*1 + 6
	if tmpCondition.Page == 0 {
		offset = 0
	}

	uid := c.GetString("uid")
	result, err := model.RequirementFind(tmpCondition.TypeOne, uid, tmpCondition.Date, tmpCondition.TimeFrom, tmpCondition.TimeEnd, tag, place, tmpCondition.Limit, offset)

	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}

	if len(result) == 0 {
		c.JSON(200, gin.H{
			"msg":     "none",
			"num":     0,
			"content": result,
		})
		//返回没有数值了，重新输入吧
		return
	}
	//返回数据

	if len(result) != 0 {
		c.JSON(200, gin.H{
			"msg":     "get result successful",
			"num":     len(result),
			"content": result,
		})
	}


	/*if len(result) != 6 {
		c.JSON(200, gin.H{
			"msg":     "get result successful",
			"num":     len(result),
			"content": result,
		})
	}*/
	return
}

func ViewRequirement(c *gin.Context) {
	if len(c.Param("requirement_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	requirementId, _ := strconv.Atoi(c.Param("requirement_id"))
	result, status, err := model.RequirementInfo(requirementId)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	if status {
		c.JSON(200, gin.H{
			"msg":"不见啦",  //删除了
			"content":nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":"success",
		"content":result,
	})
	return
}

func DeleteRequirement(c *gin.Context) {
	if len(c.Param("requirement_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	requirementId, _ := strconv.Atoi(c.Param("requirement_id"))
	uid := c.GetString("uid")
	err, status := model.RequirementDelete(requirementId, uid)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	if !status {
		ErrUnauthorized(c, error2.Unauthorized)
		return
	}
	if status {
		c.JSON(200, gin.H{
			"msg": "success",
		})
	}
	return
}

func PostRequirement(c *gin.Context) {
	var newRequirement model.Requirements
	err := c.BindJSON(&newRequirement)
	if err != nil {
		log.Print("PostRequirement err")
		fmt.Println(err)
		ErrBadRequest(c, error2.BadRequest)
		return
	}
	uid := c.GetString("uid")
	newRequirement.PostTime = model.NowTimeStampStr()
	newRequirement.SenderSid = uid
	tmpDate := strconv.Itoa(newRequirement.Date)
	newRequirement.Date = model.BinStr2Dec(tmpDate)
	err = model.CreatRequirement(newRequirement)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg": "success",
	})
	return
}

func HistoryRequirement(c *gin.Context) {
	uid := c.GetString("uid")
	if len(c.Query("limit")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	if len(c.Query("page")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))

	var offset int
	offset = page* limit

	tmpRequirement, err := model.HistoryRequirementFind(uid, offset, limit)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg":     "success",
		"num":     len(tmpRequirement),    //记得提醒到底啦，或者当page=0的时候显示没有记录！
		"history": tmpRequirement,
	})
	return
}

func ApplyRequirement(c *gin.Context) {
	uid := c.GetString("uid")
	if len(c.Param("requirement_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	requirementId, _ := strconv.Atoi(c.Param("requirement_id"))
	var tmpContract contractWay
	err := c.BindJSON(&tmpContract)
	if err != nil {
		log.Print("ApplyRequirement err")
		fmt.Println(err)
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	exist, err := model.RequirementApply(uid, requirementId, tmpContract.ContractWayType, tmpContract.ContractWay)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}

	if exist {
		c.JSON(200, gin.H{
			"msg": "application already exist!",    //提示已经申请过啦！
		})
		return
	}

	if !exist {
		c.JSON(200, gin.H{
			"msg": "success",
		})
	}
	return
}
