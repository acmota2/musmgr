package main

import (
	"backend/controller"
	"backend/model"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load("./.env.local")
	if err != nil {
		log.Fatal(err)
	}
}

func setRoutes(router *gin.Engine, cc *controller.ControllerContext) {
	router.GET("/works/:id/files", cc.GetWorkFiles)
	router.GET("/events/:id/works", cc.GetEventWorks)
	router.GET("/works", cc.GetWorks)

	// router.POST("/files/song", controller.CreateWorkFile)  TODO: Create me, on version 2.0
	router.POST("/events/:event_id/works/:work_id", cc.CreateWorkEvent)
	router.POST("/event_types/:id/events", cc.CreateEvent)
	router.POST("/works", cc.CreateEvent)
}

func main() {
	loadEnv()

	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close(ctx)
	fmt.Println("Successfully connected to the database")

	port := os.Getenv("BACKEND_PORT")
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
	backendHost := os.Getenv("BACKEND_HOST")
	router.Run(backendHost + ":" + port)
}
