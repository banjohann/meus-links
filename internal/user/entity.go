package user

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	Nome      string    `json:"nome"`
	Sobrenome string    `json:"sobrenome"`
	Email     string    `json:"email"`
	Senha     Password  `json:"senha"`
}

func NewUser(nome, sobrenome, email string, senha Password) *User {
	return &User{
		ID:        uuid.New(),
		Nome:      nome,
		Sobrenome: sobrenome,
		Email:     email,
		Senha:     senha,
	}
}
