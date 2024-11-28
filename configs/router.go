package configs

import (
	api "github.com/JohannBandelow/meus-links-go/internal/adapters/http"
	"github.com/JohannBandelow/meus-links-go/internal/adapters/sql"
	"github.com/JohannBandelow/meus-links-go/internal/service"
	"github.com/JohannBandelow/meus-links-go/internal/usecase/link"
	"github.com/JohannBandelow/meus-links-go/internal/usecase/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(dbConn *sqlx.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	linkRepo := sql.NewLinkRepo(dbConn)
	userRepo := sql.NewUserRepo(dbConn)
	userService := service.UserService{}

	linkHandler := api.LinkController{
		CriaLinkUseCase:     link.CriaLinkUseCase{Repo: linkRepo, UserService: userService},
		AtualizaLinkUseCase: link.AtualizaLinkUseCase{Repo: linkRepo},
		RemoverLinkUseCase:  link.RemoverLinkUseCase{Repo: linkRepo},
		GetByIdUseCase:      link.GetByIdLinkUseCase{Repo: linkRepo},
		GetAllUseCase:       link.GetAllLinkUseCase{Repo: linkRepo},
	}

	userHandler := api.UserController{
		CriaUsuarioUseCase:     user.CriarUsuarioUseCase{Repo: userRepo},
		AtualizaUsuarioUseCase: user.AtualizaUsuarioUseCase{Repo: userRepo},
		GetUsuarioUseCase:      user.GetUsuarioByIdUseCase{Repo: userRepo},
		RemoveUsuarioUseCase:   user.RemoveUsuarioUseCase{Repo: userRepo},
	}

	router.Route("/api/users", userHandler.LoadRoutes())
	router.Route("/api/links", linkHandler.LoadRoutes())

	return router
}
