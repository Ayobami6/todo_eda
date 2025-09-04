package main

import (
	"log"
	"time"

	"context"

	"github.com/Ayobami6/todo_read/cmd/api"
	"github.com/Ayobami6/todo_read/config"
	"github.com/Ayobami6/todo_read/internal/consumer"
	"github.com/Ayobami6/webutils"
	_ "github.com/joho/godotenv/autoload"
)

// main entry point for the application

func main() {
	// context timeout
	dbUrl := webutils.GetEnv("MONGO_URI", "mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	dbClient, err := config.ConnectDb(ctx, dbUrl)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	kafkaBrokerUrl := webutils.GetEnv("KAFKA_BROKER_URL", "localhost:9092")
	kafkaConsumer := consumer.NewKafkaConsumer(dbClient.Database("todos"), &kafkaBrokerUrl)
	go kafkaConsumer.ConsumeDebeziumTodoTask()
	apiServer := api.NewAPIServer(":8080", dbClient)
	if err := apiServer.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
