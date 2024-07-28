package user_service

import (
	"fmt"

	"github.com/google/uuid"
)

type GetUsuarioResponse struct {
	ID        uuid.UUID `json:"id"`
	Nome      string    `json:"nome"`
	Sobrenome string    `json:"sobrenome"`
	Email     string    `json:"email"`
}

func (s *UserService) GetUsuarioByID(id string) (*GetUsuarioResponse, error) {

	user, err := s.repo.Get(id)
	if err != nil || user == nil {
		return nil, fmt.Errorf("usuário não encontrado com o ID informado: %s", id)
	}

	return &GetUsuarioResponse{
		ID:        user.ID,
		Nome:      user.Nome,
		Sobrenome: user.Sobrenome,
		Email:     user.Email.String(),
	}, nil
}
