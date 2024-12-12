package shared

import (
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var icmpQueueMessagesMap = sync.Map{}
var consumersCount = 2

func ProcessICMPMessage(msg amqp.Delivery) {
	key := string(msg.Body)
	currentConsumerCounter := 1
	if val, ok := icmpQueueMessagesMap.Load(key); ok {
		currentConsumerCounter = val.(int) + 1
	}
	icmpQueueMessagesMap.Store(key, currentConsumerCounter)

	if currentConsumerCounter == consumersCount {
		err := msg.Ack(false)
		if err != nil {
			log.Printf("error acknowledging message from icmp_queue: %s", err.Error())
		} else {
			log.Printf("Message: %s acknowledged after processing by all consumers", key)
			icmpQueueMessagesMap.Delete(key)
		}
	}
}
