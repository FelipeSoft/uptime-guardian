package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/robfig/cron"
)

func main() {
	godotenv.Load("../../.env")
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("error on scheduler connecting to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("error on scheduler connecting to channel")
	}

	q, err := ch.QueueDeclare("ping", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error on scheduler to declaring queue error: %s", err.Error())
	}

	// paralelamente, testar todos endpoint/ip ativos a cada 10 segundos

	c := cron.New()
	err = c.AddFunc("@every 10s", func() {
		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("192.168.200.1"),
		})
	})

	if err != nil {
		log.Fatalf("scheduler error: %s", err.Error())
	}

	fmt.Println("Publisher is running...")
	c.Start()

	select {}
}