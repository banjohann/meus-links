package main

import (
	"context"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	cmd "github.com/JohannBandelow/meus-links-go/cmd/app"
	"github.com/JohannBandelow/meus-links-go/internal/user"
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

	userRepo := user.NewRepo(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	app.WithUserHandler(userHandler)

	app.Run(context.TODO())
}
