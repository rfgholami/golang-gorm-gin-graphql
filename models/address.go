package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	PostalCode string `json:"postalCode"`
	PersonID   int    `json:"-"`
	Person     Person `json:"person" gorm:"foreignKey:PersonID;references:ID"`
}
