package models

type Request struct {
	Entity string      `json:"entity"`
	Data   interface{} `json:"data"`
}
