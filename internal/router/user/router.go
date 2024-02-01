package user

import (
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/clients"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/dto"
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
	r.Get("/studentcard", s.HandleGetStudentCard)
	r.Get("/studyplan", s.HandlerGetStudyPlan)
}

func (s *Router) HandleGetGroup(w http.ResponseWriter, r *http.Request) {
	token, err := jwtservice.GetDecodeToken(r, s.client.Cfg.JwtSecret)

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

	response, err := s.client.GetMyGroup(&orgfaclient.GetMyGroupInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId,
		},
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "user HandleGetGroup")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "user HandleGetGroup")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "user HandleGetGroup")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetGroup")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "user HandleGetGroup")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	// items := make([]dto.MyGroup, 0)

	// for _, group := range response {
	// 	items = append(items, dto.MyGroupFromClientModel(&group))
	// }

	responser.Success(w, r, response)
}

func (s *Router) HandleGetRecordBook(w http.ResponseWriter, r *http.Request) {
	token, err := jwtservice.GetDecodeToken(r, s.client.Cfg.JwtSecret)

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

	response, err := s.client.GetRecordBook(&orgfaclient.GetRecordBookInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId,
		},
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "user HandleGetRecordBook")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "user HandleGetRecordBook")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "user HandleGetRecordBook")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetRecordBook")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "user HandleGetRecordBook")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, response)
}

// @Summary Try to get profile in user
// @Description In success case returns profile
// @Tags user
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dto.Profile
// @Failure 401 {object} dto.ApiError
// @Failure 400 {object} dto.ApiError
// @Failure 500 {object} dto.ApiError
// @Router /user/profile [get]
func (s *Router) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	token, err := jwtservice.GetDecodeToken(r, s.client.Cfg.JwtSecret)

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

	response, err := s.client.GetProfile(&orgfaclient.GetMiniProfileInput{
		AuthSession: &orgfaclient.AuthSession{
			SessionId: sessionId,
		},
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "user HandleGetProfile")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "user HandleGetProfile")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "user HandleGetProfile")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetProfile")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "user HandleGetProfile")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	profile := dto.ProfileFromClientModel(response)

	responser.Success(w, r, profile)
}

func (s *Router) HandleGetProfileDetails(w http.ResponseWriter, r *http.Request) {
	token, err := jwtservice.GetDecodeToken(r, s.client.Cfg.JwtSecret)

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
			log.Error(err, "BadRequest", "user HandleGetProfileDetails")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "user HandleGetProfileDetails")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "user HandleGetProfileDetails")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetProfileDetails")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "user HandleGetProfileDetails")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, profile)

}

func (s *Router) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	token, err := jwtservice.GetDecodeToken(r, s.client.Cfg.JwtSecret)

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
			log.Error(err, "BadRequest", "user HandleGetOrder")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "user HandleGetOrder")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "user HandleGetOrder")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetOrder")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "user HandleGetOrder")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	// items := make([]dto.ScheduleItem, 0)

	// for _, schedule := range groupsSchedule {
	// 	items = append(items, dto.ScheduleItemFromClientModel(&schedule))
	// }

	responser.Success(w, r, order)
}

func (s *Router) HandleGetStudentCard(w http.ResponseWriter, r *http.Request) {
	profileId := r.URL.Query().Get("profileId")
	err := s.validator.Var(profileId, "required")
	if err != nil {
		log.Error(err, "BadRequest", "teachers HandleGetTeacher")
		responser.BadRequset(w, r, "Query parameter `term` must be provided")

		return
	}

	token, err := jwtservice.GetDecodeToken(r, s.client.Cfg.JwtSecret)

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
		}, ProfileId: profileId,
	})

	if err != nil {
		switch err {
		case clients.ErrRequest:
			log.Error(err, "BadRequest", "user HandleGetStudentCard")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "user HandleGetStudentCard")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "user HandleGetStudentCard")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandleGetStudentCard")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "user HandleGetStudentCard")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	responser.Success(w, r, studentCard)
}

func (s *Router) HandlerGetStudyPlan(w http.ResponseWriter, r *http.Request) {
	token, err := jwtservice.GetDecodeToken(r, s.client.Cfg.JwtSecret)

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
			log.Error(err, "BadRequest", "user HandlerGetStudyPlan")
			responser.BadRequset(w, r, "Invalid request")
		case clients.ErrInvalidEntity:
			log.Error(err, "Invalid Entity", "user HandlerGetStudyPlan")
			responser.BadRequset(w, r, "Invalid entity")
		case clients.ErrValidation:
			log.Error(err, "Error Validation", "user HandlerGetStudyPlan")
			responser.BadRequset(w, r, "Error validation")
		case clients.ErrUnauthorized:
			log.Error(err, "Unauthorized", "user HandlerGetStudyPlan")
			responser.Unauthorized(w, r)
		default:
			log.Error(err, "Internal", "user HandlerGetStudyPlan")
			responser.Internal(w, r, err.Error())
		}

		return
	}

	// items := make([]dto.ScheduleItem, 0)

	// for _, schedule := range groupsSchedule {
	// 	items = append(items, dto.ScheduleItemFromClientModel(&schedule))
	// }

	responser.Success(w, r, studyPlan)
}
