package repository

import (
	"fmt"

	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/jmoiron/sqlx"
)

type SkillsRepository struct {
	db *sqlx.DB
}

func NewSkillsRepository(db *sqlx.DB) *SkillsRepository {
	return &SkillsRepository{db: db}
}

func (r *SkillsRepository) CreateNewSkill(skill m.Skill) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "empty", err
	}

	var uuid string
	query := fmt.Sprintf("insert into %s (name, title, description) values ($1, $2, $3) returning external_uuid", skillsTable)
	row := tx.QueryRow(query, skill.Name, skill.Title, skill.Description)
	if err := row.Scan(&uuid); err != nil {
		tx.Rollback()
		return "empty", err
	}

	return uuid, tx.Commit()
}

func (r *SkillsRepository) GetAll() ([]m.Skill, error) {
	var skills []m.Skill
	query := fmt.Sprintf("select * from %s", skillsTable)
	err := r.db.Select(&skills, query)

	return skills, err
}

func (r *SkillsRepository) FindOne(uuid string) (m.Skill, error) {
	var skill m.Skill
	query := fmt.Sprintf("select * from %s where external_uuid=$1", skillsTable)
	err := r.db.Get(&skill, query, uuid)

	return skill, err
}
