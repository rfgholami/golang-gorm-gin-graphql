package repo

import (
	"github.com/kwa0x2/GoLang-Postgre-API/config"
	"github.com/kwa0x2/GoLang-Postgre-API/models"
)

func GetUserByID(id int) models.User {

	//result := config.DB.Find(&user, id).Scan(&user).First(&user)
	//return result
	var user models.User

	config.DB.Model(&user).Where("id = ?", id).Scan(&user)
	return user
}

func GetUsers(limit int) []models.User {

	var users []models.User

	config.DB.Model(models.User{}).Limit(limit).Find(&users)
	return users
}
func GetUsersPageable(page int, size int) []models.User {

	var users []models.User

	offset := page * size

	config.DB.Model(models.User{}).Offset(offset).Limit(size).Find(&users)
	return users
}

func Create(dto *models.User) models.User {

	var user models.User

	config.DB.Model(&user).Save(dto).Scan(&user)
	return user
}

func Update(dto *models.User) models.User {

	var user models.User

	config.DB.Model(&user).Where("id=?", dto.ID).Save(dto).Scan(&user)
	return user
}
