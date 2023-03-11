package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"mephiSRW/pkg/model"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, login, email, phone ,password) VALUES($1, $2, $3, $4, $5) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Login, user.Email, user.Phone, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login = $1", userTable)
	err := r.db.Get(&user, query, login)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
