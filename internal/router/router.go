package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient/dto"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/router/auth"

	_ "github.com/Xenous-Inc/finapp-api/docs"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type RootRouter struct {
	Client            *ruzfaclient.Client
	ClientOrgfaclient *orgfaclient.Client
	authRouter *auth.Router 
}

func NewRootRouter(client *ruzfaclient.Client, clientOrgfaclient *orgfaclient.Client) *RootRouter {
	return &RootRouter{
		Client:            client,
		ClientOrgfaclient: clientOrgfaclient,
		authRouter: auth.NewRouter(clientOrgfaclient),
	}
}

func (s *RootRouter) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/auth", s.authRouter.Route)
	
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
	r.Post("/finapp/api/auth/recordbook", s.HandlerGetRecordBook)
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
func (s *RootRouter) HandlerGetGroup(w http.ResponseWriter, r *http.Request) {
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
func (s *RootRouter) HandlerGetGroupSchedule(w http.ResponseWriter, r *http.Request) {
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
func (s *RootRouter) HandlerGetTeacher(w http.ResponseWriter, r *http.Request) {
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
func (s *RootRouter) HandlerGetTeacherSchedule(w http.ResponseWriter, r *http.Request) {
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
func (s *RootRouter) HandlerGetAuditorium(w http.ResponseWriter, r *http.Request) {
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
func (s *RootRouter) HandlerGetAuditoriumSchedule(w http.ResponseWriter, r *http.Request) {
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

// @Summary Auth
// @Tags OrgFaRu
// @Description auth
// @Accept x-www-form-urlencoded
// @Produce x-www-form-urlencoded
// @Param input body orgfaclient.LoginInput true "auth"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth [post]
func (s *RootRouter) HandlerAuth(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userLogin := r.URL.Query().Get(dto.USER_LOGIN)
	userPassword := r.URL.Query().Get(dto.USER_PASSWORD)

	sessionId, err := s.ClientOrgfaclient.Login(&orgfaclient.LoginInput{
		Login:    userLogin,
		Password: userPassword,
	})

	if err != nil {
		fmt.Fprintf(w, "auth:  %s", clients.ErrRequest)
	}

	res, err := json.Marshal(sessionId)

	if err != nil {
		fmt.Fprintf(w, "auth marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetMyGroup
// @Tags OrgFaRu
// @Description get myGroup
// @Accept json
// @Produce json
// @Param input body orgfaclient.GetMyGroupInput true "myGroup info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/mygroup [get]
func (s *RootRouter) HandlerGetMyGroup(w http.ResponseWriter, r *http.Request) {
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

// @Summary GetRecordBook
// @Tags OrgFaRu
// @Description get GetRecordBook
// @Accept json
// @Produce json
// @Param input body orgfaclient.GetRecordBookInput true "Record book info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/recordbook [get]
func (s *RootRouter) HandlerGetRecordBook(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.AuthSession
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.ClientOrgfaclient.GetRecordBook(input)

	if err != nil {
		fmt.Fprintf(w, "Get record book:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get record book marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}

// @Summary GetMiniProfile
// @Tags OrgFaRu
// @Description GetMiniProfile
// @Accept json
// @Produce json
// @Param input body orgfaclient.GetMiniProfileInput true "Get mini profile info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/miniprofile [get]
func (s *RootRouter) HandlerGetMiniProfile(w http.ResponseWriter, r *http.Request) {
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

// @Summary GetProfile
// @Tags OrgFaRu
// @Description GetProfile
// @Accept json
// @Produce json
// @Param input body orgfaclient.GetProfileInput true "Profile info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/miniprofile/profile [get]
func (s *RootRouter) HandlerGetProfile(w http.ResponseWriter, r *http.Request) {
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

// @Summary GetOrder
// @Tags OrgFaRu
// @Description GetOrder
// @Accept json
// @Produce json
// @Param input body orgfaclient.GetOrderInput true "Order info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/miniprofile/profile/order [get]
func (s *RootRouter) HandlerGetOrder(w http.ResponseWriter, r *http.Request) {
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

// @Summary GetStudentCard
// @Tags OrgFaRu
// @Description GetStudentCard
// @Accept json
// @Produce json
// @Param input body orgfaclient.GetStudentCardInput true "Get student card info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/miniprofile/profile/studentcard [get]
func (s *RootRouter) HandlerGetStudentCard(w http.ResponseWriter, r *http.Request) {
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

// @Summary GetStudyPlan
// @Tags OrgFaRu
// @Description GetStudyPlan
// @Accept json
// @Produce json
// @Param input body orgfaclient.GetStudyPlanInput true "Get study plan info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/miniprofile/profile/studyplan/{profileId} [get]
func (s *RootRouter) HandlerGetStudyPlan(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetStudyPlanInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.ClientOrgfaclient.GetStudyPlan(input)

	if err != nil {
		fmt.Fprintf(w, "Get study plan:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get study plan marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}
