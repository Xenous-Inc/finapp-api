package classrooms

import (
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/dto"
	"github.com/Xenous-Inc/finapp-api/internal/router/constants"
	"github.com/Xenous-Inc/finapp-api/internal/router/utils/responser"
	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
	"github.com/go-chi/chi"
	"gopkg.in/go-playground/validator.v9"
)

type Router struct {
	client    *ruzfaclient.Client
	validator *validator.Validate
}

func NewRouter(client *ruzfaclient.Client) *Router {
	return &Router{
		client:    client,
		validator: validator.New(),
	}
}

func (s *Router) Route(r chi.Router) {
	r.Get("/", s.HandleGetClassRooms)
}

// @Summary Return List of classrooms
// @Description Return list of classrooms which found by provided query term
// @Tags classrooms
// @Param term query string true "Classroom search mask"
// @Produce json
// @Success 200 {object} []dto.Classroom
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /classrooms/ [get]
func (s *Router) HandleGetClassRooms(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get(constants.QUERY_TERM)
	err := s.validator.Var(term, "required")
	if err != nil {
		log.Error(err, "BadRequest", "classrooms HandleGetClassRooms")
		responser.BadRequset(w, r, "Query parameter `term` must be provided")

		return
	}

	response, err := s.client.GetAuditorium(&ruzfaclient.GetEntitiesInput{
		Term: term,
	})
	if err != nil {
		log.Error(err, "Internal", "classrooms HandleGetClassRooms")
		responser.Internal(w, r, err.Error())

		return
	}

	classrooms := make([]dto.Classroom, 0)
	for _, classroom := range response {
		if classroom.Id != "" {
			classrooms = append(classrooms, dto.Classroom{
				Id:          classroom.Id,
				Title:       classroom.Label,
				Description: classroom.Description,
			})
		}
	}

	responser.Success(w, r, classrooms)
}
