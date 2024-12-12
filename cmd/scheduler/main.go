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

	_, err = rabbitmq.DeclareQueue("icmp_queue_websocket")
	if err != nil {
		log.Fatalf("ICMP Queue Declaring Error: %s", err.Error())
	}

	_, err = rabbitmq.DeclareQueue("icmp_queue_http")
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
