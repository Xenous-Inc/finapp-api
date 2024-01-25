package di

import (
	"github.com/Xenous-Inc/finapp-api/internal/clients/orgfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient"
	"github.com/Xenous-Inc/finapp-api/internal/router"
	"github.com/Xenous-Inc/finapp-api/internal/server"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
)

type Container struct {
	cfg *config.Config

	server *server.Server

	router *router.RootRouter

	ruzFaClient *ruzfaclient.Client
	orgFaClient *orgfaclient.Client
}

func New(cfg *config.Config) *Container {
	return &Container{
		cfg: cfg,
	}
}

func (c *Container) GetRuzClient() *ruzfaclient.Client {
	return get(&c.ruzFaClient, func() *ruzfaclient.Client {
		return ruzfaclient.NewClient("https://ruz.fa.ru/api/", c.cfg)
	})
}

func (c *Container) GetOrgClient() *orgfaclient.Client {
	return get(&c.orgFaClient, func() *orgfaclient.Client {
		return orgfaclient.NewClient("https://org.fa.ru/", c.cfg)
	})
}

func (c *Container) GetServer() *server.Server {
	return get(&c.server, func() *server.Server {
		return server.NewServer(c.cfg.Port, c.cfg.Host, c.GetRouter(c.GetRuzClient(), c.GetOrgClient(), c.cfg))
	})
}

func (c *Container) GetRouter(cl *ruzfaclient.Client, clO *orgfaclient.Client, cfg *config.Config) *router.RootRouter {
	return get(&c.router, func() *router.RootRouter {
		return router.NewRootRouter(cl, clO, c.cfg)
	})
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		log.Warn("Internal", "di get")
		return *obj
	}

	*obj = builder()
	return *obj
}
