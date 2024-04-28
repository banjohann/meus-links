package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	maxSize   = 50
	minSize   = 8
	numbers   = []rune("1234567890")
	lowercase = []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase = []rune("ABCDEFGHIJKLMNOQRSTUVWXYZ")
	symbols   = []rune("!@#$%&*()_+")
)

type Password string

func (p Password) String() string {
	return string(p)
}

func NewPassword(password string) (Password, error) {
	if err := validatePassword(password); err != nil {
		return "", err
	}

	hashPassword, err := hashPassword(password)
	if err != nil {
		return "", err
	}

	return Password(hashPassword), nil
}

func (p *Password) Compare(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*p), []byte(password))
	return err == nil
}

func validatePassword(password string) error {
	if len(password) < minSize {
		return errors.New("password must be at least 8 characters long")
	}

	if len(password) > maxSize {
		return errors.New("password must be at most 50 characters long")
	}

	if !containsRune(password, uppercase) {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !containsRune(password, lowercase) {
		return errors.New("password must contain at least one lowercase letter")
	}

	if !containsRune(password, numbers) {
		return errors.New("password must contain at least one number")
	}

	if !containsRune(password, symbols) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func containsRune(password string, runes []rune) bool {
	for _, c := range password {
		for _, r := range runes {
			if c == r {
				return true
			}
		}
	}
	return false
}
