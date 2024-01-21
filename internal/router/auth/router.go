package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/dto"
	"github.com/go-chi/chi"
)

type Router struct {
	client *orgfaclient.Client
}

func NewRouter(client *orgfaclient.Client) *Router {
	return &Router{
		client:            client,
	}
}

func (s *Router) Route(r chi.Router) {
	r.Post("/login", s.HandleAuth)
}

// @Summary Auth
// @Tags OrgFaRu
// @Description auth
// @Accept json
// @Produce json
// @Param input body dto.LoginRequest
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth [post]
func (s *Router) HandleAuth(w http.ResponseWriter, r *http.Request) {
	var input *dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sessionId, err := s.client.Login(&orgfaclient.LoginInput{
		Login:    input.Login,
		Password: input.Password,
	})

	if err != nil {
		fmt.Fprintf(w, "auth:  %s", clients.ErrRequest)
	}

	res, err := json.Marshal(sessionId)

	if err != nil {
		fmt.Fprintf(w, "auth marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}
