package user

import (
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/router/utils/responser"
	"github.com/Xenous-Inc/finapp-api/internal/utils/jwtservice"

	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
	"github.com/go-chi/chi"
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

	token, err := jwtservice.GetDecodeToken(tokenString, s.client.Cfg.JwtSecret)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetGroup")
		responser.Unauthorized(w, r)
		return
	}

	sessionId, err := jwtservice.GetSessionIdFromToken(token)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetGroup")
		responser.Unauthorized(w, r)
		return
	}

	myGroup, err := s.client.GetMyGroup(&orgfaclient.GetMyGroupInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId,
		},
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "schedule HandleGetGroup")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "schedule HandleGetGroup")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "schedule HandleGetGroup")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetGroup")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "schedule HandleGetGroup")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, myGroup)
}

func (s *Router) HandleGetRecordBook(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	token, err := jwtservice.GetDecodeToken(tokenString, s.client.Cfg.JwtSecret)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetRecordBook")
		responser.Unauthorized(w, r)
		return
	}

	sessionId, err := jwtservice.GetSessionIdFromToken(token)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetRecordBook")
		responser.Unauthorized(w, r)
		return
	}

	recordBook, err := s.client.GetRecordBook(&orgfaclient.GetRecordBookInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId,
		},
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "schedule HandleGetRecordBook")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "schedule HandleGetRecordBook")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "schedule HandleGetRecordBook")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetRecordBook")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "schedule HandleGetRecordBook")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, recordBook)
}

func (s *Router) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	token, err := jwtservice.GetDecodeToken(tokenString, s.client.Cfg.JwtSecret)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetProfile")
		responser.Unauthorized(w, r)
		return
	}

	sessionId, err := jwtservice.GetSessionIdFromToken(token)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetProfile")
		responser.Unauthorized(w, r)
		return
	}

	miniProfile, err := s.client.GetProfile(&orgfaclient.GetMiniProfileInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId,
		},
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "schedule HandleGetProfile")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "schedule HandleGetProfile")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "schedule HandleGetProfile")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetProfile")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "schedule HandleGetProfile")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, miniProfile)
}

func (s *Router) HandleGetProfileDetails(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	token, err := jwtservice.GetDecodeToken(tokenString, s.client.Cfg.JwtSecret)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetProfileDetails")
		responser.Unauthorized(w, r)
		return
	}

	sessionId, err := jwtservice.GetSessionIdFromToken(token)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetProfileDetails")
		responser.Unauthorized(w, r)
		return
	}

	profile, err := s.client.GetProfileDetails(&orgfaclient.GetProfileInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId,
		},
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "schedule HandleGetProfileDetails")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "schedule HandleGetProfileDetails")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "schedule HandleGetProfileDetails")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetProfileDetails")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "schedule HandleGetProfileDetails")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, profile)

}

func (s *Router) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	token, err := jwtservice.GetDecodeToken(tokenString, s.client.Cfg.JwtSecret)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetOrder")
		responser.Unauthorized(w, r)
		return
	}

	sessionId, err := jwtservice.GetSessionIdFromToken(token)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetOrder")
		responser.Unauthorized(w, r)
		return
	}

	order, err := s.client.GetOrder(&orgfaclient.GetOrderInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId,
		},
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "schedule HandleGetOrder")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "schedule HandleGetOrder")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "schedule HandleGetOrder")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetOrder")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "schedule HandleGetOrder")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, order)
}

func (s *Router) HandleGetStudentCard(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "profileId")

	tokenString := r.Header.Get("Authorization")

	token, err := jwtservice.GetDecodeToken(tokenString, s.client.Cfg.JwtSecret)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetStudentCard")
		responser.Unauthorized(w, r)
		return
	}

	sessionId, err := jwtservice.GetSessionIdFromToken(token)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandleGetStudentCard")
		responser.Unauthorized(w, r)
		return
	}

	studentCard, err := s.client.GetStudentCard(&orgfaclient.GetStudentCardInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId,
		}, ProfileId: url,
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "schedule HandleGetStudentCard")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "schedule HandleGetStudentCard")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "schedule HandleGetStudentCard")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetStudentCard")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "schedule HandleGetStudentCard")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, studentCard)
}

func (s *Router) HandlerGetStudyPlan(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	token, err := jwtservice.GetDecodeToken(tokenString, s.client.Cfg.JwtSecret)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandlerGetStudyPlan")
		responser.Unauthorized(w, r)
		return
	}

	sessionId, err := jwtservice.GetSessionIdFromToken(token)

	if err != nil {
		log.Error(err, "Unauthorized", "user HandlerGetStudyPlan")
		responser.Unauthorized(w, r)
		return
	}

	studyPlan, err := s.client.GetStudyPlan(&orgfaclient.GetStudyPlanInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId,
		},
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "schedule HandlerGetStudyPlan")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "schedule HandlerGetStudyPlan")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "schedule HandlerGetStudyPlan")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandlerGetStudyPlan")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "schedule HandlerGetStudyPlan")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, studyPlan)
}
