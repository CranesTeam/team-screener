package model

type UserAuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Role        string `json:"role"`
	TokenString string `json:"token"`
	Exptime     int64  `json:"exp_time"`
}

type UserRequest struct {
	Sercet string `json:"secret"  binding:"required"`
	Domain string `json:"domain"  binding:"required"`
	UserId string `json:"userId"  binding:"required"`
}
