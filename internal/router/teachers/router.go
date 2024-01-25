package teachers

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
	r.Get("/", s.HandleGetTeacher)
}

// @Summary Return List of teachers
// @Description Return list of teachers which found by provided query term
// @Tags teachers
// @Param term query string true "Teacher search mask"
// @Produce json
// @Success 200 {object} []dto.Teacher
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /teachers/ [get]
func (s *Router) HandleGetTeacher(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get(constants.QUERY_TERM)
	err := s.validator.Var(term, "required")
	if err != nil {
		log.Error(err, "BadRequest", "teachers HandleGetTeacher")
		responser.BadRequset(w, r, "Query parameter `term` must be provided")

		return
	}

	response, err := s.client.GetTeacher(&ruzfaclient.GetEntitiesInput{
		Term: term,
	})
	if err != nil {
		log.Error(err, "Internal", "teachers HandleGetTeacher")
		responser.Internal(w, r, err.Error())

		return
	}

	teachers := make([]dto.Teacher, 0)
	for _, teacher := range response {
		if teacher.Id != "" {
			teachers = append(teachers, dto.Teacher{
				Id:          teacher.Id,
				Title:       teacher.Label,
				Description: teacher.Description,
			})
		}
	}

	responser.Success(w, r, teachers)
}
