package router

import (
	"net/http"
)

func (s *Router) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
