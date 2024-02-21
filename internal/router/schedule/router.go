package schedule

import (
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient/models"
	"github.com/Xenous-Inc/finapp-api/internal/dto"
	"github.com/Xenous-Inc/finapp-api/internal/router/constants"
	"github.com/Xenous-Inc/finapp-api/internal/router/utils/responser"
	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
	"github.com/go-chi/chi"
	"golang.org/x/sync/errgroup"
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
	r.Route("/group", func(r chi.Router) {
		r.Get("/", s.HandleGetGroupSchedule)
		r.Get("/mini", s.HandleGetGroupMiniSchedule)
	})
	r.Route("/teacher", func(r chi.Router) {
		r.Get("/", s.HandleGetTeacherSchedule)
		r.Get("/mini", s.HandleGetGroupMiniSchedule)
	})
	r.Route("/classroom", func(r chi.Router) {
		r.Get("/", s.HandleGetTeacherSchedule)
		r.Get("/mini", s.HandleGetClassroomMiniSchedule)
	})
}

// @Summary Return schedule for provided group
// @Description Returns schedule for provided group Id and time interval
// @Tags schedule
// @Param id query string true "Group ID"
// @Param from query string true "Start date"
// @Param to query string true "End date"
// @Produce json
// @Success 200 {object} []dto.ScheduleItem
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /schedule/group [get]
func (s *Router) HandleGetGroupSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: dto.Date(r.URL.Query().Get(constants.QUERY_START_DATE)),
		EndDate:   dto.Date(r.URL.Query().Get(constants.QUERY_END_DATE)),
	}

	err := s.validator.Struct(input)
	if err != nil {
		log.Error(err, "BadRequest", "schedule HandleGetGroupSchedule")
		responser.BadRequset(w, r, "All required parameters must be provided")

		return
	}

	eg := errgroup.Group{}
	eg.Go(input.StartDate.Validate)
	eg.Go(input.EndDate.Validate)
	if err = eg.Wait(); err != nil {
		responser.BadRequset(w, r, err.Error())

		return
	}

	groupsSchedule, err := s.client.GetGroupSchedule(&ruzfaclient.GetScheduleInput{
		Id:        input.EntityId,
		StartDate: string(input.StartDate),
		EndDate:   string(input.EndDate),
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "schedule HandleGetGroupSchedule")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "schedule HandleGetGroupSchedule")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "schedule HandleGetGroupSchedule")
			responser.BadRequset(w, r, "Error validation")
		default:
			log.Error(err, "Internal", "schedule HandleGetGroupSchedule")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	items := make([]dto.ScheduleItem, 0)

	for _, schedule := range groupsSchedule {
		items = append(items, dto.ScheduleItemFromClientModel(&schedule))
	}

	responser.Success(w, r, items)
}

// @Summary Return schedule for provided teacher
// @Description Returns schedule for provided teacher Id and time interval
// @Tags schedule
// @Param id query string true "Teacher ID"
// @Param from query string true "Start date"
// @Param to query string true "End date"
// @Produce json
// @Success 200 {object} []dto.ScheduleItem
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /schedule/teacher [get]
func (s *Router) HandleGetTeacherSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: dto.Date(r.URL.Query().Get(constants.QUERY_START_DATE)),
		EndDate:   dto.Date(r.URL.Query().Get(constants.QUERY_END_DATE)),
	}

	err := s.validator.Struct(input)
	if err != nil {
		log.Error(err, "BadRequest", "schedule HandleGetTeacherSchedule")
		responser.BadRequset(w, r, "All required parameters must be provided")

		return
	}
	if err != nil {
		log.Error(err, "BadRequest", "schedule HandleGetTeacherSchedule")
		responser.BadRequset(w, r, "All required parameters must be provided")

		return
	}

	eg := errgroup.Group{}
	eg.Go(input.StartDate.Validate)
	eg.Go(input.EndDate.Validate)
	if err = eg.Wait(); err != nil {
		responser.BadRequset(w, r, err.Error())

		return
	}

	teacherSchedule, err := s.client.GetTeacherSchedule(&ruzfaclient.GetScheduleInput{
		Id:        input.EntityId,
		StartDate: string(input.StartDate),
		EndDate:   string(input.EndDate),
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "schedule HandleGetTeacherSchedule")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "schedule HandleGetTeacherSchedule")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "schedule HandleGetTeacherSchedule")
			responser.BadRequset(w, r, "Error validation")
		default:
			log.Error(err, "Internal", "schedule HandleGetTeacherSchedule")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	items := make([]dto.ScheduleItem, 0)

	for _, schedule := range teacherSchedule {
		items = append(items, dto.ScheduleItemFromClientModel(&schedule))
	}

	responser.Success(w, r, items)
}

