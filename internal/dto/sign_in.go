package dto

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInResponse struct {
	ErrorResponse
	Token string `json:"accessToken"`
}
