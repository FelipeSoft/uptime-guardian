package domain

import "errors"

type Endpoint struct {
	ID        uint64
	URL       string
	Method    string
	Interval  int64
	Timeout   int64
	CreatedAt string
	UpdatedAt string
}

func NewEndpoint(ID uint64, URL string, Method string, Interval int64, Timeout int64, CreatedAt string, UpdatedAt string) (*Endpoint, error) {
	if Timeout <= 0 {
		return nil, errors.New("timeout should be greater than 0 seconds")
	}
	if Interval < 10 {
		return nil, errors.New("interval should be greater than 10 seconds")
	}
	return &Endpoint{
		ID:        ID,
		URL:       URL,
		Method:    Method,
		Interval:  Interval,
		Timeout:   Timeout,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}, nil
}

type EndpointRepository interface {
	GetAll() ([]*Endpoint, error)
	GetById(id uint64) (*Endpoint, error)
	Create(endpoint *Endpoint) error
	Update(endpoint *Endpoint) error
	Delete(id uint64) error
}
