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
		//router.Get("/{id}", s.UserService.FindUserById)
		//router.Get("/email", s.UserService.FindUserByEmail)
		//router.Put("/{id}", s.UserService.UpdateUser)
		//router.Delete("/{id}", s.UserService.DeleteUser)
	})
}
