package main

import (
	"context"
	"fmt"
	"log"
	"time"

	db_connection "github.com/georgehyde-dot/mallard-migrations/pkg/dbconnection"
	queue_integration "github.com/georgehyde-dot/mallard-migrations/pkg/queueintegration"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	log.Println("Migrating Your Data, Quack Quack")

	// Initialize RabbitMQ
	var mq queue_integration.MessageQueue = &queue_integration.RabbitMQ{}
	err := mq.Connect("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer mq.Close()

	// Publish a message
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mq.Publish(ctx, "hello", "Hello, World!")
	failOnError(err, "Failed to publish a message")

	// Consume messages
	msgs, err := mq.Consume("hello")
	failOnError(err, "Failed to register a consumer")

	// execute msg on db
	dbUrl := "test@test"
	conn, err := db_connection.ConnectToDB(db_connection.MSSQL, dbUrl)
	failOnError(err, "Failed to connect to db")

	for msg := range msgs {
		startTime := time.Now()
		result, err := conn.Exec(ctx, msg)
		if err != nil {
			fmt.Printf("failed to execute %v", msg)
		}
		duration := time.Since(startTime)
		fmt.Printf("Response to message %v: %v, duration: %d", msg, result, duration)
	}

}

// TODO list

// parse message

// write to migrations DB (SQLite?, or on DB that sent the request is going to?)

// add psotgres switch for DB type

// validate SQL?
