package link

import (
	"github.com/JohannBandelow/meus-links-go/internal/domain/link"
	link_repo "github.com/JohannBandelow/meus-links-go/internal/repository/link"
	"github.com/google/uuid"
)

type GetAllRequestFilter struct {
	UsuarioId uuid.UUID `json:"usuarioId"`
}

type GetAllLinkUseCase struct {
	repo link_repo.LinkRepo
}

func (s *GetAllLinkUseCase) Handle(requestFilter GetAllRequestFilter) ([]link.Link, error) {
	return s.repo.FindByUsuarioID(requestFilter.UsuarioId.String())
}
