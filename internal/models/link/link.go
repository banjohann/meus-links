package link

import (
	"errors"

	"github.com/google/uuid"
)

type Link struct {
	ID         uuid.UUID `json:"id"`
	Nome       string    `json:"nome"`
	UsuarioID  uuid.UUID `json:"usuarioId"`
	Encurtado  string    `json:"encurtado"`
	URLDestino string    `json:"urlDestino"`
	Contagem   int       `json:"contagem"`
}

func NewLink(nome, encurtado, urlDestino string, usuarioId uuid.UUID) (*Link, error) {
	if nome == "" {
		return nil, errors.New("é obrigatório informar um nome para o link")
	}

	if urlDestino == "" {
		return nil, errors.New("é obrigatório informar uma URL para redirecionamento")
	}

	if usuarioId == uuid.Nil {
		return nil, errors.New("é obrigatório informar o ID do usuário")
	}

	return &Link{
		ID:         uuid.New(),
		Nome:       nome,
		UsuarioID:  usuarioId,
		Encurtado:  encurtado,
		URLDestino: urlDestino,
		Contagem:   0,
	}, nil
}
