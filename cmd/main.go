package main

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	cmd "github.com/JohannBandelow/meus-links-go/cmd/app"
	"github.com/JohannBandelow/meus-links-go/internal/user"
	user_api "github.com/JohannBandelow/meus-links-go/internal/user/api"
	user_service "github.com/JohannBandelow/meus-links-go/internal/user/service"

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
	userService := user_service.New(userRepo)
	userHandler := user_api.New(userService)
	app.WithUserHandler(userHandler)

	app.Run(context.TODO())
}
