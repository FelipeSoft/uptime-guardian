package repository

import (
	"database/sql"

	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type UserRepositoryMySQL struct {
	db *sql.DB
}

func NewUserRepositoryMySQL(db *sql.DB) *UserRepositoryMySQL {
	return &UserRepositoryMySQL{
		db: db,
	}
}

func (repo *UserRepositoryMySQL) GetByEmail(email string) (*domain.User, error) {
	rows, err := repo.db.Query("SELECT email, password, created_at FROM user WHERE email = ?")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var output *domain.User
	err = rows.Scan(&output.Email, &output.Password, &output.CreatedAt)
	if err != nil {
		return nil, err
	}
	return output, nil
}
