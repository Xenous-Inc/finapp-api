package router

import (
	"encoding/json"
	"fmt"

	"net/http"

	_ "github.com/Xenous-Inc/finapp-api/docs"
	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	Client *ruzfaclient.Client
}

func NewRouter(client *ruzfaclient.Client) *Router {
	return &Router{
		Client: client,
	}
}

func (s *Router) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5555/swagger/doc.json"), //The url pointing to API definition
	))

	r.Get("/", s.pingHandler)

	r.Get("/finapp/api/group", s.HandlerGetGroup)
	r.Get("/finapp/api/group/schedule/{groupid}", s.HandlerGetGroupSchedule)

	r.Get("/finapp/api/teacher", s.HandlerGetTeacher)
	r.Get("/finapp/api/teacher/schedule/{teacherid}", s.HandlerGetTeacherSchedule)

	r.Get("/finapp/api/auditorium", s.HandlerGetAuditorium)
	r.Get("/finapp/api/auditorium/schedule/{auditoriumid}", s.HandlerGetAuditoriumSchedule)

	return r
}

// @Summary GetGroup
// @Tags TimeTable
// @Description get group
// @Accept json
// @Produce json
// @Param input body ruzfaclient.GetGroupsInput true "group info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/group [get]
func (s *Router) HandlerGetGroup(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")

	groups, err := s.Client.GetGroups(&ruzfaclient.GetGroupsInput{
		GroupTerm: term,
	})

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get groups, %s", clients.ErrRequest),
		})
	}

	res, err := json.Marshal(groups)

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get groups marshal, %s", clients.ErrRequest),
		})
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetGroupSchedule
// @Tags TimeTable
// @Description get groupSchedule
// @Accept json
// @Produce json
// @Param input body ruzfaclient.GetGroupScheduleInput true "group schedule info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/group/schedule/{groupid} [get]
func (s *Router) HandlerGetGroupSchedule(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start")
	endDate := r.URL.Query().Get("finish")
	url := chi.URLParam(r, "groupid")
	groupsSchedule, err := s.Client.GetGroupSchedule(&ruzfaclient.GetGroupScheduleInput{
		GroupId:   url,
		StartDate: startDate,
		EndDate:   endDate,
	})

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get groups schedule, %s", clients.ErrRequest),
		})
	}

	res, err := json.Marshal(groupsSchedule)

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get groups schedule marshal, %s", clients.ErrRequest),
		})
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetTeacher
// @Tags TimeTable
// @Description get teacher
// @Accept json
// @Produce json
// @Param input body ruzfaclient.GetTeacherInput true "teacher info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/teacher [get]
func (s *Router) HandlerGetTeacher(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	teacher, err := s.Client.GetTeacher(&ruzfaclient.GetTeacherInput{
		TeacherTerm: term,
	})

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get teacher, %s", clients.ErrRequest),
		})
	}

	res, err := json.Marshal(teacher)

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get teacher marshal, %s", clients.ErrRequest),
		})
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetTeacherSchedule
// @Tags TimeTable
// @Description get teacherSchedule
// @Accept json
// @Produce json
// @Param input body ruzfaclient.GetTeacherScheduleInput true "teacher schedule info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/teacher/schedule [get]
func (s *Router) HandlerGetTeacherSchedule(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start")
	endDate := r.URL.Query().Get("finish")
	url := chi.URLParam(r, "teacherid")
	teacherSchedule, err := s.Client.GetTeacherSchedule(&ruzfaclient.GetTeacherScheduleInput{
		Id:        url,
		StartDate: startDate,
		EndDate:   endDate,
	})

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get teacher schedule, %s", clients.ErrRequest),
		})
	}

	res, err := json.Marshal(teacherSchedule)

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get teacher schedule marshal, %s", clients.ErrRequest),
		})
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetAuditorium
// @Tags TimeTable
// @Description get auditorium
// @Accept json
// @Produce json
// @Param input body ruzfaclient.GetAuditoriumInput true "auditorium info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auditorium [get]
func (s *Router) HandlerGetAuditorium(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	auditorium, err := s.Client.GetAuditorium(&ruzfaclient.GetAuditoriumInput{
		AuditoriumTerm: term,
	})

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get auditorium, %s", clients.ErrRequest),
		})
	}

	res, err := json.Marshal(auditorium)

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get auditorium marshal, %s", clients.ErrRequest),
		})
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetAuditoriumSchedule
// @Tags TimeTable
// @Description get auditorium schedule
// @Accept json
// @Produce json
// @Param input body ruzfaclient.GetAuditoriumScheduleInput true "auditorium schedule info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auditorium/schedule [get]
func (s *Router) HandlerGetAuditoriumSchedule(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start")
	endDate := r.URL.Query().Get("finish")
	url := chi.URLParam(r, "auditoriumid")
	auditoriumSchedule, err := s.Client.GetAuditoriumSchedule(&ruzfaclient.GetAuditoriumScheduleInput{
		Id:        url,
		StartDate: startDate,
		EndDate:   endDate,
	})

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get auditorium schedule, %s", clients.ErrRequest),
		})
	}

	res, err := json.Marshal(auditoriumSchedule)

	if err != nil {
		fmt.Println(&clients.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Get auditorium schedule marshal, %s", clients.ErrRequest),
		})
	}

	fmt.Fprintf(w, string(res))
}
