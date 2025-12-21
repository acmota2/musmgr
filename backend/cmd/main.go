package main

import (
	"backend/internal/config"
	"backend/internal/controller"
	"backend/internal/model"
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func setRoutes(router *gin.Engine, cc *controller.ControllerContext) {
	router.GET("/works/:id/files", cc.GetWorkFiles)
	router.GET("/events/:id/works", cc.GetEventWorks)
	router.GET("/works", cc.GetWorks)

	router.POST("/events/:event_id/works/:work_id", cc.CreateWorkEvent)
	router.POST("/event_types/:type/events", cc.CreateEvent)
	router.POST("/works", cc.CreateEvent)
}

func main() {
	initConfig, err := config.CreateConfig()
	if err != nil {
		log.Panic(err)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, initConfig.DatabaseUrl)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close(ctx)
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

	cc := controller.New(model.New(conn), ctx)
	setRoutes(router, cc)
	router.Run(initConfig.PortHost)
}
