package link

import (
	"errors"

	"github.com/JohannBandelow/meus-links-go/internal/repository/link"
)

type RemoverLinkUseCase struct {
	repo link.LinkRepo
}

func (s *RemoverLinkUseCase) Handle(linkID string) error {
	link, err := s.repo.FindByID(linkID)
	if err != nil || link == nil {
		return errors.New("link não encontrado")
	}

	return s.repo.Delete(link.ID.String())
}
