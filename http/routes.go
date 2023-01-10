package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s *Server) InitRoutes() {
	s.initVersion()
	//s.initHealth()
	//s.initSwagger()
	s.initUser()
	s.initUom()
}

func (s *Server) initVersion() {
	s.router.Route("/version", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			Json(w, http.StatusOK, map[string]string{"version": s.Version})
		})
	})
}

func (s *Server) initUser() {
	s.router.Route("/api/user", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", s.getUserByEmailHandler)
		router.Post("/", s.createUserHandler)
		router.Post("/login", s.loginUserHandler)
	})
}

func (s *Server) initUom() {
	s.router.Route("/api/uom", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", s.getAllUomsHandler)
		router.Get("/{id}", s.getUomByIDHandler)
		router.Get("/name/{name}", s.getUomByNameHandler)
	})
}
