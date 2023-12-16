package models

type PersonDto struct {
	ID           int    `json:"ID"`
	Name         string `json:"name"`
	LastName     string `json:"lastName"`
	MobileNumber string `json:"mobileNumber"`
	AddressId    string `json:"address_id"`
	AddressTitle string `json:"address_Title"`
}
