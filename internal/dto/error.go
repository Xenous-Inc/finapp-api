package dto

type ErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}
