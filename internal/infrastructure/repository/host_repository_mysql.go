package repository

import (
	"database/sql"
	"time"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type HostRepositoryMySQL struct {
	db *sql.DB
}

func NewHostRepositoryMySQL(db *sql.DB) *HostRepositoryMySQL {
	return &HostRepositoryMySQL{
		db: db,
	}
}

func (r *HostRepositoryMySQL) GetAll() ([]*domain.Host, error) {
	rows, err := r.db.Query("SELECT `id`, `ip_address`, `interval`, `timeout`, `created_at` FROM host")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var output []*domain.Host

	for rows.Next() {
		var e domain.Host
		err = rows.Scan(&e.ID, &e.IPAddress, &e.Interval, &e.Timeout, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		output = append(output, &e)
	}

 	return output, nil
}

func (r *HostRepositoryMySQL) GetById(id uint64) (*domain.Host, error) {
	row := r.db.QueryRow("SELECT `id`, `ip_address`, `interval`, `timeout`, `created_at` FROM host WHERE id = ?", id)
	
	var e domain.Host
	err := row.Scan(&e.ID, &e.IPAddress, &e.Interval, &e.Timeout, &e.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *HostRepositoryMySQL) Create(host *domain.Host) error {
	rows, err := r.db.Query("INSERT INTO host (`ip_address`, `interval`, `timeout`, `created_at`, `updated_at`) VALUES (?,?,?,?,?)",
	host.IPAddress, host.Interval, host.Timeout, time.Now(), time.Now())

	if err != nil {
		return err
	}

	defer rows.Close()
	return nil
}

func (r *HostRepositoryMySQL) Update(host *domain.Host) error {
	rows, err := r.db.Query("UPDATE host SET `ip_address`= ?, `interval`= ?, `timeout` = ?, `updated_at` = ? WHERE id = ?",
	host.IPAddress, host.Interval, host.Timeout, time.Now(), host.ID)

	if err != nil {
		return err
	}

	defer rows.Close()
	return nil
}

func (r *HostRepositoryMySQL) Delete(id uint64) error {
	rows, err := r.db.Query("DELETE FROM host WHERE id = ?", id)

	if err != nil {
		return err
	}

	defer rows.Close()
	return nil
}

