package migrations

const CreateLinkTable = `
	CREATE TABLE IF NOT EXISTS links (
		id uuid NOT NULL,
		user_id uuid NOT NULL,
		nome varchar(255) NOT NULL,
		short varchar(255) NOT NULL,
		redirects_to varchar(255) NOT NULL,
		clicks INTEGER NOT NULL DEFAULT 0,
	CONSTRAINT uni_short_url UNIQUE (short),
	CONSTRAINT links_pkey PRIMARY KEY (id));`
