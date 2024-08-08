package user

import "github.com/JohannBandelow/meus-links-go/internal/domain/user"

type UserRepo interface {
	FindByID(id string) (*user.Usuario, error)
	FindByEmail(email string) *user.Usuario
	FindByUsername(username string) (*user.Usuario, error)
	Save(user *user.Usuario) error
	Update(user *user.Usuario) error
	Delete(id string) error
}