// @Summary Return schedule for provided classroom
// @Description Returns schedule for provided classroom Id and time interval
// @Tags schedule
// @Param id query string true "Classroom ID"
// @Param from query string true "Start date"
// @Param to query string true "End date"
// @Produce json
// @Success 200 {object} []dto.ScheduleItem
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /schedule/classroom [get]
func (s *Router) HandleGetAuditoriumSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: dto.Date(r.URL.Query().Get(constants.QUERY_START_DATE)),
		EndDate:   dto.Date(r.URL.Query().Get(constants.QUERY_END_DATE)),
	}

	err := s.validator.Struct(input)
	if err != nil {
		log.Error(err, "BadRequest", "schedule HandleGetAuditoriumSchedule")
		responser.BadRequset(w, r, "All required parameters must be provided")

		return
	}
	if err != nil {
		log.Error(err, "BadRequest", "schedule HandleGetAuditoriumSchedule")
		responser.BadRequset(w, r, "All required parameters must be provided")

		return
	}

	eg := errgroup.Group{}
	eg.Go(input.StartDate.Validate)
	eg.Go(input.EndDate.Validate)
	if err = eg.Wait(); err != nil {
		responser.BadRequset(w, r, err.Error())

		return
	}

	auditoriumSchedule, err := s.client.GetGroupSchedule(&ruzfaclient.GetScheduleInput{
		Id:        input.EntityId,
		StartDate: string(input.StartDate),
		EndDate:   string(input.EndDate),
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "schedule HandleGetAuditoriumSchedule")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "schedule HandleGetAuditoriumSchedule")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "schedule HandleGetAuditoriumSchedule")
			responser.BadRequset(w, r, "Error validation")
		default:
			log.Error(err, "Internal", "schedule HandleGetAuditoriumSchedule")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	items := make([]dto.ScheduleItem, 0)

	for _, schedule := range auditoriumSchedule {
		items = append(items, dto.ScheduleItemFromClientModel(&schedule))
	}

	responser.Success(w, r, items)
}

