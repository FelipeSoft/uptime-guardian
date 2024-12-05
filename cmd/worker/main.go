package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/worker"
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

	icmpChannel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %s", err.Error())
	}
	defer icmpChannel.Close()

	httpChannel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %s", err.Error())
	}
	defer icmpChannel.Close()

	icmpQueue, err := icmpChannel.QueueDeclare("icmp", false, false, false, false, nil)

	if err != nil {
		log.Fatalf("declaring queue error: %s", err.Error())
	}

	httpQueue, err := httpChannel.QueueDeclare("http", false, false, false, false, nil)

	if err != nil {
		log.Fatalf("declaring queue error: %s", err.Error())
	}

	icmpMsgs, err := icmpChannel.Consume(icmpQueue.Name, "", true, false, false, false, nil)
	httpMsgs, err := httpChannel.Consume(httpQueue.Name, "", true, false, false, false, nil)

	if err != nil {
		log.Fatalf("consuming queue error: %s", err.Error())
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	forever := make(chan bool)
	go func() {
		for d := range icmpMsgs {
			res, err := worker.TestByICMP(string(d.Body))
			if err != nil {
				log.Fatalf("Error on ICMP test: %s", err.Error())
			}
			// Mandar para o Websocket um signal para se comunicar com o frontend
			fmt.Printf("Received: %v; Loss: %v; Sent: %v \n", res.PacketsRecv, res.PacketLoss, res.PacketsSent)
		}
	}()

	go func() {
		for d := range httpMsgs {
			var content worker.HttpMessageContent
			err := json.Unmarshal(d.Body, &content)
			if err != nil {
				log.Fatalf("Error on parse body from message: %s", err.Error())
				continue
			}
			res, err := worker.TestByHTTP(content.Method, content.URL)
			if err != nil {
				log.Fatalf("Error on HTTP test: %s", err.Error())
				continue
			}
			// Mandar para o Websocket um signal para se comunicar com o frontend
			fmt.Printf("Method: %v; Milliseconds: %v; StatusCode: %v \n", res.Method, res.Milliseconds, res.StatusCode)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-sigchan

	log.Printf("interrupted, shutting down")
	forever <- true
}
