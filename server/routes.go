package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

func (s *Server) routes() {
	s.router.HandleFunc("/localtime", s.handlerLocalTime)
	s.router.HandleFunc("/", handlerRoot)

	//swagger
	fs := http.FileServer(http.Dir("./swagger-ui/"))
	s.router.PathPrefix("/swagger-ui").Handler(http.StripPrefix("/swagger-ui", fs))
}

func (s *Server) middleware() {
	s.router.Use(loggingMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found."))
}

func (s *Server) handlerLocalTime(w http.ResponseWriter, r *http.Request) {
	tz := os.Getenv("TIMEZONE")
	loc, e := time.LoadLocation(tz)

	if e != nil {
		log.Fatal(e)
	}

	t := time.Now()
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	timestring := t.In(loc).String()
	w.Write([]byte(timestring))
}
