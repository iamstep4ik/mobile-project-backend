package repository

import (
	"context"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(signUp *models.User) (*models.User, error) {
	r.db.QueryRow(
		context.Background(),
		"INSERT INTO users (name, surname, email, hashed_password) VALUES ($1, $2, $3, $4) RETURNING id",
		signUp.Name,
		signUp.Surname,
		signUp.Email,
		signUp.HashedPassword,
	).Scan(&signUp.ID)
	return signUp, nil
}
