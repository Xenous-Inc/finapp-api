package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/swaggo/http-swagger"
	_"github.com/Xenous-Inc/finapp-api/docs"
)

type Router struct {
	Client *ruzfaclient.Client
	ClientOrgfaclient *orgfaclient.Client
}

func NewRouter(client *ruzfaclient.Client, clientOrgfaclient *orgfaclient.Client) *Router {
	return &Router{
		Client: client,
		ClientOrgfaclient: clientOrgfaclient,
	}
}

func (s *Router) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5555/swagger/doc.json"), //The url pointing to API definition
	))

	r.Get("/", s.pingHandler)

	r.Post("/finapp/api/group/", s.HandlerGetGroup)
	r.Post("/finapp/api/group/schedule", s.HandlerGetGroupSchedule)

	r.Post("/finapp/api/teacher", s.HandlerGetTeacher)
	r.Post("/finapp/api/teacher/schedule", s.HandlerGetTeacherSchedule)

	r.Post("/finapp/api/auditorium", s.HandlerGetAuditorium)
	r.Post("/finapp/api/auditorium/schedule", s.HandlerGetAuditoriumSchedule)

	return r
}

// @Summary GetGroup
// @Tags TimeTable
// @Description get group
// @Accept json
// @Produce json
// @Param input body ruzfaclient.GetGroupsInput true "group info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/group [post]
func (s *Router) HandlerGetGroup(w http.ResponseWriter, r *http.Request) {
	var input *ruzfaclient.GetGroupsInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	groups, err := s.Client.GetGroups(input)

	res, err := json.Marshal(groups)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
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
// @Router /finapp/api/group/schedule [post]
func (s *Router) HandlerGetGroupSchedule(w http.ResponseWriter, r *http.Request) {
	var input *ruzfaclient.GetGroupScheduleInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	groupsSchedule, err := s.Client.GetGroupSchedule(input)

	res, err := json.Marshal(groupsSchedule)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
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
// @Router /finapp/api/teacher [post]
func (s *Router) HandlerGetTeacher(w http.ResponseWriter, r *http.Request) {
	var input *ruzfaclient.GetTeacherInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	teacher, err := s.Client.GetTeacher(input)

	res, err := json.Marshal(teacher)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
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
// @Router /finapp/api/teacher/schedule [post]
func (s *Router) HandlerGetTeacherSchedule(w http.ResponseWriter, r *http.Request) {
	var input *ruzfaclient.GetTeacherScheduleInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	groupsSchedule, err := s.Client.GetTeacherSchedule(input)

	res, err := json.Marshal(groupsSchedule)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
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
// @Router /finapp/api/auditorium [post]
func (s *Router) HandlerGetAuditorium(w http.ResponseWriter, r *http.Request) {
	var input *ruzfaclient.GetAuditoriumInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	auditorium, err := s.Client.GetAuditorium(input)

	res, err := json.Marshal(auditorium)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
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
// @Router /finapp/api/auditorium/schedule [post]
func (s *Router) HandlerGetAuditoriumSchedule(w http.ResponseWriter, r *http.Request) {
	var input *ruzfaclient.GetAuditoriumScheduleInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	auditoriumSchedule, err := s.Client.GetAuditoriumSchedule(input)

	res, err := json.Marshal(auditoriumSchedule)

	if err != nil {
		fmt.Fprintf(w, "Unlucky:  %s", err)
	}

	fmt.Fprintf(w, string(res))
}

//AUTH

func (s *Router) HandlerAuth(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	auth, err := s.ClientOrgfaclient.Login(input)













	r.ParseForm()
	authForm := r.Form.Get("AUTH_FORM")
	typ := r.Form.Get("TYPE")
	userLogin := r.Form.Get("USER_LOGIN")
	userPassword := r.Form.Get("USER_PASSWORD")

	cookie, err := r.Cookie("PHPSESSID")



	auth, err := s.ClientOrgfaclient.Login(&orgfaclient.LoginInput{
		AuthForm:     authForm,
		Typ:     typ,
		Login: userLogin,
		Password: userPassword,
		Cookie: *cookie,
	})
}
