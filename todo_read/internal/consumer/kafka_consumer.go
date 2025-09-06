package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

// Todo struct (same shape as your table)
type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Completed   bool   `json:"completed"`
}

// DebeziumEvent represents the outer wrapper
type DebeziumEvent struct {
	Payload struct {
		Before *Todo  `json:"before"`
		After  *Todo  `json:"after"`
		Op     string `json:"op"`
	} `json:"payload"`
}
type KafkaConsumer struct {
	// inject dependencies here if needed
	database  *mongo.Database
	brokerUrl *string
}

func NewKafkaConsumer(dbClient *mongo.Database, brokerUrl *string) *KafkaConsumer {
	return &KafkaConsumer{
		database:  dbClient,
		brokerUrl: brokerUrl,
	}
}

func (kc *KafkaConsumer) ConsumeDebeziumTodoTask() {
	// implement the logic to consume messages from Kafka and process them
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{*kc.brokerUrl},
		Topic:   "postgres.public.todos",
		GroupID: "gin-service",
	})
	collection := kc.database.Collection("todos")
	log.Println("Kafka consumer started...")
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		var event DebeziumEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		// Only act on INSERT or UPDATE
		if event.Payload.After != nil {
			todo := *event.Payload.After
			_, err := collection.InsertOne(context.Background(), todo)
			if err != nil {
				log.Printf("Failed to insert todo into MongoDB: %v", err)
			} else {
				fmt.Printf("Inserted todo: %+v\n", todo)
			}
		}
	}
}
