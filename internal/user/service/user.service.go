package service

import "github.com/JohannBandelow/meus-links-go/internal/user"

type UserService struct {
	repo *user.UserRepo
}

func New(repo *user.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}
