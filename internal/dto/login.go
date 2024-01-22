package dto

type LoginRequest struct {
	Login    string `json:"username" example:"227789" validate:"required"`
	Password string `json:"password" example:"Bibt8877" validate:"required"`
} // @name LoginRequest

type LoginResponse struct {
	AccessToken string `json:"token" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9"`
} // @name LoginResponse

