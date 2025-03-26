package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/models"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	useCase models.UserUseCase
	logger  *zap.Logger
}

func NewUserHandler(useCase models.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: useCase,
		logger:  log.GetLogger(),
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		h.logger.Error("error binding json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.useCase.SignUp(&user)
	if err != nil {
		h.logger.Error("error signing up user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
func (h *UserHandler) Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		h.logger.Error("error binding json", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, _, err := h.useCase.Login(user.Email, user.HashedPassword)
	if err != nil {
		h.logger.Error("error signing up user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
