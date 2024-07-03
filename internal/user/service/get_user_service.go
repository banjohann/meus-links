package service

import (
	"fmt"

	"github.com/google/uuid"
)

type GetUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Nome      string    `json:"nome"`
	Sobrenome string    `json:"sobrenome"`
	Email     string    `json:"email"`
}

func (s *UserService) GetUserByID(id string) (*GetUserResponse, error) {

	user, err := s.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado com o ID informado: %s", id)
	}

	return &GetUserResponse{
		ID:        user.ID,
		Nome:      user.Nome,
		Sobrenome: user.Sobrenome,
		Email:     user.Email.String(),
	}, nil
}
