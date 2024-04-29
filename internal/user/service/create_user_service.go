package service

import (
	"errors"

	"github.com/JohannBandelow/meus-links-go/internal/user"
)

type CreateUserReq struct {
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

func NewCreateUserResp(user user.User) *CreateUserResp {
	return &CreateUserResp{
		ID:        user.ID.String(),
		FirstName: user.Nome,
		LastName:  user.Sobrenome,
		Email:     user.Email,
	}
}

func (s *UserService) CreateUser(cmd CreateUserReq) (*CreateUserResp, error) {
	var err error

	if cmd.Nome == "" {
		return nil, errors.New("nome é obrigatório")
	}

	if cmd.Sobrenome == "" {
		return nil, errors.New("sobrenome é obrigatório")
	}

	if cmd.Email == "" {
		return nil, errors.New("email é obrigatório")
	}

	existsUser := s.repo.FindByEmail(cmd.Email)
	if existsUser != nil {
		return nil, errors.New("email já cadastrado")
	}

	senha, err := user.NewPassword(cmd.Senha)

	if err != nil {
		return nil, err
	}

	user := user.NewUser(cmd.Nome, cmd.Sobrenome, cmd.Email, senha)

	err = s.repo.Save(*user)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return NewCreateUserResp(*user), nil
}
