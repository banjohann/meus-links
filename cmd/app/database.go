package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/JohannBandelow/meus-links-go/cmd/migrations"
	"github.com/jmoiron/sqlx"
)

func createDatabase(db *sqlx.DB) {
	migrations.CreateUserTable(db)
	migrations.CreateLinkTable(db)
}

func NewDBConnection() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		panic(errors.New("missing DB_USER property"))
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		panic(errors.New("missing DB_PASSWORD property"))
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		panic(errors.New("missing DB_NAME property"))
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		panic(errors.New("missing DB_PORT property"))
	}

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, password, dbname, port)

	db := sqlx.MustConnect("postgres", dsn)

	createDatabase(db)

	return db
}
