package user

import (
	"errors"
)

type UserService struct {
	repo *UserRepo
}

func NewService(repo *UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(cmd CreateUserReq) (*User, error) {
	var err error

	if cmd.Nome == "" {
		errors.Join(err, errors.New("nome é obrigatório"))
	}

	if cmd.Sobrenome == "" {
		errors.Join(err, errors.New("sobrenome é obrigatório"))
	}

	if cmd.Email == "" {
		errors.Join(err, errors.New("email é obrigatório"))
	}

	senha, errSenha := NewPassword(cmd.Senha)
	if errSenha != nil {
		errors.Join(err, errSenha)
	}

	if err != nil {
		return nil, err
	}

	user := NewUser(cmd.Nome, cmd.Sobrenome, cmd.Email, *senha)
	s.repo.Save(*user)

	return user, nil
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
