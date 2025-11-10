package main

import (
	"api_usuarios/app"
	usuarioClient "api_usuarios/clients/usuarios"
	"api_usuarios/db"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting API_Usuarios microservice...")

	// Initialize the database connection
	database := db.InitConnection()

	// Initialize usuario client with database
	usuarioClient.Db = database

	// Start the HTTP server
	app.StartRoute()

	// Close the database connection when done
	defer db.Close(database)
	log.Info("API_Usuarios shutdown complete")
}
