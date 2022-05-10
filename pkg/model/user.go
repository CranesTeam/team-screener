package model

type User struct {
	Id           int    `db:"id"`
	ExternalUuid string `db:"external_uuid"`
	Name         string `db:"name"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

type UserDto struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
