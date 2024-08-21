package server

import (
	"net/http"
)

type Server struct {
	address string
	weight  int
}

func NewServer(address string, weight int) *Server {
	return &Server{address: address, weight: weight}
}

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Forward the request to the server
}
