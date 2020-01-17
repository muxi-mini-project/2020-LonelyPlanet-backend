package router

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/handler"
	"github.com/2020-LonelyPlanet-backend/miniProject/middleware"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()
	Router.POST("/login/",handler.UserLogin)
	Router.Use(middleware.JwtAAuth())
	Router.GET("/test",handler.Test)
	Router.GET("/user/info/",handler.Homepage)
	Router.POST("/user/changeInfo/",handler.ChangeInfo)
	Router.POST("/secret/create/",handler.DebunksCreate)
	Router.DELETE("/secret/delete/",handler.DebunksDelete)
	Router.GET("/secret/history/",handler.DebunksHistory)
	Router.GET("/secret/square/",handler.DebunksSquare)
	Router.POST("/comment/create/",handler.CommentCreate)
	Router.GET("/comment/{secret_id}/history/",handler.CommentHistory)
	Router.DELETE("/comment/delete/",handler.CommentDelete)
	Router.GET("/remind/night/remindbox/view",handler.RemindNightRemindboxView)
}