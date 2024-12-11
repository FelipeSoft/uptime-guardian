package usecase

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/rabbitmq"
	"golang.org/x/net/websocket"
)

type ConsumeMetricsUseCase struct {
	rabbitmq *rabbitmq.RabbitMQ
	clients  map[*websocket.Conn]bool
}

func NewConsumeMetricsUseCase(rabbitmq *rabbitmq.RabbitMQ, clients map[*websocket.Conn]bool) *ConsumeMetricsUseCase {
	return &ConsumeMetricsUseCase{
		rabbitmq: rabbitmq,
		clients:  clients,
	}
}

func (c *ConsumeMetricsUseCase) ConsumeAvailableHostsMetrics() {
	fmt.Println("Hello From Consumer")
	msgs, err := c.rabbitmq.Consume("icmp_queue", "icmp_host_websocket_consumer")
	if err != nil {
		log.Printf("log this error: %s", err.Error())
	}
	fmt.Println(msgs)
	for msg := range msgs {
		for client := range c.clients {
			byteMsg, err := json.Marshal(msg)
			if err != nil {
				fmt.Printf("Error on consuming host websocket consumer: %s", err.Error())
				continue
			}
			websocket.Message.Send(client, byteMsg)
		}
		msg.Ack(false)
	}
}
