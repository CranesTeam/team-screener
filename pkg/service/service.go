package service

import (
	m "github.com/CranesTeam/team-screener/pkg/model"
	r "github.com/CranesTeam/team-screener/pkg/repository"
)

type Authorization interface {
	CreateUser(user m.UserDto) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Skills interface {
}

type UserSkills interface {
}

type Service struct {
	Authorization
	Skills
	UserSkills
}

func NewService(repo *r.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
