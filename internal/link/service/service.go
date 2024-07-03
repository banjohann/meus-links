package service

import (
	"net/http"

	error "github.com/JohannBandelow/meus-links-go/internal/api_error"
	"github.com/JohannBandelow/meus-links-go/internal/link"
	"github.com/JohannBandelow/meus-links-go/internal/utils"
)

type LinkService struct {
	repo *link.LinkRepo
}

const URLSize = 6

func (s *LinkService) CreateLink(w http.ResponseWriter, r *http.Request) *error.HttpError {

	link, err := NewLink(linkAPI.Name, utils.RandomURL(URLSize), linkAPI.RedirectsTo, linkAPI.UserID)
	if err != nil {
		return error.ErrorBadRequest(err.Error())
	}

	s.repo.Create(link)

	return nil
}
