package teacher

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
	r.Post("/teacher1", s.HandleGetTeacher) // name_url?
}

// @Summary GetTeacher
// @Tags TimeTable
// @Description get teacher
// @Accept json
// @Produce json
// @Param input body ruzfaclient.GetTeacherInput true "teacher info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/teacher [post]
func (s *Router) HandleGetTeacher(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetTeacherRequest{
		TeacherTerm: r.URL.Query().Get(constants.QUERY_TERM),
	}

	teacher, err := s.client.GetTeacher(&ruzfaclient.GetTeacherInput{
		TeacherTerm: input.TeacherTerm,
	})

	if err != nil {
		//TODO: add error handling
	}

	res, err := json.Marshal(teacher)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
	}

	fmt.Fprintf(w, string(res))
}