package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/response"
	"gorm.io/gorm"
)

func CheckError(ctx *gin.Context, result *gorm.DB) bool {
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   400,
			Status: "BAD REQUEST",
			Error:  string(result.Error.Error()),
		})
		return true
	}

	if result.RowsAffected == 0 {
		// ctx.JSON(http.StatusNotFound, gin.H{"error":"User not found or deleted"})
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:   404,
			Status: "NOT FOUND",
			Error:  "User not found or deleted",
		})
		return true
	}
	return false
}