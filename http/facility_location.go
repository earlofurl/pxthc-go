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

func (s *Server) getAllFacilityLocationsHandler(w http.ResponseWriter, r *http.Request) {
	facilityLocations, err := s.facilityLocationService.FindAllFacilityLocations(r.Context())
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, facilityLocations)
}

func (s *Server) getFacilityLocationByIDHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	facilityLocationIDstring := chi.URLParam(r, "id")
	if facilityLocationIDstring == "" {
		Json(w, http.StatusBadRequest, err)
		return
	}
	facilityLocationID, err := strconv.ParseInt(facilityLocationIDstring, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	facilityLocation, err := s.facilityLocationService.FindFacilityLocationByID(r.Context(), facilityLocationID)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, facilityLocation)
}

func (s *Server) getFacilityLocationByNameHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	facilityLocationName := chi.URLParam(r, "name")
	if facilityLocationName == "" {
		Json(w, http.StatusBadRequest, err)
		return
	}
	facilityLocation, err := s.facilityLocationService.FindFacilityLocationByName(r.Context(), facilityLocationName)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, facilityLocation)
}

func (s *Server) createFacilityLocationHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Create Facility Location handler called")
	var req pxthc.CreateFacilityLocationRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.CreateFacilityLocationParams{
		Name:       req.Name,
		FacilityID: req.FacilityID,
	}

	facilityLocation, err := s.facilityLocationService.CreateFacilityLocation(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, facilityLocation)
}

func (s *Server) updateFacilityLocationHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Update Facility Location handler called")
	var req pxthc.UpdateFacilityLocationRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.UpdateFacilityLocationParams{
		ID:         req.ID,
		Name:       nulls.NewString(req.Name),
		FacilityID: nulls.NewInt64(req.FacilityID),
	}

	facilityLocation, err := s.facilityLocationService.UpdateFacilityLocation(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, facilityLocation)
}

func (s *Server) deleteFacilityLocationHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Delete Facility Location handler called")
	var err error
	i := chi.URLParam(r, "id")
	if i == "" {
		Json(w, http.StatusBadRequest, err)
		return
	}

	n, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	if n <= 0 {
		Json(w, http.StatusBadRequest, err)
		return
	}

	err = s.facilityLocationService.DeleteFacilityLocation(r.Context(), n)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, nil)
}
