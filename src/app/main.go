package main

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"

	"github.com/web-service-gin/src/controllers"
	"github.com/web-service-gin/src/interfaces"
	"github.com/web-service-gin/src/util"
)

var db *sql.DB

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	util.ConnectDb()

	// load and init controller
	albumController := controllers.NewAlbumController()
	userController := controllers.NewUserController()
	controllers := []interfaces.ControllerInterface{
		albumController,
		userController,
	}
	app := newApp(controllers, "localhost:8080")
	app.run()
}
