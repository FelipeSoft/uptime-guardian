package icmp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/rabbitmq"
)

var (
	HostTaskRegistry = make(map[uint64]*HostTask)
	TaskMutex        sync.Mutex
	TaskWaitGroup    sync.WaitGroup
)

type HostTask struct {
	Host    *domain.Host
	running bool
	cancel  context.CancelFunc
	ticker  *time.Ticker
	done    chan struct{}
	ctx     context.Context
	queue   *rabbitmq.RabbitMQ
}

func NewHostTask(host *domain.Host, parentCtx context.Context, queue *rabbitmq.RabbitMQ) *HostTask {
	ctx, cancel := context.WithCancel(parentCtx)
	ticker := time.NewTicker(time.Duration(host.Interval) * time.Second)
	return &HostTask{
		Host:    host,
		running: false,
		cancel:  cancel,
		ticker:  ticker,
		done:    make(chan struct{}),
		ctx:     ctx,
		queue:   queue,
	}
}

func (ht *HostTask) Start() {
	TaskWaitGroup.Add(1)
	go func() {
		defer close(ht.done)
		defer TaskWaitGroup.Done()
		for {
			select {
			case <-ht.ticker.C:
				if !ht.running {
					ht.running = true
					go ht.executeTask()
				}
			case <-ht.ctx.Done():
				ht.ticker.Stop()
				return
			}
		}
	}()
}

func (ht *HostTask) Stop() {
	ht.cancel()
	<-ht.done
}

func (ht *HostTask) executeTask() {
	defer func() { ht.running = false }()
	ctx, cancel := context.WithTimeout(ht.ctx, time.Duration(ht.Host.Timeout)*time.Second)
	defer cancel()
	stats := TestByICMP(ctx, ht.Host.IPAddress)
	body, err := json.Marshal(map[string]interface{}{
		"host_id":           ht.Host.ID,
		"packages_received": stats.PacketsRecv,
		"packages_sent":     stats.PacketsSent,
		"packages_loss":     stats.PacketsSent - stats.PacketsRecv,
		"latency":           stats.AvgRtt,
	})
	if err != nil {
		fmt.Printf("error on logging icmp metric: %s", err.Error())
	}
	err = ht.queue.Publish("icmp_queue", body)
	if err != nil {
		fmt.Printf("error on queue message: %s", err.Error())
	}
}

func hostsEqual(a, b *domain.Host) bool {
	return a.ID == b.ID &&
		a.IPAddress == b.IPAddress &&
		a.Interval == b.Interval &&
		a.Timeout == b.Timeout
}

func UpdateHostTask(hosts []*domain.Host, ctx context.Context, queue *rabbitmq.RabbitMQ) {
	TaskMutex.Lock()
	defer TaskMutex.Unlock()

	currentHostIDs := make(map[uint64]bool)
	for id := range HostTaskRegistry {
		currentHostIDs[id] = true
	}

	for _, host := range hosts {
		if _, exists := currentHostIDs[host.ID]; !exists {
			ht := NewHostTask(host, ctx, queue)
			HostTaskRegistry[host.ID] = ht
			ht.Start()
		} else {
			currentHostIDs[host.ID] = false
			existingHT := HostTaskRegistry[host.ID]
			if !hostsEqual(existingHT.Host, host) {
				existingHT.Stop()
				newHT := NewHostTask(host, ctx, queue)
				HostTaskRegistry[host.ID] = newHT
				newHT.Start()
			}
		}
	}

	for id, exists := range currentHostIDs {
		if exists {
			HostTaskRegistry[id].Stop()
			delete(HostTaskRegistry, id)
		}
	}
}

func StartTaskManager(ctx context.Context, hostRepo domain.HostRepository, refreshInterval time.Duration, queue *rabbitmq.RabbitMQ) {
	ticker := time.NewTicker(refreshInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			hosts, err := hostRepo.GetAll()
			if err != nil {
				log.Printf("\n Error fetching hosts: %s", err.Error())
				continue
			}
			UpdateHostTask(hosts, ctx, queue)
		}
	}
}

func GracefulShutdown(ctx context.Context) {
	TaskMutex.Lock()
	defer TaskMutex.Unlock()

	for _, ht := range HostTaskRegistry {
		ht.Stop()
	}

	TaskWaitGroup.Wait()
}
