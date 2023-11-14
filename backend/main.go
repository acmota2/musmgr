package main

import (
	"log"
	"os"

	"backend/controller"
	"backend/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	conn := db.Connect()
	defer conn.Close()

	router := setRoutes()

	port := os.Getenv("BACKEND_PORT")
	router.Run("localhost:" + port)
}

func loadEnv() {
	err := godotenv.Load("../.env.local")
	if err != nil {
		log.Fatal(err)
	}
}

func setRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/categories", controller.GetAllCategories)
	router.GET("/events", controller.GetAllEvents)
	router.GET("/songs", controller.SongCategories)
	router.GET("/songs/event/:id", controller.GetAllSongsFromEvent)

	router.POST("/subcategory", controller.CreateSubCategory)
	router.POST("/event", controller.CreateEvent)
	router.POST("/song", controller.CreateSong)
	return router
}
