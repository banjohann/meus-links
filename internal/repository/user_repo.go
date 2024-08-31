package repository

import "github.com/JohannBandelow/meus-links-go/internal/models/user"

type UserRepo interface {
	FindByID(id string) (*user.Usuario, error)
	FindByEmail(email string) *user.Usuario
	Save(user user.Usuario) error
	Update(user *user.Usuario) error
	Delete(id string) error
}
