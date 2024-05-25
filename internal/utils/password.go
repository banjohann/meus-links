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
	symbols   = "!@#$%&*_"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("a senha tem que ter ao menos 8 carácteres")
	}

	if len(password) > 50 {
		return errors.New("a senha tem ter no máximo 50 carácteres")
	}

	if strings.ContainsAny(password, numbers) {
		return errors.New("a senha tem que ter ao menos 1 número")
	}

	if strings.ContainsAny(password, lowercase) {
		return errors.New("a senha tem que ter ao menos uma letra minúscula")
	}

	if strings.ContainsAny(password, uppercase) {
		return errors.New("a senha tem que ter ao menos uma letra maiúscula")
	}

	if strings.ContainsAny(password, symbols) {
		return errors.New("a senha tem que conter ao menos un caractere especial (!@#$%&*_)")
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
