package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID          string `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
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
		Topic:   "postgres.todos.todos",
		GroupID: "gin-service",
	})
	collection := kc.database.Collection("todos")
	log.Println("Kafka consumer started...")
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Consumed event: %s\n", msg.Value)
		var todo Todo
		json.Unmarshal(msg.Value, &todo)
		collection.InsertOne(context.Background(), todo)
		fmt.Printf("Consumed event: %s\n", msg.Value)
	}
}
