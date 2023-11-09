package config

import (
	"fmt"
	"time"
	
	"github.com/kwa0x2/GoLang-Postgre-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@127.18.0.2/nettasec?sslmode=disable"), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil{
		panic(err)
	}
	start := time.Now()
	for sqlDB.Ping() != nil{
		if(start.After(start.Add(10 * time.Second))) {
			fmt.Println("Failed to connect database after 10 seconds")
			break
		}
	}
	fmt.Println("Connected to database: ", sqlDB.Ping() == nil)

	db.AutoMigrate(&models.User{})

	DB = db

}
