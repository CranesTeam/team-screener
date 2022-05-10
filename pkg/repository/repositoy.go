package repository

import (
	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user m.User) (string, error)
	GetUser(username, password string) (m.User, error)
}

type Skill interface {
}

type Repository struct {
	Authorization
	Skill
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
