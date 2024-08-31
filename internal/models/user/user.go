package user

import (
	"errors"

	"github.com/google/uuid"
)

type Usuario struct {
	ID        uuid.UUID `json:"id"`
	Nome      string    `json:"nome"`
	Sobrenome string    `json:"sobrenome"`
	Email     Email     `json:"email"`
	Senha     Senha     `json:"senha"`
}

func NewUsuario(nome, sobrenome string, email Email, senha Senha) (*Usuario, error) {
	if nome == "" {
		return nil, errors.New("nome é obrigatório")
	}

	if sobrenome == "" {
		return nil, errors.New("sobrenome é obrigatório")
	}

	if email == "" {
		return nil, errors.New("email é obrigatório")
	}

	return &Usuario{
		ID:        uuid.New(),
		Nome:      nome,
		Sobrenome: sobrenome,
		Email:     email,
		Senha:     senha,
	}, nil
}
