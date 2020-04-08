package router

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/handler"
	"github.com/2020-LonelyPlanet-backend/miniProject/handler/sd"
	"github.com/2020-LonelyPlanet-backend/miniProject/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	// The health check handlers

	g.GET("/Swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	g.POST("/lonely_planet/v1/login/", handler.UserLogin) //用户登录
	g.Use(middleware.JwtAAuth())
	g.GET("/test", handler.Test)

	user := g.Group("/lonely_planet/v1/user")
	{
		user.GET("/info/", handler.Homepage)           //用户主页
		user.POST("/change_info/", handler.ChangeInfo) //更改用户信息
	}

	requirement := g.Group("/lonely_planet/v1/requirement")
	{
		requirement.GET("/square/", handler.Square)                                 //白天需求广场，即筛选需求
		requirement.GET("/view/:requirement_id/", handler.ViewRequirement)          //查看特定的需求
		requirement.DELETE("/:requirement_id/", handler.DeleteRequirement)          //删除需求
		requirement.PUT("/new/", handler.PostRequirement)                           //发布需求
		requirement.GET("/history/", handler.HistoryRequirement)                    //历史需求
		requirement.POST("/application/:requirement_id/", handler.ApplyRequirement) //申请需求
	}

	remind := g.Group("/lonely_planet/v1/remind")
	{
		remind.GET("/day/remindbox/status/", handler.DayRemindExistence)                 //查询是否显示小红点
		remind.GET("/day/remindbox/", handler.ReminderBox)                               //点击回复提醒
		remind.POST("/day/done/:application_id/", handler.UpdateRemindStatus1) //更新申请人阅读状态
		//Router.GET("/remind/night/remindbox/",handler.NightInformationExistence)

		//下面的是黑夜的
		remind.GET("/night/remindbox/view/", handler.RemindNightRemindboxView)
		remind.GET("/night/remindbox/status/",handler.NightRemindExistence)
		//remind.GET("/night/remindbox/:comment_id/",handler.UpdateNightRemindStatus)
	}

	application := g.Group("/lonely_planet/v1/application")
	{
		application.PUT("/:application_id/", handler.SolveApplication)          //确认是否接受申请
		application.GET("/todo/", handler.ViewAllApplicationRemind)             //点击申请提醒　查看所有待确认和接受的申请
		application.POST("/done/:application_id/", handler.UpdateRemindStatus2) //更新收件人阅读状态
		//Router.GET("/application/{application_id}/result/")
	}

	secret := g.Group("/lonely_planet/v1/secret")
	{
		secret.POST("/create/", handler.DebunksCreate)
		secret.GET("/view/:secret_id/", handler.GetSecret)
		secret.DELETE("/delete/:secret_id/", handler.DebunksDelete)
		secret.GET("/history/", handler.DebunksHistory)
		secret.GET("/square/", handler.DebunksSquare)
	}

	comment := g.Group("/lonely_planet/v1/comment")
	{
		comment.POST("/create/", handler.CommentCreate)
		comment.GET("/history/:secret_id/", handler.CommentHistory)
		comment.DELETE("/delete/:comment_id/", handler.CommentDelete)
	}



	return g
}