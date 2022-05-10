package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(g *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	g.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type CreateResponse struct {
	Uuid string `json:"uuid"`
}

type AuthTokenResponse struct {
	Token string `json:"token"`
}
