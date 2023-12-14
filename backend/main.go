package main

import (
	"log"
	"os"

	"backend/controller"
	"backend/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	conn := db.Connect()
	defer conn.Close()

	port := os.Getenv("BACKEND_PORT")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
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
	}))

	setRoutes(router)
	router.Run("localhost:" + port)
}

func loadEnv() {
	err := godotenv.Load("../.env.local")
	if err != nil {
		log.Fatal(err)
	}
}

func setRoutes(router *gin.Engine) {
	router.GET("/categories", controller.GetAllCategories)
	router.GET("/event_types", controller.GetAllEventTypes)
	router.GET("/events/event_type/:name", controller.GetAllEventsFromEventType)
	router.GET("/events", controller.GetAllEvents)
	router.GET("/files/song/:id", controller.GetAllFilesInformationFromSong)
	router.GET("/files/song/text/:id", controller.DownloadTextFile)
	router.GET("/songs/event/:id", controller.GetAllSongsFromEvent)
	router.GET("/songs/subcategory/:id", controller.GetAllSongsFromSubcategory)
	router.GET("/songs", controller.GetAllSongs)
	router.GET("/subcategories/song/:id", controller.GetSongSubCategories)
	router.GET("/subcategories", controller.GetAllSubcategories)

	router.POST("/subcategory", controller.CreateSubCategory)
	router.POST("/event", controller.CreateEvent)
	router.POST("/song", controller.CreateSong)
	// TODO: router.POST("/file", controller.CreateFile)
}
