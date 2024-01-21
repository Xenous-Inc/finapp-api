package router

import (
	"net/http"
)

func (s *RootRouter) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
