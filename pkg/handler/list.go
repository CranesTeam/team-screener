package handler

import (
	"net/http"

	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserSkillList(c *gin.Context) {
	uuid, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "user not found")
		return
	}

	list, err := h.services.UserSkills.GetUserSkills(uuid.(string))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) addNewSkill(c *gin.Context) {
	user_uuid, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "user not found")
		return
	}

	var dto m.AddSkillRequest
	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	uuid, err := h.services.UserSkills.AddNewSkillPointer(user_uuid.(string), dto)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, CreateResponse{Uuid: uuid})
}

func (h *Handler) findUserSkill(c *gin.Context) {

}

func (h *Handler) updateUserSkills(c *gin.Context) {

}

func (h *Handler) deteleUserSkill(c *gin.Context) {

}

func (h *Handler) deleteSkill(c *gin.Context) {

}
