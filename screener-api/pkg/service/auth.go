package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	m "github.com/CranesTeam/team-screener/pkg/model"
	r "github.com/CranesTeam/team-screener/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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
	password, err := GeneratehashPassword(userDto.Password)
	if err != nil {
		return "empty", err
	}
	userRoleId, err := s.repo.GetUserRoleId()
	if err != nil {
		return "empty", errors.New("couldn't find user role")
	}

	user := m.User{Username: userDto.Username, PasswordHash: password, RoleId: userRoleId}
	userInfo := m.UserInfo{Name: userDto.Name, Email: userDto.Email}

	uuid, err := s.repo.CreateUser(user, userInfo)

	logrus.Info("go to auth server...")
	resp, err := http.Get(fmt.Sprintf("%s/test?access_token=%s", authServerURL, globalToken.AccessToken))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	return uuid, err
}

func (s *AuthService) GenerateJWT(username, password string) (m.TokenResponse, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return m.TokenResponse{}, errors.New("could find user")
	}

	check := CheckPasswordHash(password, user.PasswordHash)
	if !check {
		return m.TokenResponse{}, errors.New("username or password is incorrect")
	}

	var mySigningKey = []byte(salt)
	token := jwt.New(jwt.SigningMethodHS256)
	time := time.Now().Add(tokenTTL).Unix()
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = user.RoleName
	claims["exp"] = time
	claims["user_uuid"] = user.ExternalUuid

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return m.TokenResponse{}, err
	}

	return m.TokenResponse{Role: user.RoleName, TokenString: tokenString, Exptime: time}, nil
}

// todo....
func (s *AuthService) ParseToken(accessToken string) (string, error) {
	var mySigningKey = []byte(salt)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in parsing token.")
		}
		return mySigningKey, nil
	})

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

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
