package models

import (
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, user.Username, user.PasswordHash).Scan(&user.ID)
}

func (r *UserRepository) FindByUsername(username string) (*User, error) {

	var user User
	err := r.db.Get(&user, "SELECT id, username, password_hash From users WHERE username=$1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
