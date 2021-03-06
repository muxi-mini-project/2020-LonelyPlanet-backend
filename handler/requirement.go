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

func detectParamSelect(tmp condition, tag []int, place []int) bool {
	if tmp.TypeOne == 1 {
		if len(tag) != 0 {
			for _, v := range tag {
				if v > 4 || v < 1 {
					return false
				}
			}
		}
		if len(place) != 0 {
			for _, v := range place {
				if v > 3 || v < 1 {
					return false
				}
			}
		}
		if !detectParamSelectTime(tmp.TimeFrom, tmp.TimeEnd) {
			return false
		}
		return true
	}
	if tmp.TypeOne == 2 {
		if len(tag) != 0 {
			for _, v := range tag {
				if v > 7 || v < 1 {
					return false
				}
			}
		}
		if len(place) != 0 {
			for _, v := range place {
				if v > 8 || v < 1 {
					return false
				}
			}
		}
		return true
	}
	if tmp.TypeOne == 3 {
		if len(tag) != 0 {
			for _, v := range tag {
				if v > 4 || v < 1 {
					return false
				}
			}
		}
		if len(place) != 0 {
			for _, v := range place {
				if v > 3 || v < 1 {
					return false
				}
			}
		}
		return true
	}
	if tmp.TypeOne == 4 {
		if len(place) != 0 {
			for _, v := range place {
				if v > 3 || v < 1 {
					return false
				}
			}
		}
		return true
	}
	return false
}

func detectParamSelectTime(from, end int) bool {
	if end != 0 {
		if from != 0 && end < from {
			return false
		}
		if end < 0 || end > 24 {
			return false
		}
	}
	if from != 0 {
		if from < 0 || from > 24 {
			return false
		}
	}
	return true
}

// @Summary 需求广场
// @Description 给出用户的筛选条件，返回若干需求
// @Tags requirement
// @Accept json
// @Produce json
// @Param limit query string true "每页数量"
// @Param page query string true "当前请求页数，从0开始"
// @Param type query string true "主类别，必要"
// @Param tag query string false "第二级标签"
// @Param place query string false "地点，将复合条件整合为字符串"
// @Param time_from query string false "起始时间"
// @Param time_end query string false "终止时间"
// @Param date query string false "复合条件的日期筛选条件，如果有整合为8位字符串'1xxxxxxx'，最低位为周一"
// @Param token header string true "token"
// @Success 200 {object} model.Square "{"msg":"get result successful", "num":数量, "content":数组，其中包含每个的id},若数量为零msg:'none'"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /requirement/square/ [get]
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
	var err error
	tmpCondition.TypeOne, err = strconv.Atoi(type1)
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	tmpCondition.Tag = c.Query("tag")
	tmpCondition.Place = c.Query("place")
	if len(c.Query("time_from")) != 0 {
		tmpCondition.TimeFrom, err = strconv.Atoi(c.Query("time_from"))
		if err != nil {
			ErrBadRequest(c, error2.ParamBadRequest)
			return
		}
	}
	if len(c.Query("time_from")) == 1 {
		tmpCondition.TimeFrom = 0
	}
	if len(c.Query("time_end")) != 0 {
		tmpCondition.TimeEnd, err = strconv.Atoi(c.Query("time_end"))
		if err != nil {
			ErrBadRequest(c, error2.ParamBadRequest)
			return
		}
	}
	if len(c.Query("time_end")) == 24 {
		tmpCondition.TimeEnd = 0
	}
	if len(c.Query("date")) != 0 {
		tmpCondition.Date, err = model.BinStr2Dec(c.Query("date"))
		if err != nil {
			ErrBadRequest(c, error2.ParamBadRequest)
			return
		}
		if tmpCondition.Date < 128 || tmpCondition.Date > 255 {
			fmt.Println(tmpCondition.Date)
			ErrBadRequest(c, error2.ParamBadRequest)
			return
		}
	}
	if len(c.Query("date")) == 0 {
		tmpCondition.Date = 0
	}

	//1.4新增筛选发布时间
	//second_time: 1:第一个选项 2:第二个选项 3:第三个选项
	tmpSecondTime := c.Query("second_time")
	secondTime, err := strconv.Atoi(tmpSecondTime)
	fmt.Println(secondTime)

	tmpCondition.Limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	tmpCondition.Page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}

	tag, err := model.ConvertStringToIntSlice(tmpCondition.Tag)
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	if tmpCondition.TypeOne == 4 {
		tag = []int{}
	}

	place, err := model.ConvertStringToIntSlice(tmpCondition.Place)
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}

	flag := detectParamSelect(tmpCondition, tag, place)
	if !flag {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}

	var offset int
	offset = tmpCondition.Page*1 + 6
	if tmpCondition.Page == 0 {
		offset = 0
	}

	uid := c.GetString("uid")
	result, err := model.RequirementFind(tmpCondition.TypeOne, uid, tmpCondition.Date, tmpCondition.TimeFrom, tmpCondition.TimeEnd, tag, place, secondTime, tmpCondition.Limit, offset)

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

