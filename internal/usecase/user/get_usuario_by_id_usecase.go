package user

import (
	"fmt"

	user_repo "github.com/JohannBandelow/meus-links-go/internal/repository/user"
	"github.com/google/uuid"
)

type GetUsuarioResponse struct {
	ID        uuid.UUID `json:"id"`
	Nome      string    `json:"nome"`
	Sobrenome string    `json:"sobrenome"`
	Email     string    `json:"email"`
}

type GetUsuarioByIdUseCase struct {
	repo user_repo.UserRepo
}

func (s *GetUsuarioByIdUseCase) Handle(id string) (*GetUsuarioResponse, error) {

	user, err := s.repo.FindByID(id)
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
