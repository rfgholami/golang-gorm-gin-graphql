package repo

import (
	"github.com/kwa0x2/GoLang-Postgre-API/config"
	"github.com/kwa0x2/GoLang-Postgre-API/models"
)

func CreatePerson(dto *models.Person) models.Person {

	var person models.Person
	config.DB.Model(&person).Save(dto).Scan(&person)
	return *dto
}

func DeletePerson(id int) int {

	tx := config.DB.Delete(&models.Person{}, id)
	tx.Commit()
	return id
}

func FindPersonById(id int) models.Person {
	var person models.Person
	config.DB.Preload("Addresses").Find(&person, id)

	return person
}

func FindPersonById2(id int) models.PersonDto {
	var person models.PersonDto
	config.DB.Find(&person, id)
	config.DB.Table(`"go"."people"`).
		Select(`"people"."id", "people"."name","addresses"."id" as address_id,"addresses"."title" as address_title`).
		InnerJoins(`left join "go"."addresses" on "addresses".person_id="people"."id"`).
		Where(&models.Person{ID: id}).
		Find(&person)

	return person
}
