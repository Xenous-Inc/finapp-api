package router

import (
	"net/http"

	_ "github.com/Xenous-Inc/finapp-api/docs"
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/router/auth"
	"github.com/Xenous-Inc/finapp-api/internal/router/classrooms"
	"github.com/Xenous-Inc/finapp-api/internal/router/groups"
	"github.com/Xenous-Inc/finapp-api/internal/router/teachers"
	"github.com/Xenous-Inc/finapp-api/internal/router/user"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type RootRouter struct {
	Client            *ruzfaclient.Client
	ClientOrgfaclient *orgfaclient.Client

	cfg *config.Config

	authRouter      *auth.Router
	classroomRouter *classrooms.Router
	groupsRouter    *groups.Router
	techersRouter   *teachers.Router
	userRouter      *user.Router
}

func NewRootRouter(ruzfaClient *ruzfaclient.Client, orgfaClient *orgfaclient.Client, cfg *config.Config) *RootRouter {
	return &RootRouter{
		Client:            ruzfaClient,
		ClientOrgfaclient: orgfaClient,

		userRouter:      user.NewRouter(orgfaClient),
		authRouter:      auth.NewRouter(orgfaClient),
		classroomRouter: classrooms.NewRouter(ruzfaClient),
		groupsRouter:    groups.NewRouter(ruzfaClient),
		techersRouter:   teachers.NewRouter(ruzfaClient),
	}
}

func (s *RootRouter) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/auth", s.authRouter.Route)
	r.Route("/user", s.userRouter.Route)
	r.Route("/classrooms", s.classroomRouter.Route)
	r.Route("/groups", s.groupsRouter.Route)
	r.Route("/teachers", s.techersRouter.Route)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5051/swagger/doc.json"), //The url pointing to API definition
	))

	return r
}
