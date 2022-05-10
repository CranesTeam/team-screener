package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	m "github.com/CranesTeam/team-screener/pkg/model"
	r "github.com/CranesTeam/team-screener/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt         = "team-screener-code" // todo: registation time, use as secret
	tokenTTL     = 1 * time.Hour        // todo: to config
	signatureKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserUuid string `json:"user_uuid"`
}

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
		PasswordHash: generateHash(userDto.Password),
	}
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generateHash(password))
	if err != nil {
		return "empty", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ExternalUuid,
	})

	return token.SignedString([]byte(signatureKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signatureKey), nil
	})
	if err != nil {
		return "empty", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "empty", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserUuid, nil
}

func generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
