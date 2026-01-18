package server

import (
	"backend/internal/config"
	"backend/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func baseRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func setPublicRoutes(router *gin.Engine, handler *controller.Handler) {
	router.GET("/events", handler.GetEvents)
	router.GET("/events/:event_id/works", handler.GetEventWorks)
	router.GET("/works", handler.GetWorks)
	router.GET("/works/:work_id/files", handler.GetWorkFiles)
	router.GET("/works/:work_id/events", handler.GetWorkEvents)
}

func NewPublicRouter(cfg *config.Config, handler *controller.Handler) *gin.Engine {
	router := baseRouter()

	router.Use(cors.New(cors.Config{
		AllowOrigins: cfg.PublicRoutes,
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Content-Type"},
	}))

	setPublicRoutes(router, handler)

	return router
}

func setAdminOnlyRoutes(router *gin.Engine, handler *controller.Handler) {
	router.POST("/events", handler.CreateEvent)
	router.POST("/work_event", handler.CreateWorkEvent)
	router.POST("/works", handler.CreateEvent)
}

func NewAdminRouter(cfg *config.Config, handler *controller.Handler) *gin.Engine {
	router := baseRouter()

	router.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AdminRoutes,
		AllowMethods: []string{"DELETE", "GET", "PATCH", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))

	setPublicRoutes(router, handler)
	setAdminOnlyRoutes(router, handler)

	return router
}
