package user

import (
	"errors"

	"github.com/JohannBandelow/meus-links-go/internal/models/user"
	"github.com/JohannBandelow/meus-links-go/internal/repository"
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

type CriarUsuarioUseCase struct {
	Repo repository.UserRepo
}

func (s *CriarUsuarioUseCase) Handle(cmd CriarUsuarioCmd) (*CriarUsuarioResponse, error) {
	senha, err := user.NewSenha(cmd.Senha)
	if err != nil {
		return nil, err
	}

	email, err := user.NewEmail(cmd.Email)
	if err != nil {
		return nil, err
	}

	user, err := user.NewUsuario(cmd.Nome, cmd.Sobrenome, email, senha)
	if err != nil {
		return nil, err
	}

	existsUser := s.Repo.FindByEmail(email.String())
	if existsUser != nil {
		return nil, errors.New("email já cadastrado")
	}

	err = s.Repo.Save(*user)
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
