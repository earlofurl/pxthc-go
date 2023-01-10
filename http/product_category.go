package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (s *Server) getAllProductCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	productCategories, err := s.store.ListProductCategories(r.Context())
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, productCategories)
}

func (s *Server) getProductCategoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	productCategoryIDstring := chi.URLParam(r, "id")
	if productCategoryIDstring == "" {
		Json(w, http.StatusBadRequest, err)
		return
	}
	productCategoryID, err := strconv.ParseInt(productCategoryIDstring, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	productCategory, err := s.store.GetProductCategoryByID(r.Context(), productCategoryID)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, productCategory)
}

func (s *Server) getProductCategoryByNameHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	productCategoryName := chi.URLParam(r, "name")
	if productCategoryName == "" {
		Json(w, http.StatusBadRequest, err)
		return
	}
	productCategory, err := s.store.GetProductCategoryByName(r.Context(), productCategoryName)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, productCategory)
}
