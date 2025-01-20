package main

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	emailConsumer, err := channel.ConsumeWithContext(ctx, "email", "consumer-email", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for message := range emailConsumer {
		log.Printf("Routing Key: %s", string(message.RoutingKey))
		log.Printf("Body: %s", message.Body)
	}
}
