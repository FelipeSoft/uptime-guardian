package publisher

import (
	"encoding/json"

	"github.com/FelipeSoft/uptime-guardian/internal/domain"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/worker"
	amqp "github.com/rabbitmq/amqp091-go"
)

type CaptureHttpPublisher struct {
	repo domain.EndpointRepository
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewCaptureHttpPublisher(repo domain.EndpointRepository, ch *amqp.Channel, q amqp.Queue) *CaptureHttpPublisher {
	return &CaptureHttpPublisher{
		repo: repo,
		ch:   ch,
		q:    q,
	}
}

func (cp *CaptureHttpPublisher) CaptureHttpAndPublish() error {
	endpoints, err := cp.repo.GetAll()
	if err != nil {
		return err
	}
	for _, currentEndpoint := range endpoints {
		body, err := json.Marshal(worker.HttpMessageContent{
			Method: currentEndpoint.Method,
			URL:    currentEndpoint.URL,
		})
		if err != nil {
			return err
		}
		err = cp.ch.Publish("", cp.q.Name, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
