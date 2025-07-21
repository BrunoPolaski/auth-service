package crypto

import (
	"github.com/BrunoPolaski/go-rest-err/rest_err"
	"golang.org/x/crypto/bcrypt"
)

type bcryptAdapter struct{}

func NewBcryptAdapter() *bcryptAdapter {
	return &bcryptAdapter{}
}

func (b *bcryptAdapter) EncryptPassword(password string) (string, *rest_err.RestErr) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	return string(hash), nil
}

func (b *bcryptAdapter) ComparePasswords(hashPassword, password string) *rest_err.RestErr {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return rest_err.NewUnauthorizedError("invalid password")
	}

	return nil
}
