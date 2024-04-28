package user

type UserService struct {
	repo *UserRepo
}

func NewService(repo *UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) DeleteUser(userID string) error {
	// implementation of the DeleteUser method
	s.repo.Delete(userID)
	return nil
}

func (s *UserService) GetUser(userID string) (*User, error) {
	// implementation of the GetUser method

	return nil, nil
}

func (s *UserService) UpdateUser(user User) error {
	// implementation of the UpdateUser method
	return nil
}
