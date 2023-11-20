package models

type LoginRequest struct {
	Captcha      string `json:"captcha"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Password     string `json:"password"`
	Username     string `json:"username"`
}
