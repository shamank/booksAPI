package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shamank/booksAPI/models"
	"github.com/shamank/booksAPI/pkg/repository"
	"time"
)

const (
	salt = "sjgbwerslfaifpluarn"

	tokenTTL      = 6 * time.Hour
	jwtSigningKey = "asd09j123asdko@#adskolDAS"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = hashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {

	user, err := s.repo.GetUser(username, password)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.ID,
	})

	tokenString, err := token.SignedString([]byte(jwtSigningKey))
	if err != nil {
		fmt.Println(tokenString, err)
		return "", err
	}

	return tokenString, nil
}

func hashPassword(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
