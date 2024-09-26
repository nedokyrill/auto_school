package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"newWebServer/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegicterRoutes(subrouter)

	log.Println("Server listen on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
