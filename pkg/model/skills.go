package model

/* Skills */
type Skill struct {
	Id          int64  `db:"id"`
	ExternalId  string `db:"external_uuid"`
	Name        string `db:"name"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type SkillDto struct {
	Uuid        string `json:"external_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type SkillRequest struct {
	Name        string `json:"name" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

/** Skill list */
type UserSkills struct {
	ExternaUuid string `db:"external_uuid"`
	Name        string `db:"name"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Points      int    `db:"points"`
}

type UserSkillsDto struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Points      int    `json:"points"`
}

type SkillListDto struct {
	UserId        string          `json:"user_uuid"`
	SkillPointers []UserSkillsDto `json:"skills"`
}

type AddSkillRequest struct {
	SkillUuid string `json:"skill_uuid"`
	Point     int    `json:"point"`
}
