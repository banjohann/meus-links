package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/JohannBandelow/meus-links-go/internal/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

type App struct {
	router      *chi.Mux
	port        int
	db          *gorm.DB
	userHandler *user.UserHandler
}

// Injeta tudo que é necessário para rodar o App
func NewApp(port int, db *gorm.DB, router *chi.Mux) *App {

	router.Use(middleware.Logger)

	return &App{
		router: router,
		port:   3030,
		db:     db,
	}
}

func (a *App) Run(ctx context.Context) error {
	log.Println("Starting Server")

	ch := make(chan error, 1)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", a.port),
		Handler: a.router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return server.Shutdown(ctx)

	case err := <-ch:
		return err
	}
}

func (a *App) WithUserHandler(userHandler *user.UserHandler) {
	a.userHandler = userHandler

	a.router.Route("/users", a.userHandler.LoadUserRoutes())
}
