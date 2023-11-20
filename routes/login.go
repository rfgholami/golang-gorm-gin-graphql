package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(router *gin.Engine) {

	router.POST("login", func(ctx *gin.Context) {
		status := new(Status)
		status.Status = "UP"
		ctx.JSON(http.StatusOK, status)
	})

	//users := router.Group("/api")
	//users.GET("/users", controller.GetUsers)

}
