package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:"message"`
}

func NewErrorResponse(g *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	g.AbortWithStatusJSON(statusCode, error{message})
}

type CreateUserResponse struct {
	Uuid string `json:"uuid"`
}
