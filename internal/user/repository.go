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

func (r *UserRepo) Save(user User) error {
	_, err := r.db.NamedExec("INSERT INTO users (id, nome, sobrenome, email, senha) VALUES (:id, :nome, :sobrenome, :email, :senha)", user)

	return err
}

func (r *UserRepo) Delete(userID string) {

	_, err := r.db.Exec("DELETE FROM users WHERE id = ($1)", userID)
	if err != nil {
		log.Println(err)
	}
}

func (r *UserRepo) Get(userID string) (*User, error) {
	user := User{}

	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) FindByEmail(email string) *User {
	user := User{}

	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil
	}

	return &user
}

func (r *UserRepo) Update(user *User) error {
	_, err := r.db.NamedExec("UPDATE users SET nome = :nome, sobrenome = :sobrenome, email = :email, senha = :senha WHERE id = :id", user)

	return err
}
