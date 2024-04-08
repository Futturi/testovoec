package server

import (
	"net/http"
	"time"
)

type Server struct {
	Server *http.Server
}

func (s *Server) InitServ(port string, handler http.Handler) error {
	s.Server = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	return s.Server.ListenAndServe()
}
