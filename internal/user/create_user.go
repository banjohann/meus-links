package user

import (
	"errors"
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

func NewCreateUserResp(user User) CreateUserResp {
	return CreateUserResp{
		ID:        user.ID.String(),
		FirstName: user.Nome,
		LastName:  user.Sobrenome,
		Email:     user.Email,
	}
}

func (s *UserService) CreateUser(cmd CreateUserReq) (*User, error) {
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

	senha, err := NewPassword(cmd.Senha)

	if err != nil {
		return nil, err
	}

	user := NewUser(cmd.Nome, cmd.Sobrenome, cmd.Email, senha)
	s.repo.Save(*user)

	return user, nil
}
