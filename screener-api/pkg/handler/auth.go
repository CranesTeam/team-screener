package handler

import (
	"net/http"

	"github.com/CranesTeam/team-screener/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary      Sing Up
// @Description  User sing up method
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body model.UserDto true "account info"
// @Success      200  {object}  CreateResponse
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /auth/sing-up [POST]
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

// @Summary      Sing In
// @Description  User sing in method
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body model.TokenResponse true "user token info"
// @Success      200  {object}  CreateResponse
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /auth/sing-in [POST]
func (h *Handler) singIn(c *gin.Context) {
	var user model.UserAuthRequest
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logrus.Infof("Generate user token for user:%s", user)

	token, err := h.services.GenerateJWT(user.Username, user.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}
