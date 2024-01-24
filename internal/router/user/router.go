package user

import (
	"net/http"
	"strings"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/router/utils/responser"
	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/go-playground/validator.v9"
)

type Router struct {
	client    *orgfaclient.Client
	validator *validator.Validate
}

func NewRouter(client *orgfaclient.Client) *Router {
	return &Router{
		client:    client,
		validator: validator.New(),
	}
}

func (s *Router) Route(r chi.Router) {
	r.Get("/group", s.HandleGetGroup)
	r.Get("/profile", s.HandleGetProfile)
	r.Get("/profile/details", s.HandleGetProfileDetails)
	r.Get("/order", s.HandleGetOrder)
	r.Get("/recordbook", s.HandleGetRecordBook)
	r.Get("/studentcard/{profileId}", s.HandleGetStudentCard)
	r.Get("/studyplan", s.HandlerGetStudyPlan)
}

func (s *Router) HandleGetGroup(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		log.Warn("Authorization header is empty")
		responser.Unauthorized(w, r)
		return
	}

	tokenSlice := strings.Split(tokenString, "Bearer ")
	if len(tokenSlice) != 2 {
		log.Warn("Invalid Authorization header format")
		responser.Unauthorized(w, r)
		return
	}

	tokenString = tokenSlice[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.client.Cfg.JwtSecret), nil
	})

	if err != nil {
		responser.Unauthorized(w, r)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		sessionId := claims["sessionId"]

		myGroup, err := s.client.GetMyGroup(&orgfaclient.GetMyGroupInput{
			AuthSession: &orgfaclient.AuthSession{
				SessionId: sessionId.(string),
			},
		})

		if err != nil {
			switch err {
			case clients.ErrRequest:
				responser.BadRequset(w, r, "Invalid request")
			case clients.ErrInvalidEntity:
				responser.BadRequset(w, r, "Invalid entity")
			case clients.ErrValidation:
				responser.BadRequset(w, r, "Error validation")
			case clients.ErrUnauthorized:
				responser.Unauthorized(w, r)
			default:
				responser.Internal(w, r, err.Error())
			}

			return
		}

		responser.Success(w, r, myGroup)
	}
}

func (s *Router) HandleGetRecordBook(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		log.Warn("Authorization header is empty")
		responser.Unauthorized(w, r)
		return
	}

	tokenSlice := strings.Split(tokenString, "Bearer ")
	if len(tokenSlice) != 2 {
		log.Warn("Invalid Authorization header format")
		responser.Unauthorized(w, r)
		return
	}

	tokenString = tokenSlice[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.client.Cfg.JwtSecret), nil
	})

	if err != nil {
		responser.Unauthorized(w, r)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		sessionId := claims["sessionId"]

		recordBook, err := s.client.GetRecordBook(&orgfaclient.GetRecordBookInput{
			AuthSession: &orgfaclient.AuthSession{
				SessionId: sessionId.(string),
			},
		})

		if err != nil {
			switch err {
			case clients.ErrRequest:
				responser.BadRequset(w, r, "Invalid request")
			case clients.ErrInvalidEntity:
				responser.BadRequset(w, r, "Invalid entity")
			case clients.ErrValidation:
				responser.BadRequset(w, r, "Error validation")
			case clients.ErrUnauthorized:
				responser.Unauthorized(w, r)
			default:
				responser.Internal(w, r, err.Error())
			}

			return
		}

		responser.Success(w, r, recordBook)
	}
}

func (s *Router) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		log.Warn("Authorization header is empty")
		responser.Unauthorized(w, r)
		return
	}

	tokenSlice := strings.Split(tokenString, "Bearer ")
	if len(tokenSlice) != 2 {
		log.Warn("Invalid Authorization header format")
		responser.Unauthorized(w, r)
		return
	}

	tokenString = tokenSlice[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.client.Cfg.JwtSecret), nil
	})

	if err != nil {
		responser.Unauthorized(w, r)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		sessionId := claims["sessionId"]

		miniProfile, err := s.client.GetMiniProfile(&orgfaclient.GetMiniProfileInput{
			AuthSession: &orgfaclient.AuthSession{
				SessionId: sessionId.(string),
			},
		})

		if err != nil {
			switch err {
			case clients.ErrRequest:
				responser.BadRequset(w, r, "Invalid request")
			case clients.ErrInvalidEntity:
				responser.BadRequset(w, r, "Invalid entity")
			case clients.ErrValidation:
				responser.BadRequset(w, r, "Error validation")
			case clients.ErrUnauthorized:
				responser.Unauthorized(w, r)
			default:
				responser.Internal(w, r, err.Error())
			}

			return
		}

		responser.Success(w, r, miniProfile)
	}
}

