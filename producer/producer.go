package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("Error: %s: %s", msg, err)
	}
}

func main() {
	var (
		bodyPtr       = flag.String("body", "", "message body")
		exchangePtr   = flag.String("exchange", "", "post message to this exchange")
		routingkeyPtr = flag.String("routingkey", "", "post message with this routing key")
		serverPtr     = flag.String("server", "localhost:5672", "post message to this AMQP server")
		userPtr       = flag.String("user", "guest", "rabbitmq user")
		passwordPtr   = flag.String("password", "guest", "rabbitmq password")
	)

	flag.Parse()

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", *userPtr, *passwordPtr, *serverPtr))
	log.Printf("Connecting to %s", *serverPtr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	if *exchangePtr == "" {
		log.Fatalf("Error: Exchange required")
	}

	err = ch.Publish(
		*exchangePtr,   // exchange
		*routingkeyPtr, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(*bodyPtr),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent \"%s\"", *bodyPtr)
}
