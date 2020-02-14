package main

import (
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	"github.com/2020-LonelyPlanet-backend/miniProject/router"
	"log"
)

// @title lonely planet
// @version 1.0
// @description 孤独星球

// @host 47.97.74.180:9090
// @BasePath /

// @Schemas http

func main() {
	model.Db.Init()
	defer model.Db.Close()
	router.Init()
	if err := router.Router.Run(":9090"); err != nil {
		log.Print("路由错误")
		log.Fatal(err)
	}
}
