package repository

import (
	"context"

	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type ProfileRepository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewProfileRepository(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{
		db:     db,
		logger: log.GetLogger(),
	}
}

func (r *ProfileRepository) CreateProfile(p *models.Profile) (*models.Profile, error) {
	r.logger.Info("Received profile data", zap.Any("profile", p))
	images := pq.Array(p.ImagesURL)
	interests := pq.Array(p.Interests)

	err := r.db.QueryRow(
		context.Background(),
		"INSERT INTO profiles (id, user_id, name, surname,imagesurl,description,gender,age,location,interests) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		p.ID,
		p.UserID,
		p.Name,
		p.Surname,
		images,
		p.Description,
		p.Gender,
		p.Age,
		p.Location,
		interests,
	).Scan(&p.ID)

	if err != nil {
		r.logger.Error("Failed to insert profile into database", zap.Error(err))
		return nil, err
	}
	return p, nil

}
