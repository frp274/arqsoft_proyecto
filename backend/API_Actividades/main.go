package main

import (
	"api_actividades/queue"
	repository "api_actividades/repositories/actividades"
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"api_actividades/app"
	"api_actividades/cache"
	db "api_actividades/db" // Importar como 'db' para facilitar el uso
)

func main() {
	// --- 1. CONFIGURACIÓN DE LA CONEXIÓN ---
	// La URI debe leerse de una variable de entorno para Docker
	cfg := db.MongoConfig{
		URI:        os.Getenv("MONGO_URI"),     // mongodb://mongouser:mongopass@mongodb:27017
		Database:   os.Getenv("MONGO_DB_NAME"), // actividades_db
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
	// Inicializar la doble capa de caché (L1 + L2)
	cache.InitCache()

	// --- 3. INICIALIZAR RABBITMQ PRODUCER ---
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@localhost:5672/" // Default para desarrollo local
	}
	queueName := os.Getenv("RABBITMQ_QUEUE")
	if queueName == "" {
		queueName = "actividades_events"
	}
	exchangeName := os.Getenv("RABBITMQ_EXCHANGE")
	if exchangeName == "" {
		exchangeName = "actividades_exchange"
	}

	// Reintentar la conexión a RabbitMQ con backoff exponencial
	// (RabbitMQ puede tardar unos segundos más que el contenedor en estar listo)
	var rabbitErr error
	for i := 0; i < 10; i++ {
		rabbitErr = queue.InitProducer(rabbitURL, queueName, exchangeName)
		if rabbitErr == nil {
			break
		}
		waitSec := time.Duration(i+1) * 2 * time.Second
		log.Warnf("WARNING: RabbitMQ no disponible (intento %d/10), reintentando en %v: %v", i+1, waitSec, rabbitErr)
		time.Sleep(waitSec)
	}
	if rabbitErr != nil {
		log.Fatalf("FATAL: No se pudo conectar a RabbitMQ después de 10 intentos: %v", rabbitErr)
	}
	log.Info("INFO: RabbitMQ producer initialized successfully.")

	// --- 4. CIERRE DE CONEXIONES ---
	// Usar defer para asegurar que las conexiones se cierren al finalizar main()
	defer func() {
		// Cerrar RabbitMQ
		if err := queue.Close(); err != nil {
			log.Printf("WARNING: Failed to close RabbitMQ connection: %v", err)
		}

		// Cerrar MongoDB
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = client.Disconnect(ctx); err != nil {
			log.Printf("WARNING: Failed to disconnect from MongoDB: %v", err)
		} else {
			log.Println("Database connection closed gracefully.")
		}
	}()

	// --- 5. CORRER MIGRACIONES / SEED ---
	// RunMigrations retorna los IDs de TODAS las actividades (nuevas o existentes).
	// Los publicamos en RabbitMQ para que API_Busquedas los indexe en Solr.
	// Esto garantiza que Solr siempre esté sincronizado, incluso si el contenedor
	// se reinició antes de que RabbitMQ estuviera listo la primera vez.
	migrationCtx, migrationCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer migrationCancel()
	actividadIDs, err := db.RunMigrations(migrationCtx, database)
	if err != nil {
		log.Fatalf("FATAL: Failed to run database migrations/seed: %v", err)
	}
	log.Infof("INFO: Database migrations/seed completed. Publicando %d actividades en RabbitMQ para sincronizar Solr...", len(actividadIDs))

	for _, id := range actividadIDs {
		if pubErr := queue.PublishEvent(queue.EventCreate, id); pubErr != nil {
			log.Warnf("WARNING: No se pudo publicar evento CREATE para actividad %s: %v", id, pubErr)
		}
	}
	log.Info("INFO: Sincronización con Solr via RabbitMQ completada.")

	// --- 6. INICIAR EL ROUTER Y LA API ---
	// Aquí debes pasar la instancia 'database' al resto de tu aplicación
	// (clients/repositories) para que puedan usarla.
	repository.Db = database

	// **PENDIENTE**: Necesitas una forma de inyectar 'database' al resto de la aplicación.
	// Por ahora, solo iniciamos las rutas.
	app.StartRoute()
}
