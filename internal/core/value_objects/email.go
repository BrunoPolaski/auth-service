package valueobjects

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(e string) (Email, error) {
	if ok := isValidEmail(e); !ok {
		return Email{}, errors.New("Invalid email")
	}

	return Email{e}, nil
}

func isValidEmail(email string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func (p *Email) Value() string {
	return p.value
}
