package link_service

import (
	"errors"
	"fmt"

	"github.com/JohannBandelow/meus-links-go/internal/link"
)

type AtualizaLinkCmd struct {
	Nome       string `json:"nome"`
	URLDestino string `json:"urlDestino"`
	URLCustom  string `json:"urlCustom"`
}

func (s *LinkService) AtualizaLink(linkID string, cmd AtualizaLinkCmd) (*link.Link, error) {
	link, err := s.repo.FindByID(linkID)
	if err != nil || link == nil {
		return nil, errors.New("link não encontrado")
	}

	link.Nome = cmd.Nome
	link.URLDestino = cmd.URLDestino

	if cmd.URLCustom != "" {
		linkCustom, _ := s.repo.FindByEncurtado(cmd.URLCustom)
		if linkCustom != nil {
			return nil, fmt.Errorf("link Curto com a URL %s já existe", cmd.URLCustom)
		}

		link.Encurtado = cmd.URLCustom
	}

	s.repo.Save(*link)

	return link, nil
}
