package http

import (
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/gobuffalo/nulls"
	"net/http"
	"strconv"
)

func (s *Server) getAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	items, err := s.itemService.FindAllItems(r.Context())
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, items)
}

func (s *Server) getItemByIDHandler(w http.ResponseWriter, r *http.Request) {
	i := chi.URLParam(r, "id")
	if i == "" {
		Json(w, http.StatusBadRequest, nil)
		return
	}

	n, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	if n < 1 {
		Json(w, http.StatusBadRequest, err)
		return
	}

	item, err := s.itemService.FindItemByID(r.Context(), n)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, item)
}

func (s *Server) createItemHandler(w http.ResponseWriter, r *http.Request) {
	var req pxthc.CreateItemRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.CreateItemParams{
		Description: req.Description,
		IsUsed:      req.IsUsed,
		ItemTypeID:  req.ItemTypeID,
		StrainID:    req.StrainID,
	}

	item, err := s.itemService.CreateItem(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusCreated, item)
}

func (s *Server) updateItemHandler(w http.ResponseWriter, r *http.Request) {
	var req pxthc.UpdateItemRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.UpdateItemParams{
		Description: nulls.NewString(req.Description),
		ID:          req.ID,
		IsUsed:      nulls.NewBool(req.IsUsed),
		ItemTypeID:  nulls.NewInt64(req.ItemTypeID),
		StrainID:    nulls.NewInt64(req.StrainID),
	}

	item, err := s.itemService.UpdateItem(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, item)
}

func (s *Server) deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	i := chi.URLParam(r, "id")
	if i == "" {
		Json(w, http.StatusBadRequest, nil)
		return
	}

	n, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	if n < 1 {
		Json(w, http.StatusBadRequest, err)
		return
	}

	err = s.itemService.DeleteItem(r.Context(), n)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusNoContent, nil)
}
