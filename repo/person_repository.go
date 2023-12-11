package repo

import (
	"github.com/kwa0x2/GoLang-Postgre-API/config"
	"github.com/kwa0x2/GoLang-Postgre-API/models"
)

func CreatePerson(dto *models.Person) models.Person {

	var person models.Person
	config.DB.Model(&person).Save(dto).Scan(&person)
	return person
}
