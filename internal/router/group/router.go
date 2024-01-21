package group

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/dto"
	"github.com/Xenous-Inc/finapp-api/internal/router/constants"
	"github.com/go-chi/chi"
)

type Router struct {
	client *ruzfaclient.Client
}

func NewRouter(client *ruzfaclient.Client) *Router {
	return &Router{
		client:            client,
	}
}

func (s *Router) Route(r chi.Router) {
	r.Get("/", s.HandleGetGroup)
}

// @Summary GetGroup
// @Tags TimeTable
// @Description get group
// @Accept json
// @Produce json
// @Param input body dto.GetGroupsRequest
// @Success 200 {integer} integer 1
// @Router /finapp/api/group [post]
func (s *Router) HandleGetGroup(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetGroupsRequest{
		GroupTerm: r.URL.Query().Get(constants.QUERY_TERM),
	}

	groups, err := s.client.GetGroups(&ruzfaclient.GetGroupsInput{
		GroupTerm: input.GroupTerm,
	})

	if err != nil {
		//TODO: add error handling
	}

	res, err := json.Marshal(groups)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
	}

	fmt.Fprintf(w, string(res))
}