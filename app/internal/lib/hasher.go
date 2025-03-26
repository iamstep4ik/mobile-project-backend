package lib

import (
	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type hasher struct {
	logger *zap.Logger
}

func NewPasswordHasher() PasswordHasher {
	return &hasher{logger: log.GetLogger()}
}

func (h *hasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		h.logger.Fatal("Failed to hash password", zap.Error(err))
	}
	return string(bytes), nil
}

func (h *hasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
