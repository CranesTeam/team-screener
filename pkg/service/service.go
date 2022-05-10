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
	CreateNewSkill(skill m.SkillRequest) (string, error)
	GetAll() ([]m.SkillDto, error)
	FindOne(name string) (m.SkillDto, error)
}

type UserSkills interface {
	GetUserSkills(user_uuid string) (m.SkillListDto, error)
	AddNewSkillPointer(user_uuid string, skillRequest m.AddSkillRequest) (string, error)
}

type Service struct {
	Authorization
	Skills
	UserSkills
}

func NewService(repo *r.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Skills:        NewSkillService(repo.Skills),
		UserSkills:    NewUserSkillsService(repo.UserSkills),
	}
}
