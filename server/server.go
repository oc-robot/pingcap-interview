package server

import (
	"log"
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
	write := func(code int, resp string) {
		w.WriteHeader(code)
		w.Write([]byte(resp))
		log.Printf("Path: %s, Resp: %d, %s", r.URL.Path, code, resp)
	}

	t := strings.TrimPrefix(r.URL.Path, "/latency/")
	_, err := time.ParseDuration(t)
	if err != nil {
		write(http.StatusBadRequest, "invalid request body")
		return
	}
	if err := s.exector.Exec(Change, t); err != nil {
		write(http.StatusInternalServerError, err.Error())
		return
	}
	write(http.StatusOK, "OK.")
}
