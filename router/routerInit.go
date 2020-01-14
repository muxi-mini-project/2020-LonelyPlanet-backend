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
}