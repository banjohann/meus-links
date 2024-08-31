package link

import (
	"github.com/JohannBandelow/meus-links-go/internal/models/link"
	"github.com/JohannBandelow/meus-links-go/internal/repository"
)

type GetByIdLinkUseCase struct {
	Repo repository.LinkRepo
}

func (s *GetByIdLinkUseCase) Handle(linkID string) (*link.Link, error) {
	link, err := s.Repo.FindByID(linkID)
	if err != nil {
		return nil, err
	}

	return link, nil
}
