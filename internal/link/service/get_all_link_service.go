package link_service

import (
	"github.com/JohannBandelow/meus-links-go/internal/link"
	"github.com/google/uuid"
)

type GetAllRequestFilter struct {
	UsuarioId uuid.UUID `json:"usuarioId"`
}

func (s *LinkService) GetAll(requestFilter GetAllRequestFilter) ([]link.Link, error) {

	return s.repo.GetAllLinksDoUsuario(requestFilter.UsuarioId.String())
}
