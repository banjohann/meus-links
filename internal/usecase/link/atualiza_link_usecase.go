package link

import (
	"errors"
	"fmt"

	"github.com/JohannBandelow/meus-links-go/internal/domain/link"
	link_repo "github.com/JohannBandelow/meus-links-go/internal/repository/link"
)

type AtualizaLinkCmd struct {
	Nome       string `json:"nome"`
	URLDestino string `json:"urlDestino"`
	URLCustom  string `json:"urlCustom"`
}

type AtualizaLinkUseCase struct {
	repo link_repo.LinkRepo
}

func (s *AtualizaLinkUseCase) Handle(linkID string, cmd AtualizaLinkCmd) (*link.Link, error) {
	linkFound, err := s.repo.FindByID(linkID)
	if err != nil || linkFound == nil {
		return nil, errors.New("link não encontrado")
	}

	linkFound.Nome = cmd.Nome
	linkFound.URLDestino = cmd.URLDestino

	if cmd.URLCustom != "" {
		linkCustom, _ := s.repo.FindByEncurtado(cmd.URLCustom)
		if linkCustom != nil {
			return nil, fmt.Errorf("link Curto com a URL %s já existe", cmd.URLCustom)
		}

		linkFound.Encurtado = cmd.URLCustom
	}

	s.repo.Save(linkFound)

	return linkFound, nil
}
