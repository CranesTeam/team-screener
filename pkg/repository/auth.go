package repository

import (
	"fmt"

	"github.com/CranesTeam/team-screener/pkg/model"
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
	tx, err := r.db.Begin()
	if err != nil {
		return "empty", err
	}

	var uuid string
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) returning external_uuid", usersTable)
	row := tx.QueryRow(query, user.Name, user.Username, user.PasswordHash)
	if err := row.Scan(&uuid); err != nil {
		tx.Rollback()
		return "empty", err
	}

	return uuid, tx.Commit()
}

func (r *AuthRepository) GetUser(username, password string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("select * from %s where username=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
