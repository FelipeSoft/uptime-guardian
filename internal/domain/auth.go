package domain

type Hashable interface {
	Hash(password string, salt int) (string, error)
	Compare(password string, hash string) bool
}

type Jwt interface {
	Expired(tokenString string) bool
	Read(tokenString string) (*string, error)
	Generate(metadata string) (string, error)
}
