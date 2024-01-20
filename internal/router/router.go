package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"

	_ "github.com/Xenous-Inc/finapp-api/docs"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/swaggo/http-swagger"
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

	r.Post("/finapp/api/auth", s.HandlerAuth)
	r.Post("/finapp/api/auth/mygroup", s.HandlerGetMyGroup)
	r.Post("/finapp/api/auth/zachetka", s.HandlerGetZachetka)
	r.Post("/finapp/api/auth/miniprofile", s.HandlerGetMiniProfile)
	r.Post("/finapp/api/auth/miniprofile/profile", s.HandlerGetProfile)
	r.Post("/finapp/api/auth/miniprofile/profile/order", s.HandlerGetOrder)
	r.Post("/finapp/api/auth/miniprofile/profile/studentcard", s.HandlerGetStudentCard)
	r.Post("/finapp/api/auth/miniprofile/profile/studyplan/{profileId}", s.HandlerGetStudyPlan)
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

// @Summary GetGroup
// @Tags OrgFaRu
// @Description auth
// @Accept json
// @Produce json
// @Param input body orgfaclient.LoginInput true "auth"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth [post]
func (s *Router) HandlerAuth(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userLogin := r.URL.Query().Get("USER_LOGIN")
	userPassword := r.URL.Query().Get("USER_PASSWORD")

	sessionId, err := s.ClientOrgfaclient.Login(&orgfaclient.LoginInput{
		Login: userLogin,
		Password: userPassword,
	})

	if err != nil {
		fmt.Fprintf(w, "Get my group:  %s", clients.ErrRequest)
	}

	res, err := json.Marshal(sessionId)
	//res2, err := json.Marshal(profileId)

	if err != nil {
		fmt.Fprintf(w, "Get my group:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetMyGroup
// @Tags OrgFaRu
// @Description get myGroup
// @Accept json
// @Produce json
// @Param input body orgfaclient.AuthSession true "myGroup info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/mygroup [get]
func (s *Router) HandlerGetMyGroup(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetMyGroupInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.ClientOrgfaclient.GetMyGroup(input)

	if err != nil {
		fmt.Fprintf(w, "Get my group:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get my group marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetMyGroup
// @Tags OrgFaRu
// @Description get myGroup
// @Accept json
// @Produce json
// @Param input body orgfaclient.AuthSession true "myGroup info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/mygroup [get]
func (s *Router) HandlerGetZachetka(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.AuthSession
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.ClientOrgfaclient.GetRecordBook(input)

	if err != nil {
		fmt.Fprintf(w, "Get zachetka:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get zachetka marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetMyGroup
// @Tags OrgFaRu
// @Description get myGroup
// @Accept json
// @Produce json
// @Param input body orgfaclient.AuthSession true "myGroup info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/mygroup [get]
func (s *Router) HandlerGetMiniProfile(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetMiniProfileInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.ClientOrgfaclient.GetMiniProfile(input)

	if err != nil {
		fmt.Fprintf(w, "Get mini profile:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get mini profile marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetMyGroup
// @Tags OrgFaRu
// @Description get myGroup
// @Accept json
// @Produce json
// @Param input body orgfaclient.AuthSession true "myGroup info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/mygroup [get]
func (s *Router) HandlerGetProfile(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetProfileInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.ClientOrgfaclient.GetProfile(input)

	if err != nil {
		fmt.Fprintf(w, "Get profile:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get profile marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetMyGroup
// @Tags OrgFaRu
// @Description get myGroup
// @Accept json
// @Produce json
// @Param input body orgfaclient.AuthSession true "myGroup info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/mygroup [get]
func (s *Router) HandlerGetOrder(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetOrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.ClientOrgfaclient.GetOrder(input)

	if err != nil {
		fmt.Fprintf(w, "Get order:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get order marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetMyGroup
// @Tags OrgFaRu
// @Description get myGroup
// @Accept json
// @Produce json
// @Param input body orgfaclient.AuthSession true "myGroup info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/mygroup [get]
func (s *Router) HandlerGetStudentCard(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "profileId")
	myGroup, err := s.ClientOrgfaclient.GetStudentCard(&orgfaclient.GetStudentCardInput{
		ProfileId: url,
	})

	if err != nil {
		fmt.Fprintf(w, "Get student card:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get student card marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetMyGroup
// @Tags OrgFaRu
// @Description get myGroup
// @Accept json
// @Produce json
// @Param input body orgfaclient.AuthSession true "myGroup info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/mygroup [get]
func (s *Router) HandlerGetStudyPlan(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetStudyPlanInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.ClientOrgfaclient.GetStudyPlan(input)

	if err != nil {
		fmt.Fprintf(w, "Get my study plan:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get my study plan marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}