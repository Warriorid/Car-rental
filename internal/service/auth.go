package service

import (
	"CarRentalService/internal/models"
	"CarRentalService/internal/repository"
	"CarRentalService/pkg/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	tokenTLL = 24 * time.Hour
	signingKey = "2o1gs32p9d2igcvwyeoc"
)
type castomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{ repo: repo }
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = config.GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email string, password string) (string, error) {
	user, err := s.repo.GetUser(email, config.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &castomClaims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTLL)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &castomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("infalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*castomClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}
	return claims.UserId, nil
}