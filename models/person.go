package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	ID           int       `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	LastName     string    `json:"lastName"`
	MobileNumber string    `json:"mobileNumber"`
	AddressItems []Address `json:"addressItems"`
}
