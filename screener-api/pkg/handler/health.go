package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) health(c *gin.Context) {
	logrus.Info("Heath check request was reveived")
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
