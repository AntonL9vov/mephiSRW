package service

import (
	"golang.org/x/crypto/bcrypt"
	"mephiSRW/pkg/model"
	"mephiSRW/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	cost := 10
	bPassword := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(bPassword, cost)
	return string(hash)
}
