package service

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/GoLang-Postgre-API/models"
	"github.com/kwa0x2/GoLang-Postgre-API/repo"
	"log"
	"net/http"
	"strconv"
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

func DeletePerson(ctx *gin.Context) {

	personID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	createdPerson := repo.DeletePerson(personID)
	ctx.JSON(http.StatusOK, createdPerson)
}

func FindPersonById(ctx *gin.Context) {

	personID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	person := repo.FindPersonById(personID)
	ctx.JSON(http.StatusOK, person)
}
