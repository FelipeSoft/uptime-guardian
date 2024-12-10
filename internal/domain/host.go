package domain

import (
	"errors"
	"time"
)

type Host struct {
	ID        uint64
	IPAddress string
	Interval  int64
	Timeout   int64
	Period    int64
	CreatedAt string
	UpdatedAt string
}

type HostMetric struct {
	ID                uint64
	Host_ID           uint64
	Packages_Received int
	Packages_Sent     int
	Packages_Loss     int
	Latency           time.Duration
}

type WsHostMessage struct {
	Host_ID uint64
	Status  bool
}

func NewHostMetric(ID uint64, Host_ID uint64, Packages_Received int, Packages_Sent int, Packages_Loss int, Latency time.Duration) *HostMetric {
	return &HostMetric{
		ID:                ID,
		Host_ID:           Host_ID,
		Packages_Received: Packages_Received,
		Packages_Sent:     Packages_Sent,
		Packages_Loss:     Packages_Loss,
		Latency:           Latency,
	}
}

func NewHost(ID uint64, IPAddress string, Interval int64, Timeout int64, Period int64, CreatedAt string, UpdatedAt string) (*Host, error) {
	if Timeout <= 0 {
		return nil, errors.New("timeout should be greater than 0 seconds")
	}
	if Interval < 10 {
		return nil, errors.New("interval should be greater than 10 seconds")
	}
	return &Host{
		ID:        ID,
		IPAddress: IPAddress,
		Interval:  Interval,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
		Timeout:   Timeout,
		Period:    Period,
	}, nil
}

type HostRepository interface {
	GetAll() ([]*Host, error)
	GetById(id uint64) (*Host, error)
	Create(host *Host) error
	Update(host *Host) error
	Delete(id uint64) error
}

type HostMetricRepository interface {
	Log(metric *HostMetric) error
}
