package dto

type NewsUserCreatedErrorDTO struct {
	AuthID string `json:"id"`
	Message string `json:"message"`
}

type NewsUserCreatedDTO struct {
	AuthID string `json:"id"`
	Message string `json:"message"`
}