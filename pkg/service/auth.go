package service

import (
	"crypto/sha256"
	"fmt"

	m "github.com/CranesTeam/team-screener/pkg/model"
	r "github.com/CranesTeam/team-screener/pkg/repository"
)

const salt = "team-screener-code"

type AuthService struct {
	repo r.Authorization
}

func NewAuthService(repo r.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(userDto m.UserDto) (string, error) {
	user := m.User{
		Name:         userDto.Name,
		Username:     userDto.Username,
		PasswordHash: s.generateHash(userDto.Password),
	}
	return s.repo.CreateUser(user)
}

func (s *AuthService) generateHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
