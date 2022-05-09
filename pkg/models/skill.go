package models

/* Skills */
type Skill struct {
	Id          int64  `db:"id"`
	ExternalId  string `db:"external_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type SkillDto struct {
	Id          string `json:"external_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type SkillRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

/** Skill list */
type SkillList struct {
	UserId string `db:"external_id"`
}

/** Skill pointer */
type SkillPointer struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Point       int    `json:"point"`
}
