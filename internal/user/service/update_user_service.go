package service

import (
	"errors"
	"fmt"

	user_domain "github.com/JohannBandelow/meus-links-go/internal/user"
	"github.com/JohannBandelow/meus-links-go/internal/utils"
)

type UpdateUserCmd struct {
	ID        string `json:"id"`
	NovoEmail string `json:"novo_email"`
	NovaSenha string `json:"nova_senha"`
}

type UpdateUserResponse struct {
	ID string `json:"id"`
}

func (s *UserService) UpdateUser(cmd UpdateUserCmd) error {

	if cmd.ID == "" {
		return errors.New("é necessário informar o ID")
	}

	user, err := s.repo.Get(cmd.ID)
	if err != nil {
		return fmt.Errorf("usuário não encontrado com o ID informado: %s", cmd.ID)
	}

	if cmd.NovoEmail != "" {
		email, err := user_domain.NewEmail(cmd.NovoEmail)
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
		err := utils.ValidatePassword(cmd.NovaSenha)
		if err != nil {
			return fmt.Errorf("a nova senha informada é inválida pois: %s", err.Error())
		}
	}

	err = s.repo.Update(user)
	if err != nil {
		return fmt.Errorf("erro ao salvar usuário no banco")
	}

	return nil
}
