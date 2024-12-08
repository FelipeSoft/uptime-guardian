package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/repository"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/scheduler/icmp"
	"github.com/joho/godotenv"
)

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	refreshInterval := 10 * time.Second
	go icmp.StartTaskManager(ctx, hostRepository, refreshInterval)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	cancel()
	icmp.GracefulShutdown(ctx)

	log.Println("Shutdown complete. Exiting.")
}