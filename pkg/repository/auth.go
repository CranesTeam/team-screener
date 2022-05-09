package repository

import (
	"fmt"

	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user m.User) (string, error) {
	var uuid string

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) returning external_uuid", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.PasswordHash)
	if err := row.Scan(&uuid); err != nil {
		return "empty", err
	}

	return uuid, nil
}
