package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userUuid"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" { // todo: fixed it
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userUuid, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userUuid)
}

// func getUserUuid(c *gin.Context) (string, error) {
// 	uuid, ok := c.Get(userCtx)
// 	if !ok {
// 		return "empty", errors.New("user id not found")
// 	}

// 	externalUuid, ok := uuid.(string)
// 	if !ok {
// 		return "empty", errors.New("user id is of invalid type")
// 	}

// 	return externalUuid, nil
// }
