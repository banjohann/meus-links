package service

import (
	"errors"

	"github.com/JohannBandelow/meus-links-go/internal/user"
	"github.com/JohannBandelow/meus-links-go/internal/vo"
)

type CriarUsuarioCmd struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
}

type CriarUsuarioResponse struct {
	ID        string `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
}

func (s *UserService) CriaUsuario(cmd CriarUsuarioCmd) (*CriarUsuarioResponse, error) {
	senha, err := vo.NewSenha(cmd.Senha)
	if err != nil {
		return nil, err
	}

	email, err := vo.NewEmail(cmd.Email)
	if err != nil {
		return nil, err
	}

	user, err := user.NewUsuario(cmd.Nome, cmd.Sobrenome, email, senha)
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

	return &CriarUsuarioResponse{
		ID:        user.ID.String(),
		Nome:      user.Nome,
		Sobrenome: user.Sobrenome,
		Email:     user.Email.String(),
	}, nil
}
