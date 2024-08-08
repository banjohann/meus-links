package migrations

const CreateUserTable = `
	CREATE TABLE IF NOT EXISTS users (
		id uuid NOT NULL,
		nome varchar(255) NOT NULL,
		sobrenome varchar(255) NOT NULL,
		email varchar(255) NOT NULL,
		senha varchar(255) NOT NULL,
	CONSTRAINT uni_users_email UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id));`
