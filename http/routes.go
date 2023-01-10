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
	s.initStrain()
	s.initProductCategory()
	s.initFacility()
	s.initFacilityLocation()
	s.initItemType()
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

func (s *Server) initStrain() {
	s.router.Route("/api/strain", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", s.getAllStrainsHandler)
		router.Get("/{id}", s.getStrainByIDHandler)
		router.Get("/name/{name}", s.getStrainByNameHandler)
		router.Post("/", s.createStrainHandler)
	})
}

func (s *Server) initProductCategory() {
	s.router.Route("/api/product-category", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", s.getAllProductCategoriesHandler)
		router.Get("/{id}", s.getProductCategoryByIDHandler)
		router.Get("/name/{name}", s.getProductCategoryByNameHandler)
	})
}

func (s *Server) initFacility() {
	s.router.Route("/api/facility", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", s.getAllFacilitiesHandler)
		router.Get("/{id}", s.getFacilityByIDHandler)
		router.Get("/name/{name}", s.getFacilityByNameHandler)
		router.Post("/", s.createFacilityHandler)
		router.Put("/", s.updateFacilityHandler)
		router.Delete("/", s.deleteFacilityHandler)
	})
}

func (s *Server) initFacilityLocation() {
	s.router.Route("/api/facility-location", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", s.getAllFacilityLocationsHandler)
		router.Get("/{id}", s.getFacilityLocationByIDHandler)
		router.Get("/name/{name}", s.getFacilityLocationByNameHandler)
		router.Post("/", s.createFacilityLocationHandler)
		router.Put("/", s.updateFacilityLocationHandler)
		router.Delete("/", s.deleteFacilityLocationHandler)
	})
}

func (s *Server) initItemType() {
	s.router.Route("/api/item-type", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", s.getAllItemTypesHandler)
		router.Get("/{id}", s.getItemTypeByIDHandler)
		router.Post("/", s.createItemTypeHandler)
		router.Put("/", s.updateItemTypeHandler)
		router.Delete("/", s.deleteItemTypeHandler)
	})
}
