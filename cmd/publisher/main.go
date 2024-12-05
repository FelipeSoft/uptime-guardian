package main

import (
	"database/sql"
	"fmt"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/publisher"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/robfig/cron"
	"log"
	"os"
)

func main() {
	godotenv.Load("../../.env")
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("error on scheduler connecting to RabbitMQ")
	}
	defer conn.Close()

	icmpChannel, err := conn.Channel()
	if err != nil {
		log.Fatalf("error on scheduler connecting to channel")
	}

	httpChannel, err := conn.Channel()
	if err != nil {
		log.Fatalf("error on scheduler connecting to channel")
	}

	icmpQueue, err := icmpChannel.QueueDeclare("icmp", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error on scheduler to declaring icmp queue error: %s", err.Error())
	}

	httpQueue, err := httpChannel.QueueDeclare("http", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error on scheduler to declaring http queue error: %s", err.Error())
	}

	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatalf("MySQL Connection Error: %s", err.Error())
	}

	endpointRepository := repository.NewEndpointRepositoryMySQL(db)
	hostRepository := repository.NewHostRepositoryMySQL(db)

	c := cron.New()
	captureIcmpPublisher := publisher.NewCaptureIcmpPublisher(hostRepository, icmpChannel, icmpQueue)
	captureHttpPublisher := publisher.NewCaptureHttpPublisher(endpointRepository, httpChannel, httpQueue)

	err = c.AddFunc("@every 10s", func() {
		err = captureIcmpPublisher.CaptureIcmpAndPublish()
		if err != nil {
			log.Fatalf("error on icmp test: %s", err.Error())
		}
	})

	err = c.AddFunc("@every 10s", func() {
		err = captureHttpPublisher.CaptureHttpAndPublish()
		if err != nil {
			log.Fatalf("error on http test: %s", err.Error())
		}
	})

	if err != nil {
		log.Fatalf("scheduler error: %s", err.Error())
	}

	fmt.Println("Publisher is running...")
	c.Start()

	select {}
}
