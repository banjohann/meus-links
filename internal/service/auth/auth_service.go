package auth

import "errors"

type AuthService struct {
}

type LoginCmd struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type LoginUserResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func Login(cmd LoginCmd) (*LoginUserResp, error) {
	user := s.repo.FindByEmail(cmd.Email)
	if user == nil {
		return nil, errors.New("credenciais inválidas")
	}

	if !user.Senha.Compare(cmd.Senha) {
		return nil, errors.New("credenciais inválidas")
	}

	return &LoginUserResp{"", ""}, nil
}