func (s *Router) HandleGetProfileDetails(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		log.Warn("Authorization header is empty")
		responser.Unauthorized(w, r)
		return
	}

	tokenSlice := strings.Split(tokenString, "Bearer ")
	if len(tokenSlice) != 2 {
		log.Warn("Invalid Authorization header format")
		responser.Unauthorized(w, r)
		return
	}

	tokenString = tokenSlice[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.client.Cfg.JwtSecret), nil
	})

	if err != nil {
		responser.Unauthorized(w, r)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		sessionId := claims["sessionId"]

		profile, err := s.client.GetProfile(&orgfaclient.GetProfileInput{
			AuthSession: &orgfaclient.AuthSession{
				SessionId: sessionId.(string),
			},
		})

		if err != nil {
			switch err {
			case clients.ErrRequest:
				responser.BadRequset(w, r, "Invalid request")
			case clients.ErrInvalidEntity:
				responser.BadRequset(w, r, "Invalid entity")
			case clients.ErrValidation:
				responser.BadRequset(w, r, "Error validation")
			case clients.ErrUnauthorized:
				responser.Unauthorized(w, r)
			default:
				responser.Internal(w, r, err.Error())
			}

			return
		}

		responser.Success(w, r, profile)

	}
}

func (s *Router) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		log.Warn("Authorization header is empty")
		responser.Unauthorized(w, r)
		return
	}

	tokenSlice := strings.Split(tokenString, "Bearer ")
	if len(tokenSlice) != 2 {
		log.Warn("Invalid Authorization header format")
		responser.Unauthorized(w, r)
		return
	}

	tokenString = tokenSlice[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.client.Cfg.JwtSecret), nil
	})

	if err != nil {
		responser.Unauthorized(w, r)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		sessionId := claims["sessionId"]

		order, err := s.client.GetOrder(&orgfaclient.GetOrderInput{
			AuthSession: &orgfaclient.AuthSession{
				SessionId: sessionId.(string),
			},
		})

		if err != nil {
			switch err {
			case clients.ErrRequest:
				responser.BadRequset(w, r, "Invalid request")
			case clients.ErrInvalidEntity:
				responser.BadRequset(w, r, "Invalid entity")
			case clients.ErrValidation:
				responser.BadRequset(w, r, "Error validation")
			case clients.ErrUnauthorized:
				responser.Unauthorized(w, r)
			default:
				responser.Internal(w, r, err.Error())
			}

			return
		}

		responser.Success(w, r, order)
	}
}

func (s *Router) HandleGetStudentCard(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "profileId")

	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		log.Warn("Authorization header is empty")
		responser.Unauthorized(w, r)
		return
	}

	tokenSlice := strings.Split(tokenString, "Bearer ")
	if len(tokenSlice) != 2 {
		log.Warn("Invalid Authorization header format")
		responser.Unauthorized(w, r)
		return
	}

	tokenString = tokenSlice[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.client.Cfg.JwtSecret), nil
	})

	if err != nil {
		responser.Unauthorized(w, r)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		sessionId := claims["sessionId"]

		studentCard, err := s.client.GetStudentCard(&orgfaclient.GetStudentCardInput{
			AuthSession: &orgfaclient.AuthSession{
				SessionId: sessionId.(string),
			}, ProfileId: url,
		})

		if err != nil {
			switch err {
			case clients.ErrRequest:
				responser.BadRequset(w, r, "Invalid request")
			case clients.ErrInvalidEntity:
				responser.BadRequset(w, r, "Invalid entity")
			case clients.ErrValidation:
				responser.BadRequset(w, r, "Error validation")
			case clients.ErrUnauthorized:
				responser.Unauthorized(w, r)
			default:
				responser.Internal(w, r, err.Error())
			}

			return
		}

		responser.Success(w, r, studentCard)
	}
}

func (s *Router) HandlerGetStudyPlan(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		log.Warn("Authorization header is empty")
		responser.Unauthorized(w, r)
		return
	}

	tokenSlice := strings.Split(tokenString, "Bearer ")
	if len(tokenSlice) != 2 {
		log.Warn("Invalid Authorization header format")
		responser.Unauthorized(w, r)
		return
	}

	tokenString = tokenSlice[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.client.Cfg.JwtSecret), nil
	})

	if err != nil {
		responser.Unauthorized(w, r)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		sessionId := claims["sessionId"]

		studyPlan, err := s.client.GetStudyPlan(&orgfaclient.GetStudyPlanInput{
			AuthSession: &orgfaclient.AuthSession{
				SessionId: sessionId.(string),
			},
		})

		if err != nil {
			switch err {
			case clients.ErrRequest:
				responser.BadRequset(w, r, "Invalid request")
			case clients.ErrInvalidEntity:
				responser.BadRequset(w, r, "Invalid entity")
			case clients.ErrValidation:
				responser.BadRequset(w, r, "Error validation")
			case clients.ErrUnauthorized:
				responser.Unauthorized(w, r)
			default:
				responser.Internal(w, r, err.Error())
			}

			return
		}

		responser.Success(w, r, studyPlan)
	}
}
