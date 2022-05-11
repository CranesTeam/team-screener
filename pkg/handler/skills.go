package handler

import (
	"net/http"

	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary Create user skill
// @Security ApiKeyAuth
// @Tags skills
// @Description create user skill
// @ID create-skill
// @Accept  json
// @Produce  json
// @Param input body model.SkillRequest true "skill info"
// @Success 200 {object} CreateResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/skill [post]
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

// @Summary Get All Skills
// @Security ApiKeyAuth
// @Tags skills
// @Description get all skills
// @ID get-all-skills
// @Accept  json
// @Produce  json
// @Success 200 {array} model.SkillDto
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/skills [get]
func (h *Handler) getAllSkills(c *gin.Context) {
	skills, err := h.services.Skills.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, skills)
}

// @Summary Get skill by uuid
// @Security ApiKeyAuth
// @Tags skills
// @Description get skill by uuid
// @ID get-list-skill-uuid
// @Accept  json
// @Produce  json
// @Success 200 {object} model.SkillDto
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/skill/:uuid [get]
func (h *Handler) findSkill(c *gin.Context) {
	uuid := c.Param("id")
	skill, err := h.services.Skills.FindOne(uuid)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, skill)
}
