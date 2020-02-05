package router

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/handler"
	"github.com/2020-LonelyPlanet-backend/miniProject/middleware"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()
	Router.POST("/login/", handler.UserLogin) //用户登录
	Router.Use(middleware.JwtAAuth())
	Router.GET("/test", handler.Test)
	Router.GET("/user/info/", handler.Homepage)          //用户主页
	Router.POST("/user/changeInfo/", handler.ChangeInfo) //更改用户信息

	Router.GET("/requirement/square/", handler.Square)                           //白天需求广场，即筛选需求
	Router.GET("/requirement/view/:requirement_id/", handler.ViewRequirement)     //查看特定的需求
	Router.DELETE("/requirement/:requirement_id/", handler.DeleteRequirement)    //删除需求
	Router.PUT("/requirement/create/", handler.PostRequirement)                 //发布需求
	Router.GET("/requirement/history/", handler.HistoryRequirement)              //历史需求
	Router.POST("/requirement/apply/:requirement_id/", handler.ApplyRequirement) //申请需求

	Router.GET("/remind/day/remindbox/status/", handler.DayInformationExistence)            //查询是否显示小红点
	Router.GET("/remind/day/remindbox/", handler.ReminderBox)                               //点击回复提醒
	Router.POST("/remind/day/remindbox/info/:application_id/", handler.UpdateRemindStatus1) //更新申请人阅读状态
	//Router.GET("/remind/night/remindbox/",handler.NightInformationExistence)

	Router.PUT("/application/solve/:application_id/", handler.SolveApplication)         //确认是否接受申请
	Router.GET("/application/unsolve/", handler.ViewAllApplicationRemind)               //点击申请提醒　查看所有待确认和接受的申请
	Router.POST("/application/view_info/:application_id/", handler.UpdateRemindStatus2) //更新收件人阅读状态
	//Router.GET("/application/{application_id}/result/")

//	Router.GET("/test",handler.Test)
//	Router.GET("/user/info/",handler.Homepage)
//	Router.POST("/user/changeInfo/",handler.ChangeInfo)
	Router.POST("/secret/create/",handler.DebunksCreate)
	Router.DELETE("/secret/delete/",handler.DebunksDelete)
	Router.GET("/secret/history/",handler.DebunksHistory)
	Router.GET("/secret/square/",handler.DebunksSquare)
	Router.POST("/comment/create/",handler.CommentCreate)
	Router.GET("/comment/{secret_id}/history/",handler.CommentHistory)
	Router.DELETE("/comment/delete/",handler.CommentDelete)
	Router.GET("/remind/night/remindbox/view",handler.RemindNightRemindboxView)

}