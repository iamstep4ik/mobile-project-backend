package lib

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/models"
	"go.uber.org/zap"
	"os"
	"time"
)

type JWTManager interface {
	GenerateJWT(user *models.User) (string, error)
}

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		log.Error("Failed to get JWT_SECRET from env")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Error("Failed to generate JWT: %v", zap.Error(err))
		return "", err
	}

	return tokenString, nil

}
