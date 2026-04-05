package server

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/config"
	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/middleware"
	"github.com/acmota2/musmgr/backend/internal/policies"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func baseRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(cfg.TrustedProxies)
	return router
}

func setPublicRoutes(router *gin.Engine, handler *controller.Handler, fileGetter *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	router.GET("/composer", middleware.RequirePerm(policies.PermList), handler.GetComposer)
	router.GET("/composer/picture", middleware.RequirePerm(policies.PermList), handler.GetComposerPicture)

	router.GET("/event_types", middleware.RequirePerm(policies.PermList), handler.GetEventTypes)
	router.GET("/file_types", middleware.RequirePerm(policies.PermList), handler.GetFileTypes)
	router.GET("/instrumentation_names", middleware.RequirePerm(policies.PermList), handler.GetInstrumentationNames)

	router.GET("/events", middleware.RequirePerm(policies.PermList), handler.GetEvents)
	router.GET("/events/:event_id", middleware.RequirePerm(policies.PermList), handler.GetEvent)
	router.GET("/events/:event_id/pieces", middleware.RequirePerm(policies.PermList), handler.GetEventPieces)
	router.GET("/pieces", middleware.RequirePerm(policies.PermList), handler.GetPieces)
	router.GET("/pieces/:piece_id", middleware.RequirePerm(policies.PermList), handler.GetPiece)
	router.GET("/pieces/:piece_id/events", middleware.RequirePerm(policies.PermList), handler.GetPieceEvents)
	router.GET("/pieces/:piece_id/files", middleware.RequirePerm(policies.PermList), handler.GetPieceFiles)

	fileGetter.GET("", middleware.RequirePerm(policies.PermGet), handler.GetFile)
}

func NewPublicRouter(cfg *config.Config, handler *controller.Handler) *gin.Engine {
	router := baseRouter(cfg)

	router.Use(cors.New(cors.Config{
		AllowOrigins: cfg.PublicRoutes,
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Content-Type"},
	}))

	router.Use(middleware.SetPublicRouterScope(), middleware.SetPublicRouterClass())
	fileGetter := router.Group("/pieces/:piece_id/files/:file_id", middleware.FileClassificationBlocking(handler.Queries))

	setPublicRoutes(router, handler, fileGetter)

	return router
}

func setAdminOnlyRoutes(router *gin.Engine, handler *controller.Handler, fileGetter *gin.RouterGroup) {
	router.POST("/composer", middleware.RequirePerm(policies.PermWrite), handler.CreateComposer)
	router.POST("/events", middleware.RequirePerm(policies.PermWrite), handler.CreateEvent)
	router.POST("/pieces/:piece_id/events/:event_id", middleware.RequirePerm(policies.PermWrite), handler.CreatePieceEvent)
	router.POST("/pieces", middleware.RequirePerm(policies.PermWrite), handler.CreatePiece)
	router.POST("/pieces/:piece_id/files", middleware.RequirePerm(policies.PermWrite), handler.CreateFile)

	router.PATCH("/composer", middleware.RequirePerm(policies.PermWrite), handler.UpdateComposer)
	router.PATCH("/events/:event_id", middleware.RequirePerm(policies.PermWrite), handler.UpdateEvent)
	router.PATCH("/pieces/:piece_id", middleware.RequirePerm(policies.PermWrite), handler.UpdatePiece)

	router.PUT("/composer/picture", middleware.RequirePerm(policies.PermWrite), handler.UpdateComposerPicture)

	router.DELETE("/composer/picture", middleware.RequirePerm(policies.PermDelete), handler.DeleteComposerPicture)
	router.DELETE("/pieces/:piece_id", middleware.RequirePerm(policies.PermDelete), handler.DeletePiece)
	router.DELETE("/events/:event_id", middleware.RequirePerm(policies.PermDelete), handler.DeleteEvent)

	fileGetter.DELETE("", middleware.RequirePerm(policies.PermDelete), handler.DeleteFile)
}

func NewAdminRouter(cfg *config.Config, handler *controller.Handler) *gin.Engine {
	router := baseRouter(cfg)
	router.MaxMultipartMemory = 64 << 20 // 64MiB

	router.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AdminRoutes,
		AllowMethods: []string{"DELETE", "GET", "PATCH", "POST", "PUT"},
		AllowHeaders: []string{"Content-Type"},
	}))

	router.Use(middleware.SetAdminRouterScope(), middleware.SetAdminRouterClass())
	fileGetter := router.Group("/pieces/:piece_id/files/:file_id", middleware.FileClassificationBlocking(handler.Queries))

	setPublicRoutes(router, handler, fileGetter)
	setAdminOnlyRoutes(router, handler, fileGetter)

	return router
}
