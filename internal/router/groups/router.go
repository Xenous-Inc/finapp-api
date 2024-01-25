package groups

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
	client *ruzfaclient.Client

	validator *validator.Validate
}

func NewRouter(client *ruzfaclient.Client) *Router {
	return &Router{
		client:    client,
		validator: validator.New(),
	}
}

func (s *Router) Route(r chi.Router) {
	r.Get("/", s.HandleGetGroup)
}

// @Summary Return List of groups
// @Description Return list of student groups which found by provided query term
// @Tags groups
// @Param term query string true "Group search mask"
// @Produce json
// @Success 200 {object} []dto.Group
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /groups/ [get]
func (s *Router) HandleGetGroup(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get(constants.QUERY_TERM)
	err := s.validator.Var(term, "required")
	if err != nil {
		log.Error(err, "BadRequest", "groups HandleGetGroup")
		responser.BadRequset(w, r, "Query parameter `term` must be provided")

		return
	}

	response, err := s.client.GetGroups(&ruzfaclient.GetEntitiesInput{
		Term: term,
	})
	if err != nil {
		log.Error(err, "Internal", "groups HandleGetGroup")
		responser.Internal(w, r, err.Error())

		return
	}

	groups := make([]dto.Classroom, 0)
	for _, group := range response {
		if group.Id != "" {
			groups = append(groups, dto.Classroom{
				Id:          group.Id,
				Title:       group.Label,
				Description: group.Description,
			})
		}
	}

	responser.Success(w, r, groups)
}
