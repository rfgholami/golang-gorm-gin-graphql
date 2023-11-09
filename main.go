package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/config"
	"github.com/kwa0x2/GoLang-Postgre-API/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":9898")
}