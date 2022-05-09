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

	api := router.Group("/api")
	{
		skills := api.Group("/skills")
		{
			skills.POST("/", h.createSkill)
			skills.GET("/", h.getAllSkills)
			skills.GET("/:id", h.getOneSkill)
		}

		// TODO:
		// userSkills := api.Group(":id/skills")
		// {
		// 	userSkills.GET("/", h.getAllSkill)
		// 	userSkills.POST("/", h.addSkill)
		// 	userSkills.GET("/:skill_id", h.getSkillById)
		// 	userSkills.PUT("/:skill_id", h.updateSkills)
		// 	userSkills.DELETE("/:skill_id", h.deteleSkill)
		// }

		service := api.Group("/health")
		{
			service.GET("/", h.health)
		}
	}

	return router
}
