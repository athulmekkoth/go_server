package api

import (
	"database/sql"
	"log"
	"net/http"

	user "github.com/athulmekkoth/go_server.git/service"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) RUN() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/vi").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Println("listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
