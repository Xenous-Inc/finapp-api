package classrooms

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
	r.Get("/", s.HandleGetClassRooms)
}

// @Summary GetAuditorium
// @Tags TimeTable
// @Description get auditorium
// @Accept json
// @Produce json
// @Param term query string true "Audience search mask"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auditorium [get]
func (s *Router) HandleGetClassRooms(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetClassRoomRequest{
		ClassRoomTerm: r.URL.Query().Get(constants.QUERY_TERM),
	}
	auditorium, err := s.client.GetAuditorium(&ruzfaclient.GetAuditoriumInput{
		AuditoriumTerm: input.ClassRoomTerm,
	})

	if err != nil {
		//TODO: add error handling
	}

	res, err := json.Marshal(auditorium)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
	}

	fmt.Fprintf(w, string(res))
}