package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/shamank/booksAPI/models"
	"github.com/shamank/booksAPI/pkg/repository"
)

const salt = "sjgbwerslfaifpluarn"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: *repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = hashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func hashPassword(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
