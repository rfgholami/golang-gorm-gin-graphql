package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/service"
)

func UserManagement(router *gin.Engine) {

	users := router.Group("/")
	users.POST("/login", service.LoginRest)
	users.POST("/get-user-info", service.GetUserInfo)

}
