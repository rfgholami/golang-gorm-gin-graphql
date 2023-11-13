package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

/*func (User) TableName() string {
	return "go.users"
}
*/
