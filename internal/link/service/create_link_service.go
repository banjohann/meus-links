package service

import (
	"github.com/JohannBandelow/meus-links-go/internal/link"
	"github.com/JohannBandelow/meus-links-go/internal/utils"
	"github.com/google/uuid"
)

type CreateLinkCmd struct {
	Name        string `json:"name"`
	RedirectsTo string `json:"redirects_to"`
	UserID      string `json:"user_id"`
	CustomURL   string `json:"custom_url"`
}

func (s *LinkService) CreateLink(cmd CreateLinkCmd) (*link.Link, error) {

	link, err := link.NewLink(cmd.Name, utils.RandomURL(URLSize), cmd.RedirectsTo, uuid.MustParse(cmd.UserID))
	if err != nil {
		return nil, err
	}

	s.repo.Create(link)

	return link, nil
}
