package models

type Profile struct {
	Name        string   `json:"name"`
	Surname     string   `json:"surname"`
	ImagesURL   []string `json:"images_url"`
	Description string   `json:"description"`
}

type ProfileUseCase interface {
}

type ProfileRepository interface {
}
