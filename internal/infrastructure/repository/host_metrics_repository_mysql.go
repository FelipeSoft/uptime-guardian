package repository

import (
	"database/sql"

	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type HostMetricsRepositoryMySQL struct {
	db *sql.DB
}

func NewHostMetricsRepositoryMySQL(db *sql.DB) *HostMetricsRepositoryMySQL {
	return &HostMetricsRepositoryMySQL{
		db: db,
	}
}

func (repo *HostMetricsRepositoryMySQL) Log(metric *domain.HostMetric) error {
	query := "INSERT INTO host_metrics (host_id, packages_received, packages_sent, packages_loss, latency) VALUES (?,?,?,?,?)"
	_, err := repo.db.Query(query, metric.Host_ID, metric.Packages_Received, metric.Packages_Sent, metric.Packages_Loss, metric.Latency)
	if err != nil {
		return err
	}
	return nil
}
