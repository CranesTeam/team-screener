package repository

import (
	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user m.User) (string, error)
	GetUser(username, password string) (m.User, error)
}

type Skills interface {
	CreateNewSkill(skill m.Skill) (string, error)
	GetAll() ([]m.Skill, error)
	FindOne(uuid string) (m.Skill, error)
}

type UserSkills interface {
	GetUserSkills(user_uuid string) ([]m.UserSkills, error)
	AddNewSkill(user_uuid string, skill_uuid string, point int) (string, error)
}

type Repository struct {
	Authorization
	Skills
	UserSkills
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Skills:        NewSkillsRepository(db),
		UserSkills:    NewUserSkillsRepository(db),
	}
}
