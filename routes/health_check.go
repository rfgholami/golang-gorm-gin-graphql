package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Status struct {
	Status string `json:"status"`
}

func HealthCheck(router *gin.Engine) {

	router.GET("health-check", func(ctx *gin.Context) {
		status := new(Status)
		status.Status = "UP"
		ctx.JSON(http.StatusOK, status)
	})

	//users := router.Group("/api")
	//users.GET("/users", controller.GetUsers)

}
