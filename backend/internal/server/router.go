package server

import (
	"net/http"

	"github.com/acmota2/musmgr/backend/internal/config"
	"github.com/acmota2/musmgr/backend/internal/controller"
	"github.com/acmota2/musmgr/backend/internal/controller/composer"
	enumscontroller "github.com/acmota2/musmgr/backend/internal/controller/enums"
	"github.com/acmota2/musmgr/backend/internal/controller/events"
	filescontroller "github.com/acmota2/musmgr/backend/internal/controller/files"
	"github.com/acmota2/musmgr/backend/internal/controller/pieces"
	piecesevent "github.com/acmota2/musmgr/backend/internal/controller/pieces_event"
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

func setPublicRoutes(router *gin.Engine, baseHandler *controller.BaseHandler, filesHandler *controller.FilesHandler, fileGetter *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	router.GET("/composer", middleware.RequirePerm(policies.PermList), composer.GetComposer(baseHandler))
	router.GET("/composer/picture", middleware.RequirePerm(policies.PermList), composer.GetComposerPicture(filesHandler, baseHandler))

	router.GET("/event_types", middleware.RequirePerm(policies.PermList), enumscontroller.GetEventTypes(baseHandler))
	router.GET("/file_types", middleware.RequirePerm(policies.PermList), enumscontroller.GetFileTypes(baseHandler))
	router.GET("/instrumentation_names", middleware.RequirePerm(policies.PermList), enumscontroller.GetInstrumentationNames(baseHandler))

	router.GET("/events", middleware.RequirePerm(policies.PermList), events.GetEvents(baseHandler))
	router.GET("/events/:event_id", middleware.RequirePerm(policies.PermList), events.GetEvent(baseHandler))
	router.GET("/events/:event_id/pieces", middleware.RequirePerm(policies.PermList), events.GetEventPieces(baseHandler))
	router.GET("/pieces", middleware.RequirePerm(policies.PermList), pieces.GetPieces(baseHandler))
	router.GET("/pieces/:piece_id", middleware.RequirePerm(policies.PermList), pieces.GetPiece(baseHandler))
	router.GET("/pieces/:piece_id/events", middleware.RequirePerm(policies.PermList), pieces.GetPieceEvents(baseHandler))
	router.GET("/pieces/:piece_id/files", middleware.RequirePerm(policies.PermList), filescontroller.GetPieceFiles(baseHandler))

	fileGetter.GET("", middleware.RequirePerm(policies.PermGet), filescontroller.GetFile(filesHandler, baseHandler))
}

func NewPublicRouter(cfg *config.Config, handler *controller.BaseHandler, filesHandler *controller.FilesHandler) *gin.Engine {
	router := baseRouter(cfg)

	router.Use(cors.New(cors.Config{
		AllowOrigins: cfg.PublicRoutes,
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Content-Type"},
	}))

	router.Use(middleware.SetPublicRouterScope(), middleware.SetPublicRouterClass())
	fileGetter := router.Group("/pieces/:piece_id/files/:file_id", middleware.FileClassificationBlocking(handler.Queries))

	setPublicRoutes(router, handler, filesHandler, fileGetter)

	return router
}

func setAdminOnlyRoutes(router *gin.Engine, handler *controller.BaseHandler, filesHandler *controller.FilesHandler, fileGetter *gin.RouterGroup) {
	router.POST("/composer", middleware.RequirePerm(policies.PermWrite), composer.CreateComposer(handler))
	router.POST("/events", middleware.RequirePerm(policies.PermWrite), events.CreateEvent(handler))
	router.POST("/pieces/:piece_id/events/:event_id", middleware.RequirePerm(policies.PermWrite), piecesevent.CreatePieceEvent(handler))
	router.POST("/pieces", middleware.RequirePerm(policies.PermWrite), pieces.CreatePiece(handler))
	router.POST("/pieces/:piece_id/files", middleware.RequirePerm(policies.PermWrite), filescontroller.CreateFile(filesHandler, handler))

	router.PATCH("/composer", middleware.RequirePerm(policies.PermWrite), composer.UpdateComposer(handler))
	router.PATCH("/events/:event_id", middleware.RequirePerm(policies.PermWrite), events.UpdateEvent(handler))
	router.PATCH("/pieces/:piece_id", middleware.RequirePerm(policies.PermWrite), pieces.UpdatePiece(handler))
	router.PATCH("/pieces/:piece_id/files/:file_id", middleware.RequirePerm(policies.PermWrite), filescontroller.UpdateFileMetadata(handler))

	router.PUT("/composer/picture", middleware.RequirePerm(policies.PermWrite), composer.UpdateComposerPicture(filesHandler, handler))

	router.DELETE("/composer/picture", middleware.RequirePerm(policies.PermDelete), composer.DeleteComposerPicture(filesHandler, handler))
	router.DELETE("/pieces/:piece_id", middleware.RequirePerm(policies.PermDelete), pieces.DeletePiece(filesHandler, handler))
	router.DELETE("/events/:event_id", middleware.RequirePerm(policies.PermDelete), events.DeleteEvent(handler))

	fileGetter.DELETE("", middleware.RequirePerm(policies.PermDelete), filescontroller.DeleteFile(filesHandler, handler))
}

func NewAdminRouter(cfg *config.Config, handler *controller.BaseHandler, filesHandler *controller.FilesHandler) *gin.Engine {
	router := baseRouter(cfg)
	router.MaxMultipartMemory = 64 << 20 // 64MiB

	router.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AdminRoutes,
		AllowMethods: []string{"DELETE", "GET", "PATCH", "POST", "PUT"},
		AllowHeaders: []string{"Content-Type"},
	}))

	router.Use(middleware.SetAdminRouterScope(), middleware.SetAdminRouterClass())
	fileGetter := router.Group("/pieces/:piece_id/files/:file_id", middleware.FileClassificationBlocking(handler.Queries))

	setPublicRoutes(router, handler, filesHandler, fileGetter)
	setAdminOnlyRoutes(router, handler, filesHandler, fileGetter)

	return router
}
