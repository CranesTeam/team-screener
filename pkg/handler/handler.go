package handler

import (
	"github.com/CranesTeam/team-screener/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.POST("sing-in", h.singIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		skill := api.Group("/skill")
		{
			skill.POST("/", h.createSkill)
			skill.GET("/:id", h.findSkill)
		}
		skills := api.Group("/skills")
		{
			skills.GET("/", h.getAllSkills)
		}

		list := api.Group("/list")
		{
			list.GET("/", h.getUserSkillList)
			list.POST("/", h.addNewSkill)
			list.GET("/:skill_uuid", h.findUserSkill)
			list.POST("/:skill_uuid/:points", h.updateUserSkills)
			list.DELETE("/:skill_uuid", h.deteleUserSkill)
		}
		service := api.Group("/health")
		{
			service.GET("/", h.health)
		}
	}

	return router
}