// @Summary 查看特定需求
// @Description 根据id来查看特定的需求
// @Tags requirement
// @Accept json
// @Produce json
// @Param requirement_id path string true "查看需求的id，会在别的api中给出"
// @Param token header string true "token"
// @Success 200 {object} model.ViewRequirement "{"msg":"success", "content":数组，其中包含每个的id},若该需求被删除 msg:'不见啦'"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /requirement/view/:requirement_id/ [get]
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
			"msg":     "不见啦", //删除了
			"content": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":     "success",
		"content": result,
	})
	return
}

// @Summary 删除已发布的需求
// @Description 根据id来删除特定的需求
// @Tags requirement
// @Accept json
// @Produce json
// @Param requirement_id path string true "删除需求的id，会在别的api中给出"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"msg":"success"} 成功"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /requirement/:requirement_id/ [delete]
func DeleteRequirement(c *gin.Context) {
	if len(c.Param("requirement_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	requirementId, err := strconv.Atoi(c.Param("requirement_id"))
	if err != nil {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	uid := c.GetString("uid")
	err, status := model.RequirementDelete(requirementId, uid) //判断是否是本人操作 多次删除 无影响
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

func detectPostRequirement(tmp model.Requirements) bool {
	if tmp.TimeEnd == 0 || tmp.TimeFrom == 0 || tmp.Date == 0 || len(tmp.Title) == 0 || len(tmp.Content) == 0 {
		return false
	}
	//各属性长度 限定 -> 确认
	if tmp.Tag == 0 || tmp.Type == 0 || tmp.RequirementId != 0 || tmp.Status != 0 {
		if tmp.Type != 4 {
			return false
		}
	}
	//确定 是否 越界
	tmpCondition := condition{
		TypeOne: tmp.Type,
	}
	if !detectParamSelect(tmpCondition, []int{tmp.Tag}, []int{tmp.Place}) {
		return false
	}
	if !detectParamSelectTime(tmp.TimeFrom, tmp.TimeEnd) {
		return false
	}
	return true
}

// @Summary 发布需求
// @Description 用户发布需求
// @Tags requirement
// @Accept json
// @Produce json
// @Param requirement body model.NewRequirements true "新发布的需求详情"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"msg":"success"} 成功 {"msg":"requirement already exist."} 提示重复发布需求了"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /requirement/new/ [put]
func PostRequirement(c *gin.Context) {
	uid := c.GetString("uid")

	black, err := model.ConfirmBlackList(uid)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	if black {
		ErrUnauthorized(c, error2.BlackList)
		return
	}

	var newRequirement model.Requirements
	err = c.BindJSON(&newRequirement)
	if err != nil {
		log.Println("PostRequirement err", err)
		ErrBadRequest(c, error2.BadRequest)
		return
	}

	if newRequirement.IsDraft != 1 { //不是草稿
		err = model.NewRecode(uid, 1, 60) //新增记录
		if err != nil {
			ErrServerError(c, error2.ServerError)
			return
		}
		num, err := model.InspectNum(uid, 1)
		if err != nil {
			ErrServerError(c, error2.ServerError)
			return
		}
		if num > 2 {
			ErrBadRequest(c, error2.FrequentRequests1) //暂且这样
			return
		}

		num2, err := model.InspectNum(uid, 3)
		if err != nil {
			ErrServerError(c, error2.ServerError)
			return
		}

		if num2 > 15 {
			ErrBadRequest(c, error2.FrequentRequests2) //暂且这样
			return
		}
		//fmt.Println(num, num2)

		if !detectPostRequirement(newRequirement) {
			ErrBadRequest(c, error2.ParamBadRequest)
			return
		}

		newRequirement.PostTime = model.NowTimeStampStr()
		newRequirement.SenderSid = uid

		if newRequirement.Date < 1000000 || newRequirement.Date > 11111111 {
			ErrBadRequest(c, error2.ParamBadRequest)
			return
		}
		tmpDate := strconv.Itoa(newRequirement.Date)
		newRequirement.Date, err = model.BinStr2Dec(tmpDate)
		if err != nil {
			ErrBadRequest(c, error2.ParamBadRequest)
			return
		}
		err, status := model.ConfirmRequirementExist(newRequirement)
		if err != nil {
			ErrServerError(c, error2.ServerError)
			return
		}
		if status {
			c.JSON(200, gin.H{
				"msg": "requirement already exist.",
			})
			return
		}
	} else {
		/*
			newRequirement.PostTime = model.NowTimeStampStr()
			newRequirement.SenderSid = uid
			tmpDate := strconv.Itoa(newRequirement.Date)
			newRequirement.Date, err = model.BinStr2Dec(tmpDate)
			if err != nil {
				ErrBadRequest(c, error2.ParamBadRequest)
				return
			}
		*/
		newRequirement.PostTime = model.NowTimeStampStr()
		newRequirement.SenderSid = uid
		newRequirement.Status = 3
	}

	err = model.CreatRequirement(newRequirement)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	/*
		err = model.NewRecode(uid, 1) //新增记录
		if err != nil {
			ErrServerError(c, error2.ServerError)
			return
		}
	*/
	c.JSON(200, gin.H{
		"msg": "success",
	})
	return
}

// @Summary 历史发布需求
// @Description 用户历史所发布需求
// @Tags requirement
// @Accept json
// @Produce json
// @Param limit query string true "每页数量"
// @Param page query string true "当前请求页数，从0开始"
// @Param token header string true "token"
// @Success 200 {object} model.ViewHistoryRequirement "{"msg":"success", "num":数量, "history":数组，其中包含每个的id}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /requirement/history/ [get]
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
	offset = page * limit

	tmpRequirement, err := model.HistoryRequirementFind(uid, offset, limit)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg":     "success",
		"num":     len(tmpRequirement), //记得提醒到底啦，或者当page=0的时候显示没有记录！
		"history": tmpRequirement,
	})
	return
}

// @Summary 申请需求
// @Description 根据id来申请特定的需求
// @Tags requirement
// @Accept json
// @Produce json
// @Param requirement_id path string true "申请需求的id，会在别的api中给出"
// @Param application body model.Application true "联系方式和附加信息"
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"msg":"success"}/{"msg":"不能申请自己的需求!"}/{"msg":"已经申请过了!"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /requirement/application/:requirement_id/ [post]
func ApplyRequirement(c *gin.Context) {
	uid := c.GetString("uid")

	black, err := model.ConfirmBlackList(uid)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	if black {
		ErrUnauthorized(c, error2.BlackList)
		return
	}

	if len(c.Param("requirement_id")) == 0 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	requirementId, _ := strconv.Atoi(c.Param("requirement_id"))
	var tmpContract model.Application
	err = c.BindJSON(&tmpContract)
	//fmt.Println(tmpContract)
	if len(tmpContract.ContactWay) != 2 {
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	if err != nil {
		log.Print("ApplyRequirement err")
		fmt.Println(err)
		ErrBadRequest(c, error2.ParamBadRequest)
		return
	}
	flag, err := model.RequirementApply(uid, requirementId, tmpContract.ContactWay[0], tmpContract.ContactWay[1], tmpContract.Content)
	if flag == 2 {
		ErrBadRequest(c, error2.BadRequest)
		return
	}

	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}

	if flag == 4 {
		c.JSON(200, gin.H{
			"msg": "不能申请自己的需求!",
		})
		return
	}

	if flag == 3 {
		c.JSON(200, gin.H{
			"msg": "已经申请过了!", //提示已经申请过啦！
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

func FindDraft(c *gin.Context) {
	uid := c.GetString("uid")

	err, num, tmpResult := model.FindDraft(uid)
	if err != nil {
		log.Println("find draft err: ", err)
		ErrServerError(c, error2.ServerError)
		return
	}
	var result model.Draft
	if num == 0 {
		result.HasDraft = 2 //没有草稿
	} else {
		result.HasDraft = 1
		result.Content = tmpResult
	}
	c.JSON(200, gin.H{
		"msg":   "success",
		"draft": result,
	})
	return
}

func DeleteDraft(c *gin.Context) {
	uid := c.GetString("uid")

	err := model.DeleteDraft(uid)
	if err != nil {
		log.Println("delete draft err: ", err)
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg": "success",
	})
	return
}
