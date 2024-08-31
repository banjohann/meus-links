package service

import (
	"errors"

	"github.com/JohannBandelow/meus-links-go/internal/repository"
)

type AuthService struct {
	userRepo repository.UserRepo
}

type LoginCmd struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type LoginUserResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *AuthService) Login(cmd LoginCmd) (*LoginUserResp, error) {
	user := s.userRepo.FindByEmail(cmd.Email)
	if user == nil {
		return nil, errors.New("credenciais inválidas")
	}

	if !user.Senha.Compare(cmd.Senha) {
		return nil, errors.New("credenciais inválidas")
	}

	return &LoginUserResp{"", ""}, nil
}
