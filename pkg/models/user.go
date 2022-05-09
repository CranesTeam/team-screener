package models

type User struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Username string `db:"username"`
	Password string `db:password`
}

type UserDto struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:password`
}
