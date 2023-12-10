package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	port uint16
	host string
	*http.Server
}

func NewServer(port uint16, host string) *Server {
	newServer := &Server{port: port, host: host}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	newServer.Server = server

	return newServer
}

func (s *Server) StartListening() {
	log.Printf("Server successfully started on %s:%d", s.host, s.port)
	err := s.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
