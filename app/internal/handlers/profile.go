package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/lib"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/log"
	"github.com/iamstep4ik/mobile-project-backend/app/internal/models"
	"go.uber.org/zap"
)

type ProfileHandler struct {
	useCase models.ProfileUseCase
	logger  *zap.Logger
}

func NewProfileHandler(useCase models.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{
		useCase: useCase,
		logger:  log.GetLogger(),
	}
}

func (h *ProfileHandler) FillProfile(c *gin.Context) {
	h.logger.Info("Received request to fill profile")

	profileJSON := c.PostForm("profile")
	if profileJSON == "" {
		h.logger.Error("missing profile form field")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing profile data"})
		return
	}

	var profile models.Profile
	if err := json.Unmarshal([]byte(profileJSON), &profile); err != nil {
		h.logger.Error("failed to parse profile JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid profile data"})
		return
	}

	files := c.Request.MultipartForm.File["profile_photos[]"]
	var imagePaths []string

	for _, file := range files {
		userID := profile.UserID.String()
		filePath, err := lib.SaveUploadedFile(file, userID)
		if err != nil {
			h.logger.Error("failed to save profile photo", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		imagePaths = append(imagePaths, filePath)
	}
	profile.ImagesURL = imagePaths

	createdProfile, err := h.useCase.FillProfile(&profile)
	if err != nil {
		h.logger.Error("error creating profile", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("Profile created successfully", zap.Any("profile", createdProfile))
	c.JSON(http.StatusCreated, createdProfile)
}
