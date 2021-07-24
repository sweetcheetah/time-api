package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(router *mux.Router) *Server {
	s := Server{
		router: router,
	}
	s.routes()
	s.middleware()

	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Start() error {
	if os.Getenv("PORT") == "" {
		log.Fatal("PORT not set in environment")
	}

	log.Printf("Listening on port %v", os.Getenv("PORT"))
	return http.ListenAndServe(":"+os.Getenv("PORT"), s.router)
}
