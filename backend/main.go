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
	router.GET("/event_types", controller.GetAllEventTypes)
	router.GET("/events/event_type/:name", controller.GetAllEventsFromEventType)
	router.GET("/events", controller.GetAllEvents)
	router.GET("/files/song/:id", controller.GetAllFilesFromSong)
	// TODO: router.GET("/files/song/text", )
	router.GET("/songs/event/:id", controller.GetAllSongsFromEvent)
	router.GET("/songs/subcategory/:id", controller.GetAllSongsFromSubcategory)
	router.GET("/songs", controller.GetAllSongs)
	router.GET("/subcategories/song/:id", controller.GetSongSubCategories)
	router.GET("/subcategories", controller.GetAllSubcategories)

	router.POST("/subcategory", controller.CreateSubCategory)
	router.POST("/event", controller.CreateEvent)
	router.POST("/song", controller.CreateSong)
	// TODO: router.POST("/file", controller.CreateFile)
	return router
}
