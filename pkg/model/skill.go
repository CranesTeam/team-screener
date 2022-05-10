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
	Name        string `db:"name"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Points      int    `db:"point"`
}

type SkillListDto struct {
	UserId        string            `json:"user_uuid"`
	SkillPointers []SkillPointerDto `json:"skills"`
}

type SkillPointerDto struct {
	Uuid        string `json:"uuid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Point       int    `json:"point"`
}
