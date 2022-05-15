package handler

import (
	"net/http"

	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Get all user skills
// @Security ApiKeyAuth
// @Tags user-skills
// @Description get all user skills
// @ID get-all-user-skills
// @Accept  json
// @Produce  json
// @Success 200 {array} model.SkillListDto
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list [get]
func (h *Handler) getUserSkillList(c *gin.Context) {
	user_uuid, err := getUserUuid(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Infof("Get all skill by user_uuid:%s", user_uuid)

	list, err := h.services.UserSkills.GetUserSkills(user_uuid)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Add new user skill
// @Security ApiKeyAuth
// @Tags user-skills
// @Description add new user skill
// @ID create-user-skill
// @Accept  json
// @Produce  json
// @Param input body model.AddSkillRequest true "skill info"
// @Success 200 {object} CreateResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/skill [post]
func (h *Handler) addNewSkill(c *gin.Context) {
	user_uuid, err := getUserUuid(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Infof("Add new skill by user_uuid:%s", user_uuid)

	var dto m.AddSkillRequest
	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	uuid, err := h.services.UserSkills.AddNewSkillPointer(user_uuid, dto)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, CreateResponse{Uuid: uuid})
}

// @Summary Find user skills by uuid
// @Security ApiKeyAuth
// @Tags user-skills
// @Description find user skill by uuid
// @ID find-one-user-skills
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserSkillsDto
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list/:skill_uuid [get]
func (h *Handler) findUserSkill(c *gin.Context) {
	user_uuid, err := getUserUuid(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Infof("Find skill by user_uuid:%s", user_uuid)

	uuid := c.Param("skill_uuid")
	skill, err := h.services.UserSkills.FindSkill(user_uuid, uuid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, skill)
}

// @Summary update user skill point by uuid
// @Security ApiKeyAuth
// @Tags user-skills
// @Description update user skill point by uuid
// @ID update-one-user-skills
// @Accept  json
// @Produce  json
// @Success 200 {object} CreateResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list/:skill_uuid/:points [post]
func (h *Handler) updateUserSkills(c *gin.Context) {
	user_uuid, err := getUserUuid(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Infof("Find skill by user_uuid:%s", user_uuid)

	skill_uuid := c.Param("skill_uuid")
	points := c.Param("points")

	uuid, err := h.services.UserSkills.UpdatePoint(user_uuid, skill_uuid, points)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, CreateResponse{Uuid: uuid})
}

// @Summary delete user skill by uuid
// @Security ApiKeyAuth
// @Tags user-skills
// @Description delete user skill by uuid
// @ID delete-one-user-skills
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserSkillsDto
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/list/:skill_uuid [delete]
func (h *Handler) deteleUserSkill(c *gin.Context) {
	user_uuid, err := getUserUuid(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Infof("Delete skill for user_uuid:%s", user_uuid)

	skill_uuid := c.Param("skill_uuid")
	uuid, err := h.services.UserSkills.DeleteSkill(user_uuid, skill_uuid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, CreateResponse{Uuid: uuid})
}
