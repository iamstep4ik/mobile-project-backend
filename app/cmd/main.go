package main

import (
	"github.com/iamstep4ik/mobile-project-backend/app/internal/config"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/repository"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := log.Initialize()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer log.Sync()
	if err = godotenv.Load(); err != nil {
		log.Error("Error loading .env file: %v", zap.Error(err))
	}

	cfg := config.NewConfig()
	if err = config.LoadConfig(cfg); err != nil {
		log.Fatal("Failed to load config: %v", zap.Error(err))
	}
	db, err := repository.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database: %v", zap.Error(err))
	}
	defer db.Close()
	log.Info("Successfully connected to database")

}
