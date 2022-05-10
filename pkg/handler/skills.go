package handler

import (
	"net/http"

	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createSkill(c *gin.Context) {
	var dto m.SkillRequest

	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	uuid, err := h.services.Skills.CreateNewSkill(dto)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, CreateResponse{Uuid: uuid})
}

func (h *Handler) getAllSkills(c *gin.Context) {
	skills, err := h.services.Skills.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, skills)
}

func (h *Handler) findSkill(c *gin.Context) {
	uuid := c.Param("id")
	skill, err := h.services.Skills.FindOne(uuid)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, skill)
}
