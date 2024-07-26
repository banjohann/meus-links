package link

import (
	"errors"

	"github.com/google/uuid"
)

type Link struct {
	ID          uuid.UUID `json:"id"`
	Nome        string    `json:"nome"`
	UserID      uuid.UUID `json:"userId"`
	Short       string    `json:"short"`
	RedirectsTo string    `json:"redirects_to"`
	Clicks      int       `json:"clicks"`
}

func NewLink(nome, short, redirectsTo string, userID uuid.UUID) (*Link, error) {
	if nome == "" {
		return nil, errors.New("É obrigatório informar um nome para o link")
	}

	if redirectsTo == "" {
		return nil, errors.New("É obrigatório informar uma URL para redirecionamento")
	}

	if userID == uuid.Nil {
		return nil, errors.New("É obrigatório informar o ID do usuário")
	}

	return &Link{
		ID:          uuid.New(),
		Nome:        nome,
		UserID:      userID,
		Short:       short,
		RedirectsTo: redirectsTo,
		Clicks:      0,
	}, nil
}
