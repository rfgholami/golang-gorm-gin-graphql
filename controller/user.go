package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/config"
	"github.com/kwa0x2/GoLang-Postgre-API/helper"
	"github.com/kwa0x2/GoLang-Postgre-API/models"
	"github.com/kwa0x2/GoLang-Postgre-API/response"
)

var users models.User

func GetUsers(ctx *gin.Context) {
	var users []models.User
	result := config.DB.Raw("SELECT * FROM users ORDER BY id asc").Scan(&users)

	if helper.CheckError(ctx, result) {
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:     http.StatusOK,
		Status:   "OK",
		RowCount: int(result.RowsAffected),
		Data:     &users,
	})
}

func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result := config.DB.Raw("SELECT * FROM users WHERE id=?", id).Scan(&users)

	if helper.CheckError(ctx, result) {
		return
	}

	ctx.JSON(http.StatusFound, &users)
}

func DeleteUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result := config.DB.Exec("DELETE FROM users WHERE id=?", id)
	if helper.CheckError(ctx, result) {
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User deleted successfully",
	})
}

func PutUserByID(ctx *gin.Context) {
	ctx.BindJSON(&users)

	id := ctx.Param("id")

	result := config.DB.Exec("UPDATE users SET username=?, password=? WHERE id=?", users.Username, users.Password, id)

	if helper.CheckError(ctx, result) {
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User updated successfully",
	})
}

func PostUser(ctx *gin.Context) {
	ctx.BindJSON(&users)

	result := config.DB.Exec("INSERT INTO users(username,password) VALUES(?,?)", users.Username, users.Password)

	if helper.CheckError(ctx, result) {
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "User created successfully",
	})
}

func PatchUser(ctx *gin.Context) {
	ctx.BindJSON(&users)
	id := ctx.Param("id")

	var sets []string
	if users.Username != "" {
		sets = append(sets, "username='"+users.Username+"'")
	}
	if users.Password != "" {
		sets = append(sets, "password='"+users.Password+"'")
	}

	if len(sets) == 0 {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "BAD REQUEST",
			Error: "No valid fields to update",
		})
		return
	}

	result := config.DB.Exec("UPDATE users SET " + strings.Join(sets, ", ") + " WHERE id=" + id + "")

	if helper.CheckError(ctx, result) {
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "User updated successfully",
	})
}
