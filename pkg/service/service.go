package service

import "github.com/CranesTeam/team-screener/pkg/repository"

type Autorisation interface {
}

type Skill interface {
}

type Service struct {
	Autorisation
	Skill
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
