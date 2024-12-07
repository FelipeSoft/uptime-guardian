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
	row := repo.db.QueryRow("SELECT `id`, `email`, `password`, `created_at` FROM user WHERE email = ?", email)

	var output domain.User
	err := row.Scan(&output.Id, &output.Email, &output.Password, &output.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &output, nil
}
