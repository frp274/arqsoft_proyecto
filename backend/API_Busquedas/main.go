package main

import (
	"api_busquedas/app"
	"api_busquedas/cache"
	"api_busquedas/db"
	"api_busquedas/queue"
	"api_busquedas/search"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting API_Busquedas microservice...")

	// Initialize MySQL database connection
	database := db.InitConnection()
	defer db.Close(database)
	log.Info("MySQL database connected and tables migrated")

	// Initialize Solr connection
	if err := search.InitSolr(); err != nil {
		log.Fatalf("Failed to initialize Solr: %v", err)
	}
	log.Info("Solr connection initialized")

	// Initialize cache layers (local + distributed)
	cache.InitCache()
	log.Info("Cache layers initialized")

	// Start RabbitMQ consumer in background
	go func() {
		if err := queue.StartConsumer(); err != nil {
			log.Errorf("RabbitMQ consumer error: %v", err)
		}
	}()
	log.Info("RabbitMQ consumer started")

	// Start HTTP server
	app.StartRoute()
}
