package schedule

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
		client: client,
	}
}

func (s *Router) Route(r chi.Router) {
	r.Post("/group", s.HandleGetGroupSchedule)
	// r.Post("/teacher", s.HandleGetTeacherSchedule)
	// r.Post("/auditorium", s.HandleGetTeacherSchedule)
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

	groupsSchedule, err := s.client.GetGroupSchedule(&ruzfaclient.GetGroupScheduleInput{
		GroupId:   input.EntityId,
		StartDate: input.StartDate.String(),
		EndDate:   input.EndDate.String(),
	})

	res, err := json.Marshal(groupsSchedule)
	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
	}

	fmt.Fprintf(w, string(res))
}

/*func (s *Router) HandleGetTeacherSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: r.URL.Query().Get(constants.QUERY_START_DATE),
		EndDate:   r.URL.Query().Get(constants.QUERY_END_DATE),
	}

	teacherSchedule, err := s.client.GetTeacherSchedule(&ruzfaclient.GetTeacherScheduleInput{
		Id:        input.EntityId,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	})

	res, err := json.Marshal(teacherSchedule)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
	}

	fmt.Fprintf(w, string(res))
}

func (s *Router) HandleGetAuditoriumSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: r.URL.Query().Get(constants.QUERY_START_DATE),
		EndDate:   r.URL.Query().Get(constants.QUERY_END_DATE),
	}

	auditoriumSchedule, err := s.client.GetGroupSchedule(&ruzfaclient.GetGroupScheduleInput{
		GroupId:   input.EntityId,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	})

	if err != nil {
		//TODO: add error handling
	}

	res, err := json.Marshal(auditoriumSchedule)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
	}

	fmt.Fprintf(w, string(res))
}*/