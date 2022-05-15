package model

type UserInfo struct {
	Id           int    `db:"id"`
	ExternalUuid string `db:"external_uuid"`
	Name         string `db:"name"`
	Email        string `db:"email"`
}

type User struct {
	Id           int    `db:"id"`
	ExternalUuid string `db:"external_uuid"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
	RoleId       int    `db:"role_id"`
	RoleName     string `db:"role_name"`
}

type UserRoles struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type UserDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
