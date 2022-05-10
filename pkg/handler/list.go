package handler

import (
	"net/http"

	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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
