package handler

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	error2 "github.com/2020-LonelyPlanet-backend/miniProject/pkg/error"
	"github.com/2020-LonelyPlanet-backend/miniProject/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func DayReport(c *gin.Context) {
	uid := c.GetString("uid")

	var newReport model.ReportInformation
	err := c.BindJSON(&newReport)
	if err != nil {
		log.Println("PostRequirement err", err)
		ErrBadRequest(c, error2.BadRequest)
		return
	}

	requirementId, _ := strconv.Atoi(c.Param("requirement_id"))
	result, err := model.GetInfoFromRequirementId(requirementId)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	err = util.SendMail(1, result.Content, newReport.Reason, uid, result.SenderSid, newReport.Addition)
	if err != nil {
		log.Println("sendmail err: ", err)
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg": "success",
	})
	return
}