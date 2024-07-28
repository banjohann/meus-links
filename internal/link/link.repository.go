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

func (r *LinkRepo) FindByID(linkID string) (*Link, error) {
	link := Link{}

	err := r.db.Get(&link, "SELECT * FROM links WHERE id = $1", linkID)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *LinkRepo) FindByEncurtado(short string) (*Link, error) {
	link := Link{}

	err := r.db.Get(&link, "SELECT * FROM links WHERE short = $1", short)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *LinkRepo) GetAllLinksDoUsuario(userID string) ([]Link, error) {
	links := []Link{}

	err := r.db.Get(&links, "SELECT * FROM links WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	return links, nil
}

func (r *LinkRepo) RemoverByUsuarioID(userID string) {
	r.db.Exec("DELETE FROM links WHERE user_id = ($1)", userID)
}

func (r *LinkRepo) Update(link *Link) error {
	_, err := r.db.NamedExec(`UPDATE links SET nome = :nome, user_id = :user_id, short = :short, redirects_to = :redirects_to, clicks = :clicks WHERE id = :id`, link)

	return err
}
