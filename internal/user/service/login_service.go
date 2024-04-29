package service

import (
	"errors"
)

type LoginUserReq struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type LoginUserResp struct {
	ID string `json:"id"`
}

func (s *UserService) LoginUser(cmd LoginUserReq) (*LoginUserResp, error) {
	user := s.repo.FindByEmail(cmd.Email)
	if user == nil {
		return nil, errors.New("usuário não encontrado")
	}

	if !user.Senha.Compare(cmd.Senha) {
		return nil, errors.New("senha inválida")
	}

	return &LoginUserResp{user.ID.String()}, nil
}
