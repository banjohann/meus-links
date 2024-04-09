package cmd

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(user string, password string, dbname string, port string) *gorm.DB {

	if user == "" {
		panic(errors.New("missing DB_USER property"))
	}

	if password == "" {
		panic(errors.New("missing DB_PASSWORD property"))
	}

	if dbname == "" {
		panic(errors.New("missing DB_NAME property"))
	}

	if port == "" {
		panic(errors.New("missing DB_PORT property"))
	}

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// err = db.AutoMigrate(&user.User{})
	// if err != nil {
	// 	return nil, err
	// }

	return db
}
