package main

import (
	"backend/internal/config"
	"backend/internal/controller"
	"backend/internal/model"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func setRoutes(router *gin.Engine, h *controller.Handler) {
	router.GET("/works/:id/files", h.GetWorkFiles)
	router.GET("/events/:id/works", h.GetEventWorks)
	router.GET("/works", h.GetWorks)

	router.POST("/events/:event_id/works/:work_id", h.CreateWorkEvent)
	router.POST("/event_types/:type/events", h.CreateEvent)
	router.POST("/works", h.CreateEvent)
}

func main() {
	initConfig, err := config.New()
	if err != nil {
		log.Panic(err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	conn, err := pgxpool.New(ctx, initConfig.DatabaseUrl)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to the database")

	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, // TODO: Change to env variable
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Authorization",
			"Origin",
			"Accept",
		},
		ExposeHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	handler := &controller.Handler{
		Queries: model.New(conn),
	}
	setRoutes(router, handler)
	go router.Run(initConfig.PortHost)
	<-ctx.Done()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")
}
