package main

import (
	"errors"
	"github.com/2020-LonelyPlanet-backend/miniProject/config"
	"github.com/2020-LonelyPlanet-backend/miniProject/model"
	"github.com/2020-LonelyPlanet-backend/miniProject/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

// @title lonely planet
// @version 1.0
// @description 孤独星球

// @host 47.97.74.180:9090
// @BasePath /lonely_planet/v1/

// @Schemas http

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	model.Db.Init()
	defer model.Db.Close()

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,

		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":9090")
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:9090" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
