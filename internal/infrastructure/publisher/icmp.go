package publisher

import (
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type CaptureIcmpPublisher struct {
	repo domain.HostRepository
	ch *amqp.Channel
	q amqp.Queue
}

func NewCaptureIcmpPublisher(repo domain.HostRepository, ch *amqp.Channel, q amqp.Queue) *CaptureIcmpPublisher {
	return &CaptureIcmpPublisher{
		repo: repo,
		ch: ch,
		q: q,
	}
}

func (cp *CaptureIcmpPublisher) CaptureIcmpAndPublish() error {
	endpoints, err := cp.repo.GetAll()
	if err != nil {
		return err
	}
	for _, currentEndpoint := range endpoints {
		err = cp.ch.Publish("", cp.q.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(currentEndpoint.IPAddress),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
