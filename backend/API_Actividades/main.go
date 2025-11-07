package main

import (
	"api_actividades/repositories/actividades"
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"api_actividades/app"
	db "api_actividades/db" // Importar como 'db' para facilitar el uso
)

func main() {
	// --- 1. CONFIGURACIÓN DE LA CONEXIÓN ---
	// La URI debe leerse de una variable de entorno para Docker
	cfg := db.MongoConfig{
		URI:        os.Getenv("MONGO_URI"),          // mongodb://mongouser:mongopass@mongodb:27017
		Database:   os.Getenv("MONGO_DB_NAME"),       // actividades_db
		TimeoutSec: 10,
	}

	// Si no se puede obtener la URI de las variables de entorno, salir.
	if cfg.URI == "" {
		log.Fatal("FATAL: MONGO_URI environment variable not set.")
	}

	// --- 2. INICIALIZAR Y CONECTAR LA BASE DE DATOS ---
	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("FATAL: Failed to connect to MongoDB: %v", err)
	}
	// Inicializar la cache local
	repository.NewCache(repository.CacheConfig{
		MaxSize:      100000,
		ItemsToPrune: 100,
		Duration:     30 * time.Minute,
	})
	log.Info("INFO: Successfully connected to Local Cache.")

	// --- 3. CIERRE DE LA CONEXIÓN ---
	// Usar defer para asegurar que la conexión se cierre al finalizar main()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = client.Disconnect(ctx); err != nil {
			log.Printf("WARNING: Failed to disconnect from MongoDB: %v", err)
		} else {
			log.Println("Database connection closed gracefully.")
		}
	}()
	
	// --- 4. CORRER MIGRACIONES / SEED ---
	// Correr el setup de la colección y el seed.
	migrationCtx, migrationCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer migrationCancel()
	if err := db.RunMigrations(migrationCtx, database); err != nil {
		log.Fatalf("FATAL: Failed to run database migrations/seed: %v", err)
	}
	log.Info("INFO: Database migrations/seed completed successfully.")

	// --- 5. INICIAR EL ROUTER Y LA API ---
	// Aquí debes pasar la instancia 'database' al resto de tu aplicación
	// (clients/repositories) para que puedan usarla.
	repository.Db = database
	
    // **PENDIENTE**: Necesitas una forma de inyectar 'database' al resto de la aplicación.
	// Por ahora, solo iniciamos las rutas.
	app.StartRoute()
}