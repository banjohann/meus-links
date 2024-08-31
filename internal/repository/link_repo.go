package repository

import (
	"github.com/JohannBandelow/meus-links-go/internal/models/link"
)

type LinkRepo interface {
	FindByID(id string) (*link.Link, error)
	FindByEncurtado(encurtado string) (*link.Link, error)
	FindByUsuarioID(usuarioID string) ([]link.Link, error)
	Save(link link.Link) error
	Update(link *link.Link) error
	Delete(id string) error
}
