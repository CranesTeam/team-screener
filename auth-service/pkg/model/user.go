package model

type UserRequest struct {
	Sercet string `json:"secret"  binding:"required"`
	Domain string `json:"domain"  binding:"required"`
	UserId string `json:"userId"  binding:"required"`
}
