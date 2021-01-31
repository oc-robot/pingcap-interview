package server

import (
	"net/http"
	"strings"
	"time"
)

type server struct {
	exector Exector
}

// NewServer return http.Handler that impl /latency
func NewServer(exector Exector) http.Handler {
	return &server{exector: exector}
}

// ServeHTTP impl http.Handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := strings.TrimPrefix(r.URL.Path, "/latency/")
	_, err := time.ParseDuration(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request body"))
		return
	}
	if err := s.exector.Exec(Change, t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
