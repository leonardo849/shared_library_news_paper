package dto

type AuthPublishUserCreated struct {
	AuthId string `json:"auth_id"`
	Username  string `json:"username"`
	Role string `json:"role"`
}