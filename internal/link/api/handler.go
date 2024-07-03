package api

import "github.com/JohannBandelow/meus-links-go/internal/link/service"

type LinkHandler struct {
	service *service.LinkService
}
