package http

import (
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/gobuffalo/nulls"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (s *Server) getAllFacilitiesHandler(w http.ResponseWriter, r *http.Request) {
	facilities, err := s.facilityService.FindAllFacilities(r.Context())
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, facilities)
}

func (s *Server) getFacilityByIDHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	facilityIDstring := chi.URLParam(r, "id")
	if facilityIDstring == "" {
		Json(w, http.StatusBadRequest, err)
	}
	facilityID, err := strconv.ParseInt(facilityIDstring, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	facility, err := s.facilityService.FindFacilityByID(r.Context(), facilityID)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, facility)
}

func (s *Server) getFacilityByNameHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	facilityName := chi.URLParam(r, "name")
	if facilityName == "" {
		Json(w, http.StatusBadRequest, err)
		return
	}
	facility, err := s.facilityService.FindFacilityByName(r.Context(), facilityName)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, facility)
}

func (s *Server) updateFacilityHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Update Facility handler called")
	var req pxthc.UpdateFacilityRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.UpdateFacilityParams{
		ID:            req.ID,
		Name:          nulls.NewString(req.Name),
		LicenseNumber: nulls.NewString(req.LicenseNumber),
	}

	f, err := s.facilityService.UpdateFacility(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, f)
}

func (s *Server) createFacilityHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Create Facility handler called")
	var req pxthc.CreateFacilityRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.CreateFacilityParams{
		Name:          req.Name,
		LicenseNumber: req.LicenseNumber,
	}

	f, err := s.facilityService.CreateFacility(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, f)
}

func (s *Server) deleteFacilityHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	facilityIDstring := chi.URLParam(r, "id")
	if facilityIDstring == "" {
		Json(w, http.StatusBadRequest, err)
		return
	}
	facilityID, err := strconv.ParseInt(facilityIDstring, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	err = s.facilityService.DeleteFacility(r.Context(), facilityID)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, nil)
}
