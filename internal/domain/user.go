package domain

type User struct {
	Id        uint64
	Email     string
	Password  string
	CreatedAt string
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
}
