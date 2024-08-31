package link

import (
	"github.com/JohannBandelow/meus-links-go/internal/models/link"
	"github.com/JohannBandelow/meus-links-go/internal/repository"
	"github.com/google/uuid"
)

type GetAllRequestFilter struct {
	UsuarioId uuid.UUID `json:"usuarioId"`
}

type GetAllLinkUseCase struct {
	Repo repository.LinkRepo
}

func (s *GetAllLinkUseCase) Handle(requestFilter GetAllRequestFilter) ([]link.Link, error) {
	return s.Repo.FindByUsuarioID(requestFilter.UsuarioId.String())
}
