package user_service

import "fmt"

func (s *UserService) RemoveUsuario(usuarioID string) error {

	_, err := s.repo.Get(usuarioID)
	if err != nil {
		return fmt.Errorf("usuário não encontrado com o id: %s", usuarioID)
	}

	//TODO: chamar método para remover todos os links apos exluir usuário
	err = s.repo.Delete(usuarioID)
	return err
}
