package domain

type Host struct {
	ID        uint64
	IPAddress string
	Interval  int64
	Timeout   int64
	CreatedAt string
}

type HostRepository interface {
	GetAll() ([]*Host, error)
	GetById(id uint64) (*Host, error)
	Create(host *Host) error
	Update(host *Host) error
	Delete(id uint64) error
}
