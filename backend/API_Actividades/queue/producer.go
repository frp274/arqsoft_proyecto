package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

// EventType representa el tipo de evento
type EventType string

const (
	EventCreate EventType = "create"
	EventUpdate EventType = "update"
	EventDelete EventType = "delete"
)

// ActividadEvent es el mensaje que se envía a RabbitMQ
type ActividadEvent struct {
	Type         EventType `json:"type"`
	ActividadID  string    `json:"actividad_id"`
	Timestamp    time.Time `json:"timestamp"`
}

// RabbitMQProducer gestiona la conexión y publicación de mensajes
type RabbitMQProducer struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	queueName    string
	exchangeName string
}

var producer *RabbitMQProducer

// InitProducer inicializa la conexión con RabbitMQ y declara la cola
func InitProducer(rabbitURL, queueName, exchangeName string) error {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return fmt.Errorf("failed to open channel: %w", err)
	}

	// Declarar el exchange (tipo fanout para broadcasting)
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Declarar la cola
	_, err = ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Bind la cola al exchange
	err = ch.QueueBind(
		queueName,    // queue name
		"",           // routing key (vacío para fanout)
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	producer = &RabbitMQProducer{
		conn:         conn,
		channel:      ch,
		queueName:    queueName,
		exchangeName: exchangeName,
	}

	log.Infof("RabbitMQ producer initialized successfully on exchange '%s' and queue '%s'", exchangeName, queueName)
	return nil
}

// PublishEvent publica un evento de actividad en RabbitMQ
func PublishEvent(eventType EventType, actividadID string) error {
	if producer == nil {
		return fmt.Errorf("producer not initialized")
	}

	event := ActividadEvent{
		Type:        eventType,
		ActividadID: actividadID,
		Timestamp:   time.Now(),
	}

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = producer.channel.PublishWithContext(
		ctx,
		producer.exchangeName, // exchange
		"",                    // routing key (vacío para fanout)
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent, // Persistente para que sobreviva a reinicio de RabbitMQ
			Timestamp:    time.Now(),
		},
	)

	if err != nil {
		log.Errorf("Failed to publish event %s for actividad %s: %v", eventType, actividadID, err)
		return fmt.Errorf("failed to publish event: %w", err)
	}

	log.Infof("Event published: %s for actividad %s", eventType, actividadID)
	return nil
}

// Close cierra la conexión con RabbitMQ
func Close() error {
	if producer == nil {
		return nil
	}

	if producer.channel != nil {
		if err := producer.channel.Close(); err != nil {
			log.Errorf("Error closing channel: %v", err)
		}
	}

	if producer.conn != nil {
		if err := producer.conn.Close(); err != nil {
			log.Errorf("Error closing connection: %v", err)
			return err
		}
	}

	log.Info("RabbitMQ producer closed successfully")
	return nil
}
