package http

import (
	"encoding/json"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) createStrainHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Strain Create handler called")
	var req pxthc.CreateStrainRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.CreateStrainParams{
		Name: req.Name,
		Type: req.Type,
	}

	strain, err := s.strainService.CreateStrain(r.Context(), arg)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusCreated, strain)
}

func (s *Server) getAllStrainsHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Strain Get All handler called")
	var err error

	strains, err := s.strainService.FindAllStrains(r.Context())
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, strains)
}

func (s *Server) getStrainByNameHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Strain Get By Name handler called")
	var err error

	log.Debug().Msgf("Strain Name: %s", chi.URLParam(r, "name"))

	strainName := chi.URLParam(r, "name")
	if strainName == "" {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	// convert req.Name to lowercase
	n := strings.ToLower(strainName)
	if n == "" {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	strain, err := s.strainService.FindStrainByName(r.Context(), n)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, strain)
}

func (s *Server) getStrainByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Strain Get By ID handler called")
	var err error

	log.Debug().Msgf("Strain ID: %s", chi.URLParam(r, "id"))

	strainIDstring := chi.URLParam(r, "id")
	if strainIDstring == "" {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	strainID, err := strconv.ParseInt(strainIDstring, 10, 64)
	if err != nil {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	strain, err := s.strainService.FindStrainByID(r.Context(), strainID)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, strain)
}

func (s *Server) updateStrainHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Strain Update handler called")
	var req sqlc.UpdateStrainParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.UpdateStrainParams{
		Name:                    req.Name,
		Type:                    req.Type,
		YieldAverage:            req.YieldAverage,
		TerpAverageTotal:        req.TerpAverageTotal,
		Terp1:                   req.Terp1,
		Terp1Value:              req.Terp1Value,
		Terp2:                   req.Terp2,
		Terp2Value:              req.Terp2Value,
		Terp3:                   req.Terp3,
		Terp3Value:              req.Terp3Value,
		Terp4:                   req.Terp4,
		Terp4Value:              req.Terp4Value,
		Terp5:                   req.Terp5,
		Terp5Value:              req.Terp5Value,
		ThcAverage:              req.ThcAverage,
		TotalCannabinoidAverage: req.TotalCannabinoidAverage,
		LightDep2022:            req.LightDep2022,
		FallHarvest2022:         req.FallHarvest2022,
		QuantityAvailable:       req.QuantityAvailable,
		ID:                      req.ID,
	}

	strain, err := s.strainService.UpdateStrain(r.Context(), arg)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, strain)
}

func (s *Server) deleteStrainHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Strain Delete handler called")
	var err error

	log.Debug().Msgf("Strain ID: %s", chi.URLParam(r, "id"))

	strainIDstring := chi.URLParam(r, "id")
	if strainIDstring == "" {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	strainID, err := strconv.ParseInt(strainIDstring, 10, 64)
	if err != nil {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	err = s.strainService.DeleteStrain(r.Context(), strainID)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, nil)
}
