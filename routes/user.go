package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/controller"
)

func UserRoute(router *gin.Engine) {

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	users := router.Group("/api")
	users.GET("/users", controller.GetUsers)
	users.POST("/users", controller.PostUser)
	users.GET("/users/:id", controller.GetUserByID)
	users.DELETE("/users/:id", controller.DeleteUserByID)
	users.PUT("/users/:id", controller.PutUserByID)
	users.PATCH("/users/:id", controller.PatchUser)


}
