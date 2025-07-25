package valueobjects

import (
	"errors"
	"regexp"

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

func isValidPassword(password string) bool {
	var re = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$`)
	return re.MatchString(password)
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
