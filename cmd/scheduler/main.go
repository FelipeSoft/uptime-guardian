package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/rabbitmq"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/repository"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/scheduler/icmp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

// consumers: metrics & websocket

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Environment Variable Error: %s", err.Error())
	}

	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatalf("MySQL Connection Error: %s", err.Error())
	}

	hostRepository := repository.NewHostRepositoryMySQL(db)
	// hostMetricsRepository := repository.NewHostMetricsRepositoryMySQL(db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	refreshInterval := 1 * time.Second
	rabbitmq, err := rabbitmq.NewRabbitMQ(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("RabbitMQ Connection Error: %s", err.Error())
	}
	defer rabbitmq.Close()

	args := amqp.Table{
		"x-dead-letter-exchange":    "icmp_dlx",
		"x-dead-letter-routing-key": "icmp_dead_letter_queue",
	}
	_, err = rabbitmq.DeclareQueue("icmp_queue", args)
	if err != nil {
		log.Fatalf("ICMP Queue Declaring Error: %s", err.Error())
	}

	_, err = rabbitmq.DeclareQueue("icmp_dead_letter_queue", args)
	if err != nil {
		log.Fatalf("ICMP Dead Letter Queue Declaring Error: %s", err.Error())
	}

	go icmp.StartTaskManager(ctx, hostRepository, refreshInterval, rabbitmq)
	fmt.Println("Scheduler Service is running")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	cancel()
	icmp.GracefulShutdown(ctx)

	log.Println("Shutdown complete. Exiting.")
}
