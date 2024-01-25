package schedule

import (
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/dto"
	"github.com/Xenous-Inc/finapp-api/internal/router/constants"
	"github.com/Xenous-Inc/finapp-api/internal/router/utils/responser"
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
	r.Get("/group", s.HandleGetGroupSchedule)
	r.Get("/teacher", s.HandleGetTeacherSchedule)
	r.Get("/classroom", s.HandleGetTeacherSchedule)
}

// @Summary Return schedule for provided group
// @Description Returns schedule for provided group Id and time interval
// @Tags schedule
// @Param id query string true "Group ID"
// @Param from query string true "Group ID"
// @Param to query string true "Group ID"
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

	groupsSchedule, err := s.client.GetGroupSchedule(&ruzfaclient.GetGroupScheduleInput{
		GroupId:   input.EntityId,
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

	items := make([]dto.ScheduleItem, 0)

	for _, schedule := range groupsSchedule {
		items = append(items, dto.ScheduleItem{
			ClassroomNumber:   schedule.Auditorium,
			StartsAt:          schedule.BeginLesson,
			EndsAt:            schedule.EndLesson,
			Address:           schedule.Building,
			Lesson:            schedule.Discipline,
			LessonType:        schedule.KindOfWork,
			LessonNumberStart: schedule.LessonNumberStart,
			LessonNumberEnd:   schedule.LessonNumberEnd,
			Date:              schedule.Date,
			WeekDay:           schedule.DayOfWeek,
			Lecturer:          schedule.Lecturer,
		})
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
		responser.BadRequset(w, r, "All required parameters must be provided")

		return
	}
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

	teacherSchedule, err := s.client.GetTeacherSchedule(&ruzfaclient.GetTeacherScheduleInput{
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

	items := make([]dto.ScheduleItem, 0)

	for _, schedule := range teacherSchedule {
		items = append(items, dto.ScheduleItem{
			ClassroomNumber:   schedule.Auditorium,
			StartsAt:          schedule.BeginLesson,
			EndsAt:            schedule.EndLesson,
			Address:           schedule.Building,
			Lesson:            schedule.Discipline,
			LessonType:        schedule.KindOfWork,
			LessonNumberStart: schedule.LessonNumberStart,
			LessonNumberEnd:   schedule.LessonNumberEnd,
			Date:              schedule.Date,
			WeekDay:           schedule.DayOfWeek,
			Lecturer:          schedule.Lecturer,
		})
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
		responser.BadRequset(w, r, "All required parameters must be provided")

		return
	}
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

	auditoriumSchedule, err := s.client.GetGroupSchedule(&ruzfaclient.GetGroupScheduleInput{
		GroupId:   input.EntityId,
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

	items := make([]dto.ScheduleItem, 0)

	for _, schedule := range auditoriumSchedule {
		items = append(items, dto.ScheduleItem{
			ClassroomNumber:   schedule.Auditorium,
			StartsAt:          schedule.BeginLesson,
			EndsAt:            schedule.EndLesson,
			Address:           schedule.Building,
			Lesson:            schedule.Discipline,
			LessonType:        schedule.KindOfWork,
			LessonNumberStart: schedule.LessonNumberStart,
			LessonNumberEnd:   schedule.LessonNumberEnd,
			Date:              schedule.Date,
			WeekDay:           schedule.DayOfWeek,
			Lecturer:          schedule.Lecturer,
		})
	}

	responser.Success(w, r, items)
}
