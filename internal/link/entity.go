package link

import (
	"errors"

	"github.com/google/uuid"
)

type Link struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	UserID      uuid.UUID `json:"userID" gorm:"not null"`
	Short       string    `json:"short" gorm:"not null"`
	RedirectsTo string    `json:"redirects_to" gorm:"not null"`
	Clicks      int       `json:"clicks"`
}

func NewLink(name, short, redirectsTo string, userID uuid.UUID) (*Link, error) {
	if name == "" {
		return nil, errors.New("Missing property 'name'")
	}

	if short == "" {
		return nil, errors.New("Missing property 'short'")
	}

	if redirectsTo == "" {
		return nil, errors.New("Missing property 'redirects_to'")
	}

	if userID == uuid.Nil {
		return nil, errors.New("Missing property 'userID'")
	}

	return &Link{
		ID:          uuid.New(),
		Name:        name,
		UserID:      userID,
		Short:       short,
		RedirectsTo: redirectsTo,
		Clicks:      0,
	}, nil
}
