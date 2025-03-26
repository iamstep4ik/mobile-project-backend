package lib

import (
	"errors"
	"regexp"

	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/models"
	"go.uber.org/zap"
)

type Validation interface {
	ValidateUser(user *models.User) error
}

type validator struct {
	logger *zap.Logger
}

func NewValidator() *validator {
	return &validator{
		logger: log.GetLogger(),
	}
}
func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}

func (v *validator) ValidateUser(user *models.User) error {
	v.logger.Info("Validating user", zap.Any("user", user))
	if user.Name == "" {
		v.logger.Error("User name is empty")
		return errors.New("user name is empty")
	}

	if user.Surname == "" {
		v.logger.Error("User surname is empty")
		return errors.New("user surname is empty")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		v.logger.Error("invalid email format")
		return errors.New("invalid email format")
	}
	if !isValidPassword(user.HashedPassword) {
		v.logger.Error("invalid format password")
		return errors.New("invalid password format")
	}

	v.logger.Info("User validation successful", zap.Any("user", user))
	return nil
}
