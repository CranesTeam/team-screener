package repository

import (
	"fmt"

	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/jmoiron/sqlx"
)

type UserSkillRepository struct {
	db *sqlx.DB
}

func NewUserSkillsRepository(db *sqlx.DB) *UserSkillRepository {
	return &UserSkillRepository{db: db}
}

func (r *UserSkillRepository) GetUserSkills(user_uuid string) ([]m.UserSkills, error) {
	var userSkills []m.UserSkills
	query := fmt.Sprintf("select * from %s where user_uuid=$1", userSkillsTable)
	err := r.db.Select(&userSkills, query, user_uuid)

	return userSkills, err
}
