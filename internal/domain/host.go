package domain

import "errors"

type Host struct {
	ID        uint64
	IPAddress string
	Interval  int64
	Timeout   int64
	CreatedAt string
}

func NewHost(ID uint64, IPAddress string, Interval int64, Timeout int64, CreatedAt string) (*Host, error) {
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
		Timeout:   Timeout,
	}, nil
}

type HostRepository interface {
	GetAll() ([]*Host, error)
	GetById(id uint64) (*Host, error)
	Create(host *Host) error
	Update(host *Host) error
	Delete(id uint64) error
}
