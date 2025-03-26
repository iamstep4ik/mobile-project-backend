package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/lib"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/models"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/repository"
	"go.uber.org/zap"
)

type UserUseCase struct {
	repo     repository.UserRepository
	hasher   lib.PasswordHasher
	logger   *zap.Logger
	validate lib.Validation
}

func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repo:     *repo,
		hasher:   lib.NewPasswordHasher(),
		logger:   log.GetLogger(),
		validate: lib.NewValidator(),
	}
}

func (u *UserUseCase) SignUp(user *models.User) (*models.User, error) {
	u.logger.Info("Received user data", zap.Any("user", user))
	validateErr := u.validate.ValidateUser(user)
	if validateErr != nil {
		u.logger.Error("Validation error", zap.Error(validateErr))
		return nil, validateErr
	}
	user.ID = uuid.New()
	u.logger.Info("Generated UUID", zap.String("id", user.ID.String()))

	hashedPassword, err := u.hasher.HashPassword(user.HashedPassword)
	if err != nil {
		u.logger.Error("error hashing password", zap.Error(err))
		return nil, err
	}

	u.logger.Info("Password", zap.String("password", user.HashedPassword))

	user.HashedPassword = hashedPassword

	exists, err := u.repo.IsUserExists(user.Email)
	if err != nil {
		u.logger.Error("error getting user by email", zap.Error(err))
		return nil, err
	}
	if exists {
		u.logger.Error("User with such email already exists", zap.Any("Email", user.Email))
		return nil, fmt.Errorf("user with such email already exists")
	}

	_, err = u.repo.CreateUser(user)
	if err != nil {
		u.logger.Error("error creating user", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) Login(email, password string) (*models.User, string, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		u.logger.Error("error getting user by email", zap.Error(err))
		return nil, "", err
	}

	if !u.hasher.CheckPasswordHash(password, user.HashedPassword) {
		u.logger.Error("passwords do not match")
		return nil, "", fmt.Errorf("passwords do not match")
	}

	u.logger.Info("User logged in", zap.Any("user", user))

	token, err := lib.GenerateJWT(user)
	if err != nil {
		u.logger.Error("error generating JWT", zap.Error(err))
		return nil, "", err
	}
	u.logger.Info("Generated JWT", zap.String("token", token))

	return user, token, nil
}
