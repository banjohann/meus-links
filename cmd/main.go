package main

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	cmd "github.com/JohannBandelow/meus-links-go/cmd/app"
	"github.com/JohannBandelow/meus-links-go/internal/user"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	db := cmd.NewDBConnection()
	router := chi.NewRouter()

	app := cmd.NewApp(
		3030,
		db,
		router,
	)

	userRepo := user.NewRepo(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	app.WithUserHandler(userHandler)

	app.Run(context.TODO())
}
