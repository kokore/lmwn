package main

import (
	"covid/handler/covid"
	"covid/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	initial()
	ginEngine := getGinEngine()

	port := fmt.Sprintf(":%s", config.Config.App.Port)
	ginEngine.Run(port)
}

func initial() {
	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func getGinEngine() *gin.Engine {
	router := gin.Default()

	root := router.Group("/covid")

	root.GET("/summary", covid.GetCovidSummary)

	return router
}
