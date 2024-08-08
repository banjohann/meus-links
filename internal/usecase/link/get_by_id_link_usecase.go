package link

import (
	"github.com/JohannBandelow/meus-links-go/internal/domain/link"
	link_repo "github.com/JohannBandelow/meus-links-go/internal/repository/link"
)

type GetByIdLinkUseCase struct {
	repo link_repo.LinkRepo
}

func (s *GetByIdLinkUseCase) Handle(linkID string) (*link.Link, error) {
	link, err := s.repo.FindByID(linkID)
	if err != nil {
		return nil, err
	}

	return link, nil
}
