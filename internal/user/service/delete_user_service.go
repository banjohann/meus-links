package service

import "fmt"

func (s *UserService) DeleteUser(userID string) error {

	_, err := s.repo.Get(userID)
	if err != nil {
		return fmt.Errorf("usuário não encontrado com o id: %s", userID)
	}

	//TODO: chamar método para remover todos os links apos exluir usuário
	err = s.repo.Delete(userID)
	return err
}
