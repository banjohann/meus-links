package service

import (
	"errors"

	"github.com/JohannBandelow/meus-links-go/internal/user"
)

type CreateUserCmd struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
}

type CreateUserResp struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (s *UserService) CreateUser(cmd CreateUserCmd) (*CreateUserResp, error) {
	senha, err := user.NewPassword(cmd.Senha)
	if err != nil {
		return nil, err
	}

	email, err := user.NewEmail(cmd.Email)
	if err != nil {
		return nil, err
	}

	user, err := user.NewUser(cmd.Nome, cmd.Sobrenome, email, senha)
	if err != nil {
		return nil, err
	}

	existsUser := s.repo.FindByEmail(email.String())
	if existsUser != nil {
		return nil, errors.New("email já cadastrado")
	}

	err = s.repo.Save(*user)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return &CreateUserResp{
		ID:        user.ID.String(),
		FirstName: user.Nome,
		LastName:  user.Sobrenome,
		Email:     user.Email.String(),
	}, nil
}
