package link

import (
	"errors"
	"fmt"

	"github.com/JohannBandelow/meus-links-go/internal/models/link"
	"github.com/JohannBandelow/meus-links-go/internal/repository"
)

type AtualizaLinkCmd struct {
	Nome       string `json:"nome"`
	URLDestino string `json:"urlDestino"`
	URLCustom  string `json:"urlCustom"`
}

type AtualizaLinkUseCase struct {
	Repo repository.LinkRepo
}

func (s *AtualizaLinkUseCase) Handle(linkID string, cmd AtualizaLinkCmd) (*link.Link, error) {
	linkFound, err := s.Repo.FindByID(linkID)
	if err != nil || linkFound == nil {
		return nil, errors.New("link não encontrado")
	}

	linkFound.Nome = cmd.Nome
	linkFound.URLDestino = cmd.URLDestino

	if cmd.URLCustom != "" {
		linkCustom, _ := s.Repo.FindByEncurtado(cmd.URLCustom)
		if linkCustom != nil {
			return nil, fmt.Errorf("link Curto com a URL %s já existe", cmd.URLCustom)
		}

		linkFound.Encurtado = cmd.URLCustom
	}

	s.Repo.Save(*linkFound)

	return linkFound, nil
}
