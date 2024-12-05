package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FelipeSoft/uptime-guardian/internal/worker"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	godotenv.Load("../../.env")
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %s", err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("ping", false, false, false, false, nil)

	if err != nil {
		log.Fatalf("declaring queue error: %s", err.Error())
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	if err != nil {
		log.Fatalf("consuming queue error: %s", err.Error())
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			res, err := worker.TestByICMP(string(d.Body))
			if err != nil {
				log.Fatalf("Error on ICMP test: %s", err.Error())
			}
			fmt.Printf("Received: %v; Loss: %v; Sent: %v \n", res.PacketsRecv, res.PacketLoss, res.PacketsSent)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-sigchan

	log.Printf("interrupted, shutting down")
	forever <- true
}
