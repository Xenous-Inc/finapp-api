package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/dto"
	"github.com/Xenous-Inc/finapp-api/internal/router/utils/responser"
	"github.com/Xenous-Inc/finapp-api/internal/utils/jwtservice"
	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
	"github.com/go-chi/chi"
	"gopkg.in/go-playground/validator.v9"
)

type Router struct {
	client    *orgfaclient.Client
	validator *validator.Validate
}

func NewRouter(client *orgfaclient.Client) *Router {
	return &Router{
		client:    client,
		validator: validator.New(),
	}
}

func (s *Router) Route(r chi.Router) {
	r.Post("/login", s.HandleAuth)
}

// @Summary Try to sign in user
// @Description In success case returns Access JWT Token
// @Tags auth
// @Param data body dto.LoginRequest true "Credentials input"
// @Produce json
// @Success 200 {object} dto.LoginResponse
// @Failure 401 {object} dto.ApiError
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /auth/login [post]
func (s *Router) HandleAuth(w http.ResponseWriter, r *http.Request) {
	input := new(dto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.validator.Struct(input)
	if err != nil {
		responser.BadRequset(w, r, "Password and login must be provided")

		return
	}

	sessionId, err := s.client.Login(&orgfaclient.LoginInput{
		Login:    input.Login,
		Password: input.Password,
	})

	token, err := jwtservice.NewToken(sessionId, s.client.Cfg.JwtSecret)
	if err != nil {
		log.Warn("Unauthorized, Error generate Token", "orgfaclient login")
	}

	if err != nil {
		if errors.Is(err, clients.ErrUnauthorized) {
			responser.Unauthorized(w, r)

			return
		}

		responser.Internal(w, r, err.Error())

		return
	}

	response := &dto.LoginResponse{
		AccessToken: token,
	}

	responser.Success(w, r, response)
}
