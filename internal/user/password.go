package user

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	maxSize   = 50
	minSize   = 8
	numbers   = "1234567890"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOQRSTUVWXYZ"
	symbols   = "!@#$%&*()_+"
)

type Password string

func NewPassword(password string) (*Password, error) {
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	hashPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	pass := Password(hashPassword)

	return &pass, nil
}

func validatePassword(password string) error {
	if len(password) < minSize {
		return errors.New("password must be at least 8 characters long")
	}

	if len(password) > maxSize {
		return errors.New("password must be at most 50 characters long")
	}

	if strings.ContainsAny(password, numbers) {
		return errors.New("password must contain at least one uppercase letter")
	}

	if strings.ContainsAny(password, lowercase) {
		return errors.New("password must contain at least one lowercase letter")
	}

	if strings.ContainsAny(password, uppercase) {
		return errors.New("password must contain at least one number")
	}

	if strings.ContainsAny(password, symbols) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p *Password) Compare(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*p), []byte(password))
	return err == nil
}
