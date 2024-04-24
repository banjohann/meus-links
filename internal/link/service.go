package link

import (
	"github.com/JohannBandelow/meus-links-go/internal/entities"
	error "github.com/JohannBandelow/meus-links-go/internal/errors"
)

type LinkService struct {
	repo *LinkRepo
}

const URLSize = 6

func (s *LinkService) CreateLink(linkAPI Link) *error.HttpError {
	link, err := NewLink(linkAPI.Name, entities.RandomURL(URLSize), linkAPI.RedirectsTo, linkAPI.UserID)
	if err != nil {
		return error.ErrorBadRequest(err.Error())
	}

	s.repo.Create(link)

	return nil
}
