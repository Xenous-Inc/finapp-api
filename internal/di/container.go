package di

import (
	"github.com/Xenous-Inc/finapp-api/internal/server"
	"github.com/Xenous-Inc/finapp-api/internal/utils/config"
)

type Container struct {
	cfg *config.Config

	server *server.Server
}

func New(cfg *config.Config) *Container {
	return &Container{
		cfg: cfg,
	}
}

func (c *Container) GetServer() *server.Server {
	return get(&c.server, func() *server.Server {
		return server.NewServer(c.cfg.Port, c.cfg.Host)
	})
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		return *obj
	}

	*obj = builder()
	return *obj
}
