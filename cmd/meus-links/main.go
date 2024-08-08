package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"

	"github.com/JohannBandelow/meus-links-go/configs"

	_ "github.com/lib/pq"
)

type Application struct {
	Router *chi.Mux
	Port   int
	DB     *sqlx.DB
}

func main() {
	db := configs.NewDBConnection()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	app := Application{
		Port:   3030,
		DB:     db,
		Router: router,
	}

	Run(context.TODO(), app)
}

func Run(ctx context.Context, app Application) error {
	log.Printf("Starting Server at port %d", app.Port)

	ch := make(chan error, 1)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", app.Port),
		Handler: app.Router,
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
