package repository

import (
	"github.com/JohannBandelow/meus-links-go/internal/user"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) Create(user user.User) {
}
