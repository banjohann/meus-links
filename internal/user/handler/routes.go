package handler

import (
	"github.com/JohannBandelow/meus-links-go/internal/user/repository"
	"github.com/go-chi/chi"
)

type UserHandler struct {
	repo *repository.UserRepo
}

func NewUserHandler(repo *repository.UserRepo) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (h *UserHandler) LoadUserRoutes() func(chi.Router) {
	return func(router chi.Router) {
		router.Get("/", h.CreateUser)
	}
}
