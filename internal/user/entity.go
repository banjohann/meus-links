package user

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Nome      string    `json:"nome"`
	Sobrenome string    `json:"sobrenome"`
	Email     Email     `json:"email"`
	Senha     Password  `json:"senha"`
}

func NewUser(nome, sobrenome string, email Email, senha Password) (*User, error) {
	if nome == "" {
		return nil, errors.New("nome é obrigatório")
	}

	if sobrenome == "" {
		return nil, errors.New("sobrenome é obrigatório")
	}

	if email == "" {
		return nil, errors.New("email é obrigatório")
	}

	return &User{
		ID:        uuid.New(),
		Nome:      nome,
		Sobrenome: sobrenome,
		Email:     email,
		Senha:     senha,
	}, nil
}
