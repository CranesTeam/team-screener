package repository

type Autorisation interface {
}

type Skill interface {
}

type Repository struct {
	Autorisation
	Skill
}

func NewRepository() *Repository {
	return &Repository{}
}
