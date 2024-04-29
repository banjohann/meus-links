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

func (s *UserService) DeleteUser(userID string) error {
	// implementation of the DeleteUser method
	s.repo.Delete(userID)
	return nil
}

func (s *UserService) GetUser(userID string) (*user.User, error) {
	// implementation of the GetUser method

	return nil, nil
}

func (s *UserService) UpdateUser(user user.User) error {
	// implementation of the UpdateUser method
	return nil
}
