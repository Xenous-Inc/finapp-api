package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt/v5"
)

type Router struct {
	client *orgfaclient.Client
	jwtSecret string 
}

func NewRouter(client *orgfaclient.Client, jwtSecret string) *Router {
	return &Router{
		client: client,
		jwtSecret: jwtSecret,
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

func (s *Router) HandleGetGroup(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		//errors.New("Токен отсутствует в заголовке Authorization")
		return 
	}
	
	tokenSlice := strings.Split(tokenString, "Bearer ")
	if len(tokenSlice) != 2 {
		return
	}
	
	tokenString = tokenSlice[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		sessionId := claims["sessionId"]
		fmt.Println(sessionId.(string))
		
	myGroup, err := s.client.GetMyGroup(&orgfaclient.GetMyGroupInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId.(string),
		},
	})

	if err != nil {
		fmt.Fprintf(w, "Get my group:  %s", clients.ErrUnauthorized)
	}

	res, err := json.Marshal(myGroup)

	if err != nil {
		fmt.Fprintf(w, "Get my group marshal:  %s", clients.ErrRequest)
	}

	fmt.Fprintf(w, string(res))
	}
}

func (s *Router) HandleGetRecordBook(w http.ResponseWriter, r *http.Request) {
	var input *orgfaclient.GetRecordBookInput
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

