package user

import (
	"errors"
	"fmt"

	"github.com/JohannBandelow/meus-links-go/internal/models/user"
	"github.com/JohannBandelow/meus-links-go/internal/repository"
)

type AtualizaUsuarioCmd struct {
	ID        string `json:"id"`
	NovoEmail string `json:"novo_email"`
	NovaSenha string `json:"nova_senha"`
}

type AtualizaUsuarioResponse struct {
	ID string `json:"id"`
}

type AtualizaUsuarioUseCase struct {
	Repo repository.UserRepo
}

func (s *AtualizaUsuarioUseCase) Handle(cmd AtualizaUsuarioCmd) (*AtualizaUsuarioResponse, error) {

	if cmd.ID == "" {
		return nil, errors.New("é necessário informar o ID")
	}

	usuario, err := s.Repo.FindByID(cmd.ID)
	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado com o ID informado: %s", cmd.ID)
	}

	if cmd.NovoEmail != "" {
		email, err := user.NewEmail(cmd.NovoEmail)
		if err != nil {
			return nil, fmt.Errorf("email informado em formato inválido")
		}

		userWithNewEmail := s.Repo.FindByEmail(cmd.NovoEmail)
		if userWithNewEmail != nil && usuario.ID != userWithNewEmail.ID {
			return nil, fmt.Errorf("email já cadastrado")
		}

		usuario.Email = email
	}

	if cmd.NovaSenha != "" {
		senha, err := user.NewSenha(cmd.NovaSenha)
		if err != nil {
			return nil, fmt.Errorf("a nova senha informada é inválida pois: %s", err.Error())
		}
		usuario.Senha = senha
	}

	err = s.Repo.Update(usuario)
	if err != nil {
		return nil, fmt.Errorf("erro ao salvar usuário no banco")
	}

	return &AtualizaUsuarioResponse{ID: usuario.ID.String()}, nil
}
