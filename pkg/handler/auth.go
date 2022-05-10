package handler

import (
	"net/http"

	"github.com/CranesTeam/team-screener/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) singUp(c *gin.Context) {
	var user model.UserDto

	if err := c.BindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	uuid, err := h.services.CreateUser(user)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, CreateUserResponse{Uuid: uuid})
}

func (h *Handler) singIn(c *gin.Context) {
	var user model.UserAuthRequest

	if err := c.BindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(user.Username, user.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, AuthTokenResponse{Token: token})
}
