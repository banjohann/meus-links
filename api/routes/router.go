package routes

import (
	api "github.com/JohannBandelow/meus-links-go/internal/adapters/http"
	"github.com/go-chi/chi"
)

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()

	linkHandler := api.LinkController{}
	userHandler := api.UserController{}

	router.Route("/api/users", userHandler.LoadRoutes())
	router.Route("/api/links", linkHandler.LoadRoutes())

	return router
}
