package controllers

import (
	"github.com/jiharevzahar/fullstack/api/middlewares"
	"net/http"
)

func (s *Server) initializeRoutes() {

	//Users routes
	s.Router.HandleFunc("/groups/", middlewares.SetMiddlewareJSON(s.CreateGroup)).Methods(http.MethodPost)
	s.Router.HandleFunc("/groups", middlewares.SetMiddlewareJSON(s.GetGroups)).Methods(http.MethodGet)
	s.Router.HandleFunc("/groups/{id}", middlewares.SetMiddlewareJSON(s.GetGroup)).Methods(http.MethodGet)
	s.Router.HandleFunc("/groups/{id}", middlewares.SetMiddlewareJSON(s.UpdateGroup)).Methods(http.MethodPut)
	s.Router.HandleFunc("/groups/{id}", s.DeleteGroup).Methods(http.MethodDelete)

	//Posts routes
	s.Router.HandleFunc("/tasks/", middlewares.SetMiddlewareJSON(s.CreateTask)).Methods(http.MethodPost)
	s.Router.HandleFunc("/tasks", middlewares.SetMiddlewareJSON(s.GetTasks)).Methods(http.MethodGet)
	s.Router.HandleFunc("/tasks/{id}", middlewares.SetMiddlewareJSON(s.GetTask)).Methods(http.MethodGet)
	s.Router.HandleFunc("/tasks/{id}", middlewares.SetMiddlewareJSON(s.UpdateTask)).Methods(http.MethodPut)
	s.Router.HandleFunc("/tasks/{id}", s.DeleteTask).Methods(http.MethodDelete)

	//Timeframes routes
	s.Router.HandleFunc("/timeframes/", middlewares.SetMiddlewareJSON(s.CreateTimeframe)).Methods(http.MethodPost)
	s.Router.HandleFunc("/timeframes/{id}", s.DeleteTimeframe).Methods(http.MethodDelete)
}