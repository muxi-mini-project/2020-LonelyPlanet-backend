package main

import (
	"github.com/kocoler/GoExercises/miniProject/model"
	"github.com/kocoler/GoExercises/miniProject/router"
	"log"
)

func main() {
	model.Db.Init()
	defer model.Db.Close()
	router.Init()
	if err := router.Router.Run(":9090"); err != nil {
		log.Print("路由错误")
		log.Fatal(err)
	}
<<<<<<< HEAD
}
=======
}
>>>>>>> d3ce8f162ae41d571dab45e986131f21981720f3
