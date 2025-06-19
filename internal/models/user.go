package models

import (
	"errors"
	"strings"
)

type User struct {
	ID           int64  `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(u.PasswordHash) == "" {
		return errors.New("password hash is required")
	}
	return nil
}
