package link

import (
	"github.com/jmoiron/sqlx"
)

type LinkRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *LinkRepo {
	return &LinkRepo{
		db: db,
	}
}

func (r *LinkRepo) Save(link Link) error {
	_, err := r.db.NamedExec("INSERT INTO links (id, nome, user_id, short, redirects_to, clicks) VALUES (:id, :nome, :user_id, :short, :redirects_to, :clicks)", link)

	return err
}

func (r *LinkRepo) Delete(linkID string) error {

	_, err := r.db.Exec("DELETE FROM links WHERE id = ($1)", linkID)
	return err
}

func (r *LinkRepo) Get(linkID string) (*Link, error) {
	link := Link{}

	err := r.db.Get(&link, "SELECT * FROM links WHERE id = $1", linkID)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *LinkRepo) GetByUser(userID string) *Link {
	user := Link{}

	err := r.db.Get(&user, "SELECT * FROM links WHERE user_id = $1", userID)
	if err != nil {
		return nil
	}

	return &user
}

func (r *LinkRepo) Update(link *Link) error {
	_, err := r.db.NamedExec(`UPDATE links SET nome = :nome, user_id = :user_id, short = :short, redirects_to = :redirects_to, clicks = :clicks WHERE id = :id`, link)

	return err
}
