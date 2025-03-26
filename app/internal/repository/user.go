package repository

import (
	"context"

	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserRepository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: log.GetLogger(),
	}
}
func (r *UserRepository) CreateUser(signUp *models.User) (*models.User, error) {
	r.logger.Info("Received user data", zap.Any("user", signUp))

	err := r.db.QueryRow(
		context.Background(),
		"INSERT INTO users (id, name, surname, email, hashed_password) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		signUp.ID,
		signUp.Name,
		signUp.Surname,
		signUp.Email,
		signUp.HashedPassword,
	).Scan(&signUp.ID)

	if err != nil {
		r.logger.Error("Failed to insert user into database", zap.Error(err))
		return nil, err
	}

	return signUp, nil
}

func (r *UserRepository) IsUserExists(email string) (bool, error) {
	var exists bool
	r.db.QueryRow(
		context.Background(),
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)",
		email,
	).Scan(&exists)
	return exists, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(
		context.Background(),
		"SELECT id, name, surname, email, hashed_password FROM users WHERE email = $1",
		email,
	).Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.HashedPassword)

	if err != nil {
		r.logger.Error("Failed to get user by email", zap.Error(err))
		return nil, err
	}
	return user, nil
}
