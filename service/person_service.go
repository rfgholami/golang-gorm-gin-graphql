package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})

}

func List2(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World! 2"})

}
