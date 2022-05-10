package handler

import (
	"net/http"

	"github.com/CranesTeam/team-screener/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) singUp(c *gin.Context) {
	var user model.UserDto

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logrus.Infof("Create new user: %s", user)

	uuid, err := h.services.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, CreateResponse{Uuid: uuid})
}

func (h *Handler) singIn(c *gin.Context) {
	var user model.UserAuthRequest

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logrus.Infof("Generate user token for user:%s", user)

	token, err := h.services.GenerateToken(user.Username, user.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, AuthTokenResponse{Token: token})
}
