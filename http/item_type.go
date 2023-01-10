package http

import (
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/gobuffalo/nulls"
	"net/http"
	"strconv"
)

func (s *Server) getAllItemTypesHandler(w http.ResponseWriter, r *http.Request) {
	itemTypes, err := s.itemTypeService.FindAllItemTypes(r.Context())
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, itemTypes)
}

func (s *Server) getItemTypeByIDHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	itemTypeID := chi.URLParam(r, "id")
	if itemTypeID == "" {
		Json(w, http.StatusBadRequest, err)
		return
	}

	i, err := strconv.ParseInt(itemTypeID, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	if i < 1 {
		Json(w, http.StatusBadRequest, err)
		return
	}

	itemType, err := s.itemTypeService.FindItemTypeByID(r.Context(), i)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, itemType)
}

func (s *Server) createItemTypeHandler(w http.ResponseWriter, r *http.Request) {
	var req pxthc.CreateItemTypeRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.CreateItemTypeParams{
		ProductForm:       req.ProductForm,
		ProductModifier:   req.ProductModifier,
		UomDefault:        req.UomDefault,
		ProductCategoryID: req.ProductCategoryID,
	}

	itemType, err := s.itemTypeService.CreateItemType(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, itemType)
}

func (s *Server) updateItemTypeHandler(w http.ResponseWriter, r *http.Request) {
	var req pxthc.UpdateItemTypeRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.UpdateItemTypeParams{
		ProductForm:       nulls.NewString(req.ProductForm),
		ProductModifier:   nulls.NewString(req.ProductModifier),
		UomDefault:        nulls.NewInt64(req.UomDefault),
		ProductCategoryID: nulls.NewInt64(req.ProductCategoryID),
		ID:                req.ID,
	}

	itemType, err := s.itemTypeService.UpdateItemType(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, itemType)
}

func (s *Server) deleteItemTypeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	itemTypeID := chi.URLParam(r, "id")
	if itemTypeID == "" {
		Json(w, http.StatusBadRequest, err)
		return
	}

	i, err := strconv.ParseInt(itemTypeID, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	if i < 1 {
		Json(w, http.StatusBadRequest, err)
		return
	}

	err = s.itemTypeService.DeleteItemType(r.Context(), i)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, nil)
}
