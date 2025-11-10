package queue

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"api_busquedas/cache"
	"api_busquedas/clients"
	"api_busquedas/search"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/vanng822/go-solr/solr"
)

type ActividadEvent struct {
	Operation   string `json:"operation"` // create, update, delete
	ActividadID string `json:"actividad_id"`
	Timestamp   string `json:"timestamp"`
}

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

// StartConsumer starts consuming messages from RabbitMQ
func StartConsumer() error {
	rabbitURL := getEnv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/")
	queueName := getEnv("QUEUE_NAME", "actividades_events")

	log.Infof("Connecting to RabbitMQ at %s", rabbitURL)

	var err error

	// Retry connection with exponential backoff
	for i := 0; i < 5; i++ {
		conn, err = amqp.Dial(rabbitURL)
		if err == nil {
			break
		}
		log.Warnf("Failed to connect to RabbitMQ (attempt %d/5): %v", i+1, err)
		time.Sleep(time.Duration(i+1) * 2 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ after retries: %w", err)
	}

	channel, err = conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}

	// Declare queue (idempotent)
	queue, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	log.Infof("RabbitMQ connected. Listening on queue: %s", queueName)

	// Start consuming
	msgs, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	// Process messages
	for msg := range msgs {
		go processMessage(msg.Body)
	}

	return nil
}

func processMessage(body []byte) {
	var event ActividadEvent
	if err := json.Unmarshal(body, &event); err != nil {
		log.Errorf("Failed to unmarshal event: %v", err)
		return
	}

	log.Infof("Processing event: %s for actividad %s", event.Operation, event.ActividadID)

	switch event.Operation {
	case "create", "update":
		handleCreateOrUpdate(event.ActividadID)
	case "delete":
		handleDelete(event.ActividadID)
	default:
		log.Warnf("Unknown operation: %s", event.Operation)
	}
}

func handleCreateOrUpdate(actividadID string) {
	// 1. Fetch actividad from API_Actividades to ensure consistency
	actividad, err := clients.GetActividadFromAPI(actividadID)
	if err != nil {
		log.Errorf("Failed to fetch actividad %s from API: %v", actividadID, err)
		return
	}

	// 2. Index in Solr
	doc := solr.Document{
		"id":          actividadID,
		"nombre":      actividad.Nombre,
		"descripcion": actividad.Descripcion,
		"profesor":    actividad.Profesor,
		"tags":        actividad.Tags,
	}

	// Convert horarios to JSON string
	if len(actividad.Horarios) > 0 {
		horariosJSON, _ := json.Marshal(actividad.Horarios)
		doc["horarios"] = string(horariosJSON)
	}

	updateResp, err := search.SolrClient.Update(doc, nil)
	if err != nil {
		log.Errorf("Failed to index actividad %s in Solr: %v", actividadID, err)
		return
	}
	log.Debugf("Update response: %v", updateResp)

	// 3. Commit to Solr
	commitResp, err := search.SolrClient.Commit()
	if err != nil {
		log.Errorf("Failed to commit to Solr: %v", err)
		return
	}
	log.Debugf("Commit response: %v", commitResp)

	// 4. Invalidate cache
	cacheKey := fmt.Sprintf("actividad:%s", actividadID)
	cache.Delete(cacheKey)

	log.Infof("Successfully indexed actividad %s in Solr", actividadID)
}

func handleDelete(actividadID string) {
	// 1. Delete from Solr
	deleteResp, err := search.SolrClient.Delete(solr.M{"id": actividadID}, nil)
	if err != nil {
		log.Errorf("Failed to delete actividad %s from Solr: %v", actividadID, err)
		return
	}
	log.Debugf("Delete response: %v", deleteResp)

	// 2. Commit to Solr
	commitResp, err := search.SolrClient.Commit()
	if err != nil {
		log.Errorf("Failed to commit to Solr: %v", err)
		return
	}
	log.Debugf("Commit response: %v", commitResp)

	// 3. Invalidate cache
	cacheKey := fmt.Sprintf("actividad:%s", actividadID)
	cache.Delete(cacheKey)

	log.Infof("Successfully deleted actividad %s from Solr", actividadID)
}

// Close closes the RabbitMQ connection
func Close() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
