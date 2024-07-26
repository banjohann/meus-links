package service

import (
	"github.com/JohannBandelow/meus-links-go/internal/link"
)

type LinkService struct {
	repo *link.LinkRepo
}

const URLSize = 6
