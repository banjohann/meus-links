package link

import (
	"errors"

	"github.com/JohannBandelow/meus-links-go/internal/repository"
)

type RemoverLinkUseCase struct {
	Repo repository.LinkRepo
}

func (s *RemoverLinkUseCase) Handle(linkID string) error {
	link, err := s.Repo.FindByID(linkID)
	if err != nil || link == nil {
		return errors.New("link não encontrado")
	}

	return s.Repo.Delete(link.ID.String())
}
