package link_service

import (
	"github.com/JohannBandelow/meus-links-go/internal/link"
	user_service "github.com/JohannBandelow/meus-links-go/internal/user/service"
)

type LinkService struct {
	repo        *link.LinkRepo
	userService *user_service.UserService
}

const URLSize = 6
