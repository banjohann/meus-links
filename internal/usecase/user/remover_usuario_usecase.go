package user

import (
	"fmt"

	"github.com/JohannBandelow/meus-links-go/internal/repository/user"
)

type RemoveUsuarioUseCase struct {
	repo user.UserRepo
}

func (s *RemoveUsuarioUseCase) Handle(usuarioID string) error {

	usuario, err := s.repo.FindByID(usuarioID)
	if err != nil || usuario == nil {
		return fmt.Errorf("usuário não encontrado com o id: %s", usuarioID)
	}

	//TODO: chamar método para remover todos os links apos exluir usuário
	err = s.repo.Delete(usuarioID)
	return err
}
