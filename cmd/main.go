package main

import (
	"app/config"
	"app/db"
	"app/docs"
	"app/server"
	"app/server/controller"

	log "github.com/sirupsen/logrus"
)

// @title           AuthServiceAPI
// @version         1.0
// @description     Car auction app
// @termsOfService  http://swagger.io/terms/

// @BasePath /
func main() {
	docs.SwaggerInfo.Host = ""

	config, err := config.InnitConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	dbConnection, err := db.NewConnection(config)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer dbConnection.Close()

	_, err = dbConnection.RunQueryFromFile(config.GetMigrationPath())

	if err != nil {
		log.Fatalf("Migration failed to up: %s", err.Error())
	}

	controller := controller.NewController(dbConnection)

	server := server.NewServer(config, controller)

	server.Start()
}
