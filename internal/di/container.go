package di

import (
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/router"
	"github.com/Xenous-Inc/finapp-api/internal/server"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
)

type Container struct {
	cfg *config.Config

	server *server.Server

	router *router.Router
}

func New(cfg *config.Config) *Container {
	return &Container{
		cfg: cfg,
	}
}

func (c *Container) GetServer(cl *ruzfaclient.Client) *server.Server {
	return get(&c.server, func() *server.Server {
		return server.NewServer(c.cfg.Port, c.cfg.Host, c.GetRouter(cl))
	})
}

func (c *Container) GetRouter(cl *ruzfaclient.Client) *router.Router {
	return get(&c.router, func() *router.Router {
		return router.NewRouter(cl)
	})
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		return *obj
	}

	*obj = builder()
	return *obj
}
