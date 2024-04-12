package user

import (
	"os/user"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) Create(user user.User) {
}
