package user

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"    gorm:"primaryKey"`
	Nome      string    `json:"nome"  gorm:"not null"`
	Sobrenome string    `json:"sobrenome" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Senha     string    `json:"senha"`
}

func NewUser(nome, sobrenome, email, senha string) *User {
	return &User{
		ID:        uuid.New(),
		Nome:      nome,
		Sobrenome: sobrenome,
		Email:     email,
		Senha:     senha,
	}
}
