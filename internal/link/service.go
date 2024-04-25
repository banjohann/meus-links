package link

import (
	error "github.com/JohannBandelow/meus-links-go/internal/errors"
	"github.com/JohannBandelow/meus-links-go/internal/utils"
)

type LinkService struct {
	repo *LinkRepo
}

const URLSize = 6

func (s *LinkService) CreateLink(linkAPI Link) *error.HttpError {
	link, err := NewLink(linkAPI.Name, utils.RandomURL(URLSize), linkAPI.RedirectsTo, linkAPI.UserID)
	if err != nil {
		return error.ErrorBadRequest(err.Error())
	}

	s.repo.Create(link)

	return nil
}
