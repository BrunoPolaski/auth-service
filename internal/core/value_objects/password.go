package valueobjects

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const HASH_COST = 18

type Password struct {
	value string
}

func NewPassword(p string) (Password, error) {
	if ok := isValidPassword(p); !ok {
		return Password{}, errors.New("Invalid password")
	}

	return Password{p}, nil
}

func isValidPassword(p string) bool {
	if len(p) < 8 {
		return false
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, c := range p {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigit = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasDigit && hasSpecial
}

func (p *Password) EncryptPassword() error {
	if hash, err := bcrypt.GenerateFromPassword([]byte(p.value), HASH_COST); err != nil {
		return err
	} else {
		p.value = string(hash)
		return nil
	}
}

func (p *Password) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(p.value), []byte(password)); err != nil {
		return err
	} else {
		return nil
	}
}

func (p *Password) Value() string {
	return p.value
}
