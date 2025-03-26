package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/handlers"
)

func RegisterRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	auth := router.Group("/auth")
	auth.POST("/signup", userHandler.SignUp)
	auth.POST("/login", userHandler.Login)
}
