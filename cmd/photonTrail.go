package main

import (
	"PhotonTrail-backend/internal/global"
	"PhotonTrail-backend/internal/middleware"
	"PhotonTrail-backend/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	err := initConfig()
	if err != nil {
		log.Fatalf("init.initConfig err: %v", err)
	} else {
		log.Println("init.initConfig success")
	}

	err = initDB()
	if err != nil {
		log.Fatalf("init.initDB err: %v", err)
	} else {
		log.Println("init.initDB success")
	}

}

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())

	router.InitRouters(r)
	r.Run(global.Config.App.Addr)
}
