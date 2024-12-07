package domain

import "time"

type User struct {
	Id        uint64
	Email     string
	Password  string
	CreatedAt time.Time
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
}
