package schedule

import (
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/dto"
	"github.com/Xenous-Inc/finapp-api/internal/router/constants"
	"github.com/Xenous-Inc/finapp-api/internal/router/utils/responser"
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
	r.Post("/group", s.HandleGetGroupSchedule)
	r.Post("/teacher", s.HandleGetTeacherSchedule)
	r.Post("/auditorium", s.HandleGetTeacherSchedule)
}

// @Summary Return schedule for provided group
// @Description Returns schedule for provided group Id and time interval
// @Tags schedule
// @Param id query string true "Group ID"
// @Param from query string true "Group ID"
// @Param to query string true "Group ID"
// @Produce json
// @Success 200 {object} []dto.Group
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /groups/ [get]
func (s *Router) HandleGetGroupSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: dto.Date(r.URL.Query().Get(constants.QUERY_START_DATE)),
		EndDate:   dto.Date(r.URL.Query().Get(constants.QUERY_END_DATE)),
	}

	err := s.validator.Struct(input)

	if err != nil {
		responser.BadRequset(w, r, "Password and login must be provided")

		return
	}

	groupsSchedule, err := s.client.GetGroupSchedule(&ruzfaclient.GetGroupScheduleInput{
		GroupId:   input.EntityId,
		StartDate: input.StartDate.String(),
		EndDate:   input.EndDate.String(),
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			responser.BadRequset(w, r, "Error validation")
		default:
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, groupsSchedule)
}

func (s *Router) HandleGetTeacherSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: dto.Date(r.URL.Query().Get(constants.QUERY_START_DATE)),
		EndDate:   dto.Date(r.URL.Query().Get(constants.QUERY_END_DATE)),
	}

	err := s.validator.Struct(input)

	if err != nil {
		responser.BadRequset(w, r, "Password and login must be provided")

		return
	}

	teacherSchedule, err := s.client.GetTeacherSchedule(&ruzfaclient.GetTeacherScheduleInput{
		Id:        input.EntityId,
		StartDate: input.StartDate.String(),
		EndDate:   input.EndDate.String(),
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			responser.BadRequset(w, r, "Error validation")
		default:
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, teacherSchedule)
}

func (s *Router) HandleGetAuditoriumSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: dto.Date(r.URL.Query().Get(constants.QUERY_START_DATE)),
		EndDate:   dto.Date(r.URL.Query().Get(constants.QUERY_END_DATE)),
	}

	err := s.validator.Struct(input)

	if err != nil {
		responser.BadRequset(w, r, "Password and login must be provided")

		return
	}

	auditoriumSchedule, err := s.client.GetGroupSchedule(&ruzfaclient.GetGroupScheduleInput{
		GroupId:   input.EntityId,
		StartDate: input.StartDate.String(),
		EndDate:   input.EndDate.String(),
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			responser.BadRequset(w, r, "Error validation")
		default:
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, auditoriumSchedule)
}
