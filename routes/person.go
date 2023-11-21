package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/service"
)

func Person(router *gin.Engine) {

	users := router.Group("/persons")
	users.GET("/", service.List)
	users.GET("/edit/", service.List2)

}
