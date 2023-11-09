package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/config"
	"github.com/kwa0x2/GoLang-Postgre-API/models"
	"gorm.io/gorm"
)

func checkError(ctx *gin.Context, result *gorm.DB) bool {
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return true
	}
	return false
}

func GetUser(ctx *gin.Context) {
	var users []models.User
	result := config.DB.Raw("SELECT * FROM users ORDER BY id asc").Scan(&users)

	if checkError(ctx, result){
		return
	}

	ctx.JSON(http.StatusOK, &users)
}

func GetUserByID(ctx *gin.Context){
	id := ctx.Param("id")

	var users models.User
	result := config.DB.Raw("SELECT * FROM users WHERE id=?", id).Scan(&users)

	if checkError(ctx, result){
		return
	}

	ctx.JSON(http.StatusFound, &users)
}

func DeleteUserByID(ctx *gin.Context){
	id := ctx.Param("id")

	var users models.User
	result := config.DB.Raw("DELETE FROM users WHERE id=?", id).Scan(&users)
	if checkError(ctx, result){
		return
	}

	ctx.JSON(http.StatusOK, true)
}

func PostUser(ctx *gin.Context) {
	var users models.User
	ctx.BindJSON(&users)

	result := config.DB.Raw("INSERT INTO users(username,password) VALUES(?,?) RETURNING id", users.Username, users.Password).Scan(&users)

	if checkError(ctx, result){
		return
	}

	ctx.JSON(http.StatusCreated, &users)
}
