package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserSkillList(c *gin.Context) {
	uuid, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "user not found")
	}

	list, err := h.services.UserSkills.GetUserSkills(uuid.(string))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) addNewSkillPoint(c *gin.Context) {

}

func (h *Handler) addNewSkill(c *gin.Context) {

}

func (h *Handler) findUserSkill(c *gin.Context) {

}

func (h *Handler) updateUserSkills(c *gin.Context) {

}

func (h *Handler) deteleUserSkill(c *gin.Context) {

}

func (h *Handler) deleteSkill(c *gin.Context) {

}
