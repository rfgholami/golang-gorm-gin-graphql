package config

import (
	"fmt"
	"gorm.io/gorm/schema"
	"time"

	"github.com/kwa0x2/GoLang-Postgre-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//db, err := gorm.Open(postgres.Open("postgres://pfi:pfi$$123@<docker-service-name>/<database>?sslmode=disable"), &gorm.Config{})

	dsn := "host=development-db.rbc.local user=pfi password=pfi$$123 dbname=pfidb port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "go.",
				SingularTable: false,
			},
		})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	start := time.Now()
	for sqlDB.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("Failed to connect database after 10 seconds")
			break
		}
	}
	fmt.Println("Connected to database: ", sqlDB.Ping() == nil)

	db.AutoMigrate(&models.User{})

	DB = db

}
