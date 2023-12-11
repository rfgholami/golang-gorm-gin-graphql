package service

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/models"
	"github.com/kwa0x2/GoLang-Postgre-API/repo"
	"net/http"
)

func List(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})

}

func CreatePerson(ctx *gin.Context) {

	var person models.Person
	err := ctx.BindJSON(&person)
	if err != nil {
		return
	}
	createdPerson := repo.CreatePerson(&person)
	ctx.JSON(http.StatusOK, createdPerson)
}
