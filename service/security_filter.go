package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SecurityFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !CheckAccess(ctx) {
			ctx.JSON(http.StatusOK, gin.H{"message": "access.denied"})
			ctx.Abort()
		}

	}
}
