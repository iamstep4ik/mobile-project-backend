package models

import "github.com/google/uuid"

type Profile struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	ImagesURL   []string  `json:"images_url"`
	Description string    `json:"description"`
	Gender      string    `json:"gender"`
	Age         int       `json:"age"`
	Location    string    `json:"location"`
	Interests   []string  `json:"interests"`
}

type ProfileUseCase interface {
	FillProfile(p *Profile) (*Profile, error)
	EditProfile(p *Profile) (*Profile, error)
}

type ProfileRepository interface {
	CreateProfile(p *Profile) (*Profile, error)
	DeleteProfile(p *Profile) (bool, error)
	UpdateProfile(p *Profile) (*Profile, error)
}
