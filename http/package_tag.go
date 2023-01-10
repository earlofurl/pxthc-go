package http

import (
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/gobuffalo/nulls"
	"net/http"
	"strconv"
)

func (s *Server) getAllPackageTagsHandler(w http.ResponseWriter, r *http.Request) {
	var req pxthc.FindAllPackageTagsRequest
	err := req.Bind(r.Body)

	arg := &sqlc.ListPackageTagsParams{
		IsAssigned: req.IsAssigned,
		Limit:      int32(req.Limit),
		Offset:     int32(req.Offset),
	}

	tags, err := s.packageTagService.FindAllPackageTags(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, tags)
}

func (s *Server) getPackageTagByIDHandler(w http.ResponseWriter, r *http.Request) {
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

	tag, err := s.packageTagService.FindPackageTagByID(r.Context(), n)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, tag)
}

func (s *Server) getPackageTagByTagNumberHandler(w http.ResponseWriter, r *http.Request) {
	i := chi.URLParam(r, "tag")
	if i == "" {
		Json(w, http.StatusBadRequest, nil)
		return
	}

	tag, err := s.packageTagService.FindPackageTagByTagNumber(r.Context(), i)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, tag)
}

func (s *Server) updatePackageTagHandler(w http.ResponseWriter, r *http.Request) {
	var req pxthc.UpdatePackageTagRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.UpdatePackageTagParams{
		IsAssigned:        nulls.NewBool(req.IsAssigned),
		IsActive:          nulls.NewBool(req.IsActive),
		IsProvisional:     nulls.NewBool(req.IsProvisional),
		AssignedPackageID: nulls.NewInt64(req.AssignedPackageID),
		ID:                req.ID,
	}

	p, err := s.packageTagService.UpdatePackageTag(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, p)
}
