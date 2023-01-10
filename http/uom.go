package http

import (
	"github.com/earlofurl/pxthc"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

// We don't need much functionality for UoMs because they are static

func (s *Server) getUomByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Uom Get By ID handler called")
	var err error

	log.Debug().Msgf("Uom ID Param: %s", chi.URLParam(r, "id"))

	uomIDstring := chi.URLParam(r, "id")
	if uomIDstring == "" {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	uomID, err := strconv.ParseInt(uomIDstring, 10, 64)
	if err != nil {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	uom, err := s.uomService.FindUomByID(r.Context(), uomID)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, uom)
}

func (s *Server) getUomByNameHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Uom Get By Name handler called")
	var err error

	log.Debug().Msgf("Uom Name Param: %s", chi.URLParam(r, "name"))

	uomName := chi.URLParam(r, "name")
	if uomName == "" {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	// convert req.Name to lowercase
	n := strings.ToLower(uomName)
	if n == "" {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	uom, err := s.uomService.FindUomByName(r.Context(), uomName)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, uom)
}

func (s *Server) getAllUomsHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Uom Get All handler called")

	uoms, err := s.uomService.FindAllUoms(r.Context())
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	Json(w, http.StatusOK, uoms)
}
