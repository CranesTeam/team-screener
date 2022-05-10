package service

import (
	m "github.com/CranesTeam/team-screener/pkg/model"
	r "github.com/CranesTeam/team-screener/pkg/repository"
)

type UserSkillsService struct {
	repo r.UserSkills
}

func NewUserSkillsService(repo r.UserSkills) *UserSkillsService {
	return &UserSkillsService{repo: repo}
}

func (s *UserSkillsService) GetUserSkills(user_uuid string) (m.SkillListDto, error) {
	skills, err := s.repo.GetUserSkills(user_uuid)
	var userSkillsDto []m.UserSkillsDto

	for _, entity := range skills {
		userSkillsDto = append(userSkillsDto, m.UserSkillsDto{
			Name:        entity.Name,
			Title:       entity.Title,
			Description: entity.Description,
			Points:      entity.Points,
		})
	}

	return m.SkillListDto{UserId: user_uuid, SkillPointers: userSkillsDto}, err
}

func (s *UserSkillsService) AddNewSkillPointer(user_uuid string, skillRequest m.AddSkillRequest) (string, error) {
	return s.repo.AddNewSkill(user_uuid, skillRequest.SkillUuid, skillRequest.Point)
}
