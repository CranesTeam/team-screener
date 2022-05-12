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
		userSkillsDto = append(userSkillsDto, convert(entity))
	}

	return m.SkillListDto{UserId: user_uuid, SkillPointers: userSkillsDto}, err
}

func (s *UserSkillsService) AddNewSkillPointer(user_uuid string, skillRequest m.AddSkillRequest) (string, error) {
	return s.repo.AddNewSkill(user_uuid, skillRequest.SkillUuid, skillRequest.Point)
}

func (s *UserSkillsService) FindSkill(user_uuid string, skill_uuid string) (m.UserSkillsDto, error) {
	entity, err := s.repo.FindSkill(user_uuid, skill_uuid)
	skillDto := convert(entity)

	return skillDto, err
}

func (s *UserSkillsService) DeleteSkill(user_uuid string, skill_uuid string) (string, error) {
	return s.repo.DeleteSkill(user_uuid, skill_uuid)
}

func (s *UserSkillsService) UpdatePoint(user_uuid string, skill_uuid string, points string) (string, error) {
	return s.repo.UpdatePoint(user_uuid, skill_uuid, points)
}

func convert(entity m.UserSkills) m.UserSkillsDto {
	return m.UserSkillsDto{
		ExternaUuid: entity.ExternaUuid,
		Name:        entity.Name,
		Title:       entity.Title,
		Description: entity.Description,
		Points:      entity.Points,
	}
}
