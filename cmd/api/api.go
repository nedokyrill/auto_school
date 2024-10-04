package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"newWebServer/service/exam"
	"newWebServer/service/student"
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
	userHandler.RegisterRoutes(subrouter)

	studentStore := student.NewStore(s.db)
	studentHandler := student.NewHandler(studentStore)
	studentHandler.RegisterRoutes(subrouter)

	examStore := exam.NewStore(s.db)
	examHandler := exam.NewHandler(examStore)
	examHandler.RegisterRoutes(subrouter)

	log.Println("Server listen on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
