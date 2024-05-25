package service

import (
	"errors"
)

type LoginUserCmd struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type LoginUserResp struct {
	Token string `json:"token"`
}

func (s *UserService) LoginUser(cmd LoginUserCmd) (*LoginUserResp, error) {
	user := s.repo.FindByEmail(cmd.Email)
	if user == nil {
		return nil, errors.New("credenciais inválidas")
	}

	if !user.Senha.Compare(cmd.Senha) {
		return nil, errors.New("credenciais inválidas")
	}

	return &LoginUserResp{user.ID.String()}, nil
}
