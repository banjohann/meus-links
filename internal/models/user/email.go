package user

import (
	"errors"

	"regexp"
)

type Email string

func (e Email) String() string {
	return string(e)
}

func NewEmail(email string) (Email, error) {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return "", err
	}

	if !match {
		return "", errors.New("email informado em formato inválido")
	}

	return Email(email), nil
}
