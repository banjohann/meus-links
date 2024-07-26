package service

import (
	"errors"
	"fmt"

	"github.com/JohannBandelow/meus-links-go/internal/vo"
)

type AtualizaUsuarioCmd struct {
	ID        string `json:"id"`
	NovoEmail string `json:"novo_email"`
	NovaSenha string `json:"nova_senha"`
}

type UpdateUserResponse struct {
	ID string `json:"id"`
}

func (s *UserService) AtualizaUsuario(cmd AtualizaUsuarioCmd) error {

	if cmd.ID == "" {
		return errors.New("é necessário informar o ID")
	}

	user, err := s.repo.Get(cmd.ID)
	if err != nil {
		return fmt.Errorf("usuário não encontrado com o ID informado: %s", cmd.ID)
	}

	if cmd.NovoEmail != "" {
		email, err := vo.NewEmail(cmd.NovoEmail)
		if err != nil {
			return fmt.Errorf("email informado em formato inválido")
		}

		userWithNewEmail := s.repo.FindByEmail(cmd.NovoEmail)
		if userWithNewEmail != nil && user.ID != userWithNewEmail.ID {
			return fmt.Errorf("email já cadastrado")
		}

		user.Email = email
	}

	if cmd.NovaSenha != "" {
		senha, err := vo.NewSenha(cmd.NovaSenha)
		if err != nil {
			return fmt.Errorf("a nova senha informada é inválida pois: %s", err.Error())
		}
		user.Senha = senha
	}

	err = s.repo.Update(user)
	if err != nil {
		return fmt.Errorf("erro ao salvar usuário no banco")
	}

	return nil
}
