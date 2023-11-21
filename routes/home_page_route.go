package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/service"
	"net/http"
)

func HomePageRoute(router *gin.Engine) {

	router.Use(service.SecurityFilter())
	router.GET("/", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
}
