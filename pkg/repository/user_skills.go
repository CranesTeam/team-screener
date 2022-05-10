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
	query := fmt.Sprintf("select st.external_uuid, st.name, st.title, st.description, ust.points from %s ust "+
		"inner join %s st on ust.skill_uuid = st.external_uuid "+
		"where ust.user_uuid=$1", userSkillsTable, skillsTable)
	err := r.db.Select(&userSkills, query, user_uuid)

	return userSkills, err
}

func (r *UserSkillRepository) AddNewSkill(user_uuid string, skill_uuid string, point int) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "empty", err
	}

	var uuid string
	query := fmt.Sprintf("INSERT INTO %s (user_uuid, skill_uuid, points) values ($1, $2, $3) returning external_uuid", userSkillsTable)
	row := tx.QueryRow(query, user_uuid, skill_uuid, point)
	if err := row.Scan(&uuid); err != nil {
		tx.Rollback()
		return "empty", err
	}

	return uuid, tx.Commit()

}
