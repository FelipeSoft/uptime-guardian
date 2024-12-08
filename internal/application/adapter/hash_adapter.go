package adapter

import "golang.org/x/crypto/bcrypt"

type BcryptHashAdapter struct{}

func NewBcryptHashAdapter() *BcryptHashAdapter {
	return &BcryptHashAdapter{}
}

func (b *BcryptHashAdapter) Hash(password string, salt int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hash), err
}

func (b *BcryptHashAdapter) Compare(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
