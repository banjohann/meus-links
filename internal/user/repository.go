package user

import (
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

func (repo *UserRepo) Save(user User) {
	repo.db.Create(&user)
}

func (repo *UserRepo) Delete(userID string) {
}

func (repo *UserRepo) Get(userID string) User {
	return User{}
}

func (repo *UserRepo) Update(user User) {
}
