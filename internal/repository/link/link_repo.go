package link

import (
	"github.com/JohannBandelow/meus-links-go/internal/domain/link"
)

type LinkRepo interface {
	FindAll() ([]*link.Link, error)
	FindByID(id string) (*link.Link, error)
	FindByEncurtado(encurtado string) (*link.Link, error)
	FindByUsuarioID(usuarioID string) ([]link.Link, error)
	Save(link *link.Link) error
	Update(link *link.Link) error
	Delete(id string) error
}
