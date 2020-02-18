package router

import (
	_ "github.com/2020-LonelyPlanet-backend/miniProject/docs"
	"github.com/2020-LonelyPlanet-backend/miniProject/handler"
	"github.com/2020-LonelyPlanet-backend/miniProject/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()
	Router.GET("/Swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	Router.POST("/lonely_planet/v1/login/", handler.UserLogin) //用户登录
	Router.Use(middleware.JwtAAuth())
	Router.GET("/test", handler.Test)
	Router.GET("/lonely_planet/v1/user/info/", handler.Homepage)           //用户主页
	Router.POST("/lonely_planet/v1/user/change_info/", handler.ChangeInfo) //更改用户信息

	Router.GET("/lonely_planet/v1/requirement/square/", handler.Square)                                 //白天需求广场，即筛选需求
	Router.GET("/lonely_planet/v1/requirement/view/:requirement_id/", handler.ViewRequirement)          //查看特定的需求
	Router.DELETE("/lonely_planet/v1/requirement/:requirement_id/", handler.DeleteRequirement)          //删除需求
	Router.PUT("/lonely_planet/v1/requirement/new/", handler.PostRequirement)                           //发布需求
	Router.GET("/lonely_planet/v1/requirement/history/", handler.HistoryRequirement)                    //历史需求
	Router.POST("/lonely_planet/v1/requirement/application/:requirement_id/", handler.ApplyRequirement) //申请需求

	Router.GET("/lonely_planet/v1/remind/day/remindbox/status/", handler.DayRemindExistence)                 //查询是否显示小红点
	Router.GET("/lonely_planet/v1/remind/day/remindbox/", handler.ReminderBox)                               //点击回复提醒
	Router.POST("/lonely_planet/v1/remind/day/remindbox/done/:application_id/", handler.UpdateRemindStatus1) //更新申请人阅读状态
	//Router.GET("/remind/night/remindbox/",handler.NightInformationExistence)

	Router.PUT("/lonely_planet/v1/application/:application_id/", handler.SolveApplication)          //确认是否接受申请
	Router.GET("/lonely_planet/v1/application/todo/", handler.ViewAllApplicationRemind)             //点击申请提醒　查看所有待确认和接受的申请
	Router.POST("/lonely_planet/v1/application/done/:application_id/", handler.UpdateRemindStatus2) //更新收件人阅读状态
	//Router.GET("/application/{application_id}/result/")

	//	Router.GET("/test",handler.Test)
	//	Router.GET("/user/info/",handler.Homepage)
	//	Router.POST("/user/changeInfo/",handler.ChangeInfo)
	Router.POST("/lonely_planet/v1/secret/create/", handler.DebunksCreate)
	Router.DELETE("/lonely_planet/v1/secret/delete/:secret_id/", handler.DebunksDelete)
	Router.GET("/lonely_planet/v1/secret/history/", handler.DebunksHistory)
	Router.GET("/lonely_planet/v1/secret/square/", handler.DebunksSquare)
	Router.POST("/lonely_planet/v1/comment/create/", handler.CommentCreate)
	Router.GET("/lonely_planet/v1/comment/history/:secret_id/", handler.CommentHistory)
	Router.DELETE("/lonely_planet/v1/comment/delete/:comment_id/", handler.CommentDelete)
	Router.GET("/lonely_planet/v1/remind/night/remindbox/view", handler.RemindNightRemindboxView)
}
