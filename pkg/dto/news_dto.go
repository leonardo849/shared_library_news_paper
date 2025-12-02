package dto

type UserCreatedErrorDTO struct {
	AuthID string `json:"id"`
	Message string `json:"message"`
}

type UserCreatedDTO struct {
	AuthID string `json:"id"`
	Message string `json:"message"`
}