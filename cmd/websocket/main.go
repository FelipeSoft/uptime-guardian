package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/rabbitmq"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/shared"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/websocket/handler"
	"github.com/joho/godotenv"
	"golang.org/x/net/websocket"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Environment Variable Error: %s", err.Error())
	}

	queue, err := rabbitmq.NewRabbitMQ(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("RabbitMQ Connection Error: %s", err.Error())
	}
	defer queue.Close()

	msgs, err := queue.Consume("icmp_queue_websocket")
	if err != nil {
		log.Fatalf("RabbitMQ Consumer Error: %s", err.Error())
	}

	http.Handle("/host/ws", enableCORS(websocket.Handler(handler.HostMetricsWebsocketHandler)))

	wsClients := shared.GetWebsocketClients()

	for i := 0; i < 5; i++ {
		go func() {
			for msg := range msgs {
				for client := range wsClients {
					client.Write(msg.Body)
				}
			}
		}()
	}

	fmt.Printf("Websocket Server started on %s \n", os.Getenv("WEBSOCKET_URL"))
	err = http.ListenAndServe(os.Getenv("WEBSOCKET_URL"), nil)
	if err != nil {
		log.Fatalf("Websocket Server initializing start error: %s", err.Error())
	}
}
