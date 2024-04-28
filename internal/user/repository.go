package user

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Save(user User) {
	tx := r.db.MustBegin()

	tx.NamedExec("INSERT INTO users (id, name, last_name, email, password) VALUES (:id, :nome, :sobrenome, :email, :senha)", user)
	tx.Commit()
}

func (r *UserRepo) Delete(userID string) {

	_, err := r.db.Exec("DELETE FROM users WHERE id = ($1)", userID)
	if err != nil {
		log.Println(err)
	}
}

func (r *UserRepo) Get(userID string) (User, error) {
	user := User{}

	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		return User{}, err
	}

	return User{}, nil
}

func (repo *UserRepo) Update(user User) {
	tx := repo.db.MustBegin()

	tx.NamedExec("UPDATE users SET name = :nome, sobrenome = :sobrenome, email = :email, password = :password WHERE id = :id", user)
	tx.Commit()
}
