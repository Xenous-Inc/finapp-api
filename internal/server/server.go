package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Xenous-Inc/finapp-api/internal/router"
	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
)

type Server struct {
	port uint16
	host string
	*http.Server
}

func NewServer(port uint16, host string, router *router.RootRouter) *Server {
	newServer := &Server{port: port, host: host}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	newServer.Server = server

	return newServer
}

func (s *Server) StartListening() {
	log.Info(fmt.Sprintf("Server successfully started on %s:%d", s.host, s.port))
	err := s.ListenAndServe()
	if err != nil {
		log.Error(err, "Internal", "server StartListening")
		panic("cannot start server")
	}
}
