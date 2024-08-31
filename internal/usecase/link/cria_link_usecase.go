package link

import (
	"errors"
	"fmt"

	"github.com/JohannBandelow/meus-links-go/internal/models/link"
	"github.com/JohannBandelow/meus-links-go/internal/repository"
	"github.com/JohannBandelow/meus-links-go/internal/service"
	"github.com/JohannBandelow/meus-links-go/internal/shared"
	"github.com/google/uuid"
)

type CriaLinkCmd struct {
	Nome       string `json:"nome"`
	URLDestino string `json:"urlDestino"`
	UsuarioID  string `json:"usuarioId"`
	URLCustom  string `json:"urlCustom"`
}

type CriaLinkUseCase struct {
	Repo        repository.LinkRepo
	UserService service.UserService
}

func (s *CriaLinkUseCase) Handle(cmd CriaLinkCmd) (*link.Link, error) {
	user, err := s.UserService.GetUsuarioByID(cmd.UsuarioID)
	if err != nil || user == nil {
		return nil, errors.New("usuário não encontrado")
	}

	if cmd.URLCustom != "" {
		linkCustom, _ := s.Repo.FindByEncurtado(cmd.URLCustom)
		if linkCustom != nil {
			return nil, fmt.Errorf("link Curto com a URL %s já existe", cmd.URLCustom)
		}
	}

	link, err := link.NewLink(cmd.Nome, shared.RandomURL(6), cmd.URLDestino, uuid.MustParse(cmd.UsuarioID))
	if err != nil {
		return nil, err
	}

	s.Repo.Save(*link)

	return link, nil
}
