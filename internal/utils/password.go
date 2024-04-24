package utils

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	numbers   = "1234567890"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOQRSTUVWXYZ"
	symbols   = "!@#$%&*()_+"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if len(password) > 50 {
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
