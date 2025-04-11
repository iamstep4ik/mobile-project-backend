package service

import (
	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/models"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/repository"
	"go.uber.org/zap"
)

type ProfileUseCase struct {
	repo   repository.ProfileRepository
	logger *zap.Logger
}

func NewProfileUseCase(repo *repository.ProfileRepository) *ProfileUseCase {
	return &ProfileUseCase{
		repo:   *repo,
		logger: log.GetLogger(),
	}
}

func (p *ProfileUseCase) FillProfile(profile *models.Profile) (*models.Profile, error) {
	p.logger.Info("Received profile data", zap.Any("profile", profile))

	createdProfile, err := p.repo.CreateProfile(profile)
	if err != nil {
		p.logger.Error("error creating profile", zap.Error(err))
		return nil, err
	}

	return createdProfile, nil
}
func (p *ProfileUseCase) EditProfile(profile *models.Profile) (*models.Profile, error) {
	return nil, nil
}
