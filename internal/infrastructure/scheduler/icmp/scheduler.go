package icmp

import (
	"context"
	"log"
	"sync"
	"time"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
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
}

func NewHostTask(host *domain.Host, parentCtx context.Context) *HostTask {
	ctx, cancel := context.WithCancel(parentCtx)
	ticker := time.NewTicker(time.Duration(host.Interval) * time.Second)
	return &HostTask{
		Host:    host,
		running: false,
		cancel:  cancel,
		ticker:  ticker,
		done:    make(chan struct{}),
		ctx:     ctx,
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
	TestByICMP(ctx, ht.Host.IPAddress)
}

func UpdateHostTask(hosts []*domain.Host, ctx context.Context) {
	TaskMutex.Lock()
	defer TaskMutex.Unlock()

	currentHostIDs := make(map[uint64]bool)
	for id := range HostTaskRegistry {
		currentHostIDs[id] = true
	}

	for _, host := range hosts {
		if _, exists := currentHostIDs[host.ID]; !exists {
			ht := NewHostTask(host, ctx)
			HostTaskRegistry[host.ID] = ht
			ht.Start()
		} else {
			currentHostIDs[host.ID] = false
			existingHT := HostTaskRegistry[host.ID]
			if existingHT.Host.Interval != host.Interval || existingHT.Host.Timeout != host.Timeout {
				existingHT.Stop()
				newHT := NewHostTask(host, ctx)
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

func StartTaskManager(ctx context.Context, hostRepo domain.HostRepository, refreshInterval time.Duration) {
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
			UpdateHostTask(hosts, ctx)
		}
	}
}

func GracefulShutdown(ctx context.Context) {
	TaskMutex.Lock()
	defer TaskMutex.Unlock()

	for _, ht := range HostTaskRegistry {
		ht.Stop()
	}
}
