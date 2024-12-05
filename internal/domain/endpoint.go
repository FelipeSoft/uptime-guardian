package domain

type Endpoint struct {
	ID        uint64
	URL       string
	Method    string
	Interval  int64
	Timeout   int64
	CreatedAt string
}

type EndpointRepository interface {
	GetAll() ([]*Endpoint, error)
	GetById(id uint64) (*Endpoint, error)
	Create(endpoint *Endpoint) error
	Update(endpoint *Endpoint) error
	Delete(id uint64) error
}
