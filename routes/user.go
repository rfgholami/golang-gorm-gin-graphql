package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/controller"
)

func UserRoute(router *gin.Engine) {
	users := router.Group("/api/users")
	users.GET("/get", controller.ListUsers)
}
