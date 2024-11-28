package configs

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewPostgresConnection() *sqlx.DB {
	credentials := LoadConfig()

	dbUser := credentials.DBUser
	if dbUser == "" {
		panic(errors.New("missing DB_USER property"))
	}

	password := credentials.DBPassword
	if password == "" {
		panic(errors.New("missing DB_PASSWORD property"))
	}

	dbName := credentials.DBName
	if credentials.DBName == "" {
		panic(errors.New("missing DB_NAME property"))
	}

	port := credentials.DBPort
	if port == "" {
		panic(errors.New("missing DB_PORT property"))
	}

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, password, dbName, port)

	db := sqlx.MustConnect("postgres", dsn)

	err := ExecMigrations(db)
	if err != nil {
		fmt.Println("Error running migrations")
		panic(err)
	}

	return db
}
