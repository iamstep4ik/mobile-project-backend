package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/config"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/handlers"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/repository"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/routes"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/service"
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
		log.Error("Error loading .env file", zap.Error(err))
	}

	cfg := config.NewConfig()
	if err = config.LoadConfig(cfg); err != nil {
		log.Fatal("Failed to load config", zap.Error(err))
	}

	db, err := repository.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()
	log.Info("Successfully connected to database")

	repo := repository.NewUserRepository(db)
	userUseCase := service.NewUserUseCase(repo)
	userHandler := handlers.NewUserHandler(userUseCase)

	router := gin.Default()

	routes.RegisterRoutes(router, userHandler)

	log.Info("Server started on port " + cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
	}
}
