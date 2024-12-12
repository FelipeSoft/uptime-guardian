package usecase

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/rabbitmq"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/shared"
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
	msgs, err := c.rabbitmq.Consume("icmp_queue")
	if err != nil {
		log.Printf("log this error: %s", err.Error())
	}
	for msg := range msgs {
		shared.ProcessICMPMessage(msg)
		for client := range c.clients {
			byteMsg, err := json.Marshal(msg)
			if err != nil {
				fmt.Printf("Error on consuming host websocket consumer: %s", err.Error())
				continue
			}
			websocket.Message.Send(client, byteMsg)
		}
	}
}
