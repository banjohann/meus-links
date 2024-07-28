package link_service

import (
	"errors"
	"fmt"

	"github.com/JohannBandelow/meus-links-go/internal/link"
	"github.com/JohannBandelow/meus-links-go/internal/utils"
	"github.com/google/uuid"
)

type CriaLinkCmd struct {
	Nome       string `json:"nome"`
	URLDestino string `json:"urlDestino"`
	UsuarioID  string `json:"usuarioId"`
	URLCustom  string `json:"urlCustom"`
}

func (s *LinkService) CriaLink(cmd CriaLinkCmd) (*link.Link, error) {
	user, err := s.userService.GetUsuarioByID(cmd.UsuarioID)
	if err != nil || user == nil {
		return nil, errors.New("usuário não encontrado")
	}

	if cmd.URLCustom != "" {
		linkCustom, _ := s.repo.FindByEncurtado(cmd.URLCustom)
		if linkCustom != nil {
			return nil, fmt.Errorf("link Curto com a URL %s já existe", cmd.URLCustom)
		}
	}

	link, err := link.NewLink(cmd.Nome, utils.RandomURL(URLSize), cmd.URLDestino, uuid.MustParse(cmd.UsuarioID))
	if err != nil {
		return nil, err
	}

	s.repo.Save(*link)

	return link, nil
}
