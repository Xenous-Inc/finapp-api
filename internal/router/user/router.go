package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/go-chi/chi"
)

type Router struct {
	client *orgfaclient.Client
}

func NewRouter(client *orgfaclient.Client) *Router {
	return &Router{
		client:            client,
	}
}

func (s *Router) Route(r chi.Router) {
	r.Get("/group", s.HandleGetGroup)
	r.Get("/profile", s.HandleGetProfile)
	r.Get("/profile/details", s.HandleGetProfileDetails)
	r.Get("/order", s.HandleGetOrder)
	r.Get("/recordbook", s.HandleGetRecordBook)
	r.Get("/studentcard", s.HandleGetStudentCard)
}

// @Summary GetMyGroup
// @Tags OrgFaRu
// @Description get myGroup
// @Accept json
// @Produce json
// @Param input body orgfaclient.GetMyGroupInput true "myGroup info"
// @Success 200 {integer} integer 1
// @Router /finapp/api/auth/mygroup [get]
func (s *Router) HandleGetGroup(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetMyGroupInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.client.GetMyGroup(input)

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
func (s *Router) HandleGetRecordBook(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.AuthSession
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.client.GetRecordBook(input)

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
func (s *Router) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetMiniProfileInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.client.GetMiniProfile(input)

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
func (s *Router) HandleGetProfileDetails(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetProfileInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.client.GetProfile(input)

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
func (s *Router) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetOrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.client.GetOrder(input)

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
func (s *Router) HandleGetStudentCard(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "profileId")
	myGroup, err := s.client.GetStudentCard(&orgfaclient.GetStudentCardInput{
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
func (s *Router) HandlerGetStudyPlan(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetStudyPlanInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	myGroup, err := s.client.GetStudyPlan(input)

	if err != nil {
		fmt.Fprintf(w, "Get study plan:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get study plan marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
}