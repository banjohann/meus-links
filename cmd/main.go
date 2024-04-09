package main

import (
	"context"
	"os"

	cmd "github.com/JohannBandelow/meus-links-go/cmd/app"
	"github.com/JohannBandelow/meus-links-go/internal/user/handler"
	"github.com/JohannBandelow/meus-links-go/internal/user/repository"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db := cmd.NewDBConnection(dbUser, dbPass, dbPort, dbName)

	router := chi.NewRouter()
	app := cmd.NewApp(
		3030,
		db,
		router,
	)

	userRepo := repository.NewUserRepo(db)
	app.WithUserHandler(handler.NewUserHandler(userRepo))

	app.Run(context.TODO())
}
