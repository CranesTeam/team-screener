package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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

		userSkills := api.Group(":id/skills")
		{
			userSkills.GET("/", h.getAllSkill)
			userSkills.POST("/", h.addSkill)
			userSkills.GET("/:skill_id", h.getSkillById)
			userSkills.PUT("/:skill_id", h.updateSkills)
			userSkills.DELETE("/:skill_id", h.deteleSkill)
		}

		service := api.Group("/health")
		{
			service.GET("/", h.health)
		}
	}

	return router
}

// func InitHandlers(s *product.ProductService) *http.Server {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.Use(otelmux.Middleware("my-api"))
// 	xraySegment := xray.NewFixedSegmentNamer("aws-go-service")

// 	myRouter.Handle("/api/health", xray.Handler(xraySegment,
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})
// 		})))

// 	myRouter.Handle("/product/all", xray.Handler(xraySegment, http.HandlerFunc(s.FindAll)))
// 	myRouter.Handle("/product/{id}", xray.Handler(xraySegment, http.HandlerFunc(s.FindOne)))
// 	myRouter.Handle("/product/{id}/add", xray.Handler(xraySegment, http.HandlerFunc(s.Create)))

// 	srv := &http.Server{
// 		Handler:      myRouter,
// 		Addr:         "127.0.0.1:8083",
// 		WriteTimeout: 15 * time.Second,
// 		ReadTimeout:  15 * time.Second,
// 	}

// 	log.Printf("Server will be started with address: %s", srv.Addr)
// 	return srv
// }
