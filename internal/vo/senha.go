package vo

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

type Senha string

func (p Senha) String() string {
	return string(p)
}

func NewSenha(password string) (Senha, error) {
	if err := validaSenha(password); err != nil {
		return "", err
	}

	hashPassword, err := hashPassword(password)
	if err != nil {
		return "", err
	}

	return Senha(hashPassword), nil
}

func (p *Senha) Compare(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*p), []byte(password))
	return err == nil
}

func validaSenha(senha string) error {
	if len(senha) < minSize {
		return errors.New("password must be at least 8 characters long")
	}

	if len(senha) > maxSize {
		return errors.New("password must be at most 50 characters long")
	}

	if !containsRune(senha, uppercase) {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !containsRune(senha, lowercase) {
		return errors.New("password must contain at least one lowercase letter")
	}

	if !containsRune(senha, numbers) {
		return errors.New("password must contain at least one number")
	}

	if !containsRune(senha, symbols) {
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
