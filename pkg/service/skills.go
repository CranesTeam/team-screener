package service

import (
	m "github.com/CranesTeam/team-screener/pkg/model"
	r "github.com/CranesTeam/team-screener/pkg/repository"
	"github.com/sirupsen/logrus"
)

type SkillService struct {
	repo r.Skills
}

func NewSkillService(repo r.Skills) *SkillService {
	return &SkillService{repo: repo}
}

func (s *SkillService) CreateNewSkill(skill m.SkillRequest) (string, error) {
	entity := m.Skill{
		Title:       skill.Title,
		Name:        skill.Name,
		Description: skill.Description,
	}

	return s.repo.CreateNewSkill(entity)
}

func (s *SkillService) GetAll() ([]m.SkillDto, error) {
	skills, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	logrus.Info(skills)
	var skillsDto []m.SkillDto

	for _, entity := range skills {
		skillsDto = append(skillsDto, m.SkillDto{
			Uuid:        entity.ExternalId,
			Name:        entity.Name,
			Title:       entity.Title,
			Description: entity.Description,
		})
	}

	return skillsDto, nil
}

func (s *SkillService) FindOne(uuid string) (m.SkillDto, error) {
	entity, err := s.repo.FindOne(uuid)
	dto := m.SkillDto{
		Uuid:        entity.ExternalId,
		Title:       entity.Title,
		Name:        entity.Name,
		Description: entity.Description,
	}

	return dto, err
}
