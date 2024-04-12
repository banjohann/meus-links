package user

import "net/http"

type UserService struct {
	repo *UserRepo
}

func NewService(repo *UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
}
