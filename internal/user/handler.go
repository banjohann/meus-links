package user

import (
	"github.com/go-chi/chi"
)

type UserHandler struct {
	service *UserService
}

func NewHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) LoadUserRoutes() func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/", h.service.CreateUser)
	}
}
