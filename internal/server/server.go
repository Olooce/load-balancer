package server

import (
	"io"
	"net/http"
	"net/url"
)

type Server struct {
	address string
	weight  int
}

func NewServer(address string, weight int) *Server {
	return &Server{address: address, weight: weight}
}

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the server address
	targetURL, err := url.Parse(s.address)
	if err != nil {
		http.Error(w, "Failed to parse server address", http.StatusInternalServerError)
		return
	}

	// Create a new request to forward to the target server
	req, err := http.NewRequest(r.Method, targetURL.String(), r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Copy headers from the original request
	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	// Forward the request to the target server
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to forward request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy response headers and body back to the original response
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
