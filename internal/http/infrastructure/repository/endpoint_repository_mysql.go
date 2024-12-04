package repository

import (
	"database/sql"
	"time"
	"github.com/FelipeSoft/uptime-guardian/internal/http/domain"
)

type EndpointRepositoryMySQL struct {
	db *sql.DB
}

func NewEndpointRepositoryMySQL(db *sql.DB) *EndpointRepositoryMySQL {
	return &EndpointRepositoryMySQL{
		db: db,
	}
}

func (r *EndpointRepositoryMySQL) GetAll() ([]*domain.Endpoint, error) {
	rows, err := r.db.Query("SELECT `id`, `address`, `interval`, `timeout`, `created_at` FROM endpoint")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var output []*domain.Endpoint

	for rows.Next() {
		var e domain.Endpoint
		err = rows.Scan(&e.ID, &e.Address, &e.Interval, &e.Timeout, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		output = append(output, &e)
	}

 	return output, nil
}

func (r *EndpointRepositoryMySQL) GetById(id uint64) (*domain.Endpoint, error) {
	row := r.db.QueryRow("SELECT `id`, `address`, `interval`, `timeout`, `created_at` FROM endpoint WHERE id = ?", id)
	
	var e domain.Endpoint
	err := row.Scan(&e.ID, &e.Address, &e.Interval, &e.Timeout, &e.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *EndpointRepositoryMySQL) Create(endpoint *domain.Endpoint) error {
	rows, err := r.db.Query("INSERT INTO endpoint (`address`, `interval`, `timeout`, `created_at`) VALUES (?,?,?,?)",
	endpoint.Address, endpoint.Interval, endpoint.Timeout, time.Now())

	if err != nil {
		return err
	}

	defer rows.Close()
	return nil
}

func (r *EndpointRepositoryMySQL) Update(endpoint *domain.Endpoint) error {
	rows, err := r.db.Query("UPDATE endpoint SET `address` = ?, `interval` = ?, `timeout` = ? WHERE id = ?",
	endpoint.Address, endpoint.Interval, endpoint.Timeout, endpoint.ID)

	if err != nil {
		return err
	}

	defer rows.Close()
	return nil
}

func (r *EndpointRepositoryMySQL) Delete(id uint64) error {
	rows, err := r.db.Query("DELETE FROM endpoint WHERE id = ?", id)

	if err != nil {
		return err
	}

	defer rows.Close()
	return nil
}

