package migrations

import "github.com/jmoiron/sqlx"

func CreateUserTable(db *sqlx.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id uuid NOT NULL,
			nome text NOT NULL,
			sobrenome text NOT NULL,
			email text NOT NULL,
			senha text NOT NULL,
		CONSTRAINT uni_users_email UNIQUE (email),
		CONSTRAINT users_pkey PRIMARY KEY (id));
	`

	db.MustExec(query)
}
