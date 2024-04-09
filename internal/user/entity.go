package user

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `json:"id"    gorm:"primaryKey"`
	Nome  string    `json:"nome"  gorm:"not null"`
	Email string    `json:"email" gorm:"unique;not null"`
	Senha string    `json:"senha"`
}

func NewUser(nome, email, senha string) *User {
	return &User{
		ID:    uuid.New(),
		Nome:  nome,
		Email: email,
		Senha: senha,
	}
}