// @Summary Return mini schedule for provided group
// @Description Returns schedule in compact format for provided teacher Id and time interval
// @Tags schedule
// @Param id query string true "Teacher ID"
// @Param from query string true "Start date"
// @Param to query string true "End date"
// @Produce json
// @Success 200 {object} []dto.MiniScheduleItem
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /schedule/group/mini [get]
func (s *Router) HandleGetGroupMiniSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: dto.Date(r.URL.Query().Get(constants.QUERY_START_DATE)),
		EndDate:   dto.Date(r.URL.Query().Get(constants.QUERY_END_DATE)),
	}

	err := s.validator.Struct(input)
	if err != nil {
		responser.BadRequset(w, r, "All required parameters must be provided")

		return
	}

	eg := errgroup.Group{}
	eg.Go(input.StartDate.Validate)
	eg.Go(input.EndDate.Validate)
	if err = eg.Wait(); err != nil {
		responser.BadRequset(w, r, err.Error())

		return
	}

	groupsSchedule, err := s.client.GetGroupSchedule(&ruzfaclient.GetScheduleInput{
		Id:        input.EntityId,
		StartDate: string(input.StartDate),
		EndDate:   string(input.EndDate),
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			responser.Internal(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			responser.Internal(w, r, "Invalid entity")
		case clients.ErrValidation:
			responser.BadRequset(w, r, "Error validation")
		default:
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, s.filterMiniSchedule(groupsSchedule))
}

// @Summary Return mini schedule for provided classroom
// @Description Returns schedule in compact format for provided classroom Id and time interval
// @Tags schedule
// @Param id query string true "Classroom ID"
// @Param from query string true "Start date"
// @Param to query string true "End date"
// @Produce json
// @Success 200 {object} []dto.MiniScheduleItem
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /schedule/classroom/mini [get]
func (s *Router) HandleGetClassroomMiniSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: dto.Date(r.URL.Query().Get(constants.QUERY_START_DATE)),
		EndDate:   dto.Date(r.URL.Query().Get(constants.QUERY_END_DATE)),
	}

	err := s.validator.Struct(input)
	if err != nil {
		responser.BadRequset(w, r, "All required parameters must be provided")

		return
	}

	eg := errgroup.Group{}
	eg.Go(input.StartDate.Validate)
	eg.Go(input.EndDate.Validate)
	if err = eg.Wait(); err != nil {
		responser.BadRequset(w, r, err.Error())

		return
	}

	schedule, err := s.client.GetAuditoriumSchedule(&ruzfaclient.GetScheduleInput{
		Id:        input.EntityId,
		StartDate: string(input.StartDate),
		EndDate:   string(input.EndDate),
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			responser.Internal(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			responser.Internal(w, r, "Invalid entity")
		case clients.ErrValidation:
			responser.BadRequset(w, r, "Error validation")
		default:
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, s.filterMiniSchedule(schedule))
}

// @Summary Return mini schedule for provided teacher
// @Description Returns schedule in compact format for provided teacher Id and time interval
// @Tags schedule
// @Param id query string true "Teacher ID"
// @Param from query string true "Start date"
// @Param to query string true "End date"
// @Produce json
// @Success 200 {object} []dto.MiniScheduleItem
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /schedule/teacher/mini [get]
func (s *Router) HandleGetTeacherMiniSchedule(w http.ResponseWriter, r *http.Request) {
	input := &dto.GetScheduleRequest{
		EntityId:  r.URL.Query().Get(constants.QUERY_ID),
		StartDate: dto.Date(r.URL.Query().Get(constants.QUERY_START_DATE)),
		EndDate:   dto.Date(r.URL.Query().Get(constants.QUERY_END_DATE)),
	}

	err := s.validator.Struct(input)
	if err != nil {
		responser.BadRequset(w, r, "All required parameters must be provided")

		return
	}

	eg := errgroup.Group{}
	eg.Go(input.StartDate.Validate)
	eg.Go(input.EndDate.Validate)
	if err = eg.Wait(); err != nil {
		responser.BadRequset(w, r, err.Error())

		return
	}

	schedule, err := s.client.GetTeacherSchedule(&ruzfaclient.GetScheduleInput{
		Id:        input.EntityId,
		StartDate: string(input.StartDate),
		EndDate:   string(input.EndDate),
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			responser.Internal(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			responser.Internal(w, r, "Invalid entity")
		case clients.ErrValidation:
			responser.BadRequset(w, r, "Error validation")
		default:
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, s.filterMiniSchedule(schedule))
}

func (r *Router) filterMiniSchedule(scheduleResponse []models.Schedule) []dto.MiniScheduleItem {

	items := make([]dto.MiniScheduleItem, 0)
	for i := 1; i < len(scheduleResponse); i++ {
		if scheduleResponse[i].Discipline != scheduleResponse[i-1].Discipline && scheduleResponse[i].LessonNumberStart != scheduleResponse[i-1].LessonNumberStart {
			items = append(items, dto.MiniScheduleItemFromClientModel(&scheduleResponse[i-1]))
		} else {
			items = append(items, dto.MiniScheduleItemFromClientModel(&scheduleResponse[i]))
		}
	}

	return items
}
