package http

import (
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
)

func (s *Server) getAllLabTestsHandler(w http.ResponseWriter, r *http.Request) {
	labTests, err := s.labTestService.FindAllLabTests(r.Context())
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, labTests)
}

func (s *Server) getLabTestByIDHandler(w http.ResponseWriter, r *http.Request) {
	l := chi.URLParam(r, "id")
	if l == "" {
		Json(w, http.StatusBadRequest, nil)
		return
	}

	n, err := strconv.ParseInt(l, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	if n < 1 {
		Json(w, http.StatusBadRequest, err)
		return
	}

	labTest, err := s.labTestService.FindLabTestByID(r.Context(), n)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, labTest)
}

func (s *Server) createLabTestHandler(w http.ResponseWriter, r *http.Request) {
	var req pxthc.CreateLabTestRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.CreateLabTestParams{
		TestName:                req.TestName,
		BatchCode:               req.BatchCode,
		TestIDCode:              req.TestIDCode,
		LabFacilityName:         req.LabFacilityName,
		TestPerformedDateTime:   req.TestPerformedDateTime,
		OverallPassed:           req.OverallPassed,
		TestTypeName:            req.TestTypeName,
		TestPassed:              req.TestPassed,
		TestComment:             req.TestComment,
		ThcTotalPercent:         req.ThcTotalPercent,
		ThcTotalValue:           req.ThcTotalValue,
		CbdPercent:              req.CbdPercent,
		CbdValue:                req.CbdValue,
		TerpeneTotalPercent:     req.TerpeneTotalPercent,
		TerpeneTotalValue:       req.TerpeneTotalValue,
		ThcAPercent:             req.ThcAPercent,
		ThcAValue:               req.ThcAValue,
		Delta9ThcPercent:        req.Delta9ThcPercent,
		Delta9ThcValue:          req.Delta9ThcValue,
		Delta8ThcPercent:        req.Delta8ThcPercent,
		Delta8ThcValue:          req.Delta8ThcValue,
		ThcVPercent:             req.ThcVPercent,
		ThcVValue:               req.ThcVValue,
		CbdAPercent:             req.CbdAPercent,
		CbdAValue:               req.CbdAValue,
		CbnPercent:              req.CbnPercent,
		CbnValue:                req.CbnValue,
		CbgAPercent:             req.CbgAPercent,
		CbgAValue:               req.CbgAValue,
		CbgPercent:              req.CbgPercent,
		CbgValue:                req.CbgValue,
		CbcPercent:              req.CbcPercent,
		CbcValue:                req.CbcValue,
		TotalCannabinoidPercent: req.TotalCannabinoidPercent,
		TotalCannabinoidValue:   req.TotalCannabinoidValue,
	}
	labTest, err := s.labTestService.CreateLabTest(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, labTest)
}

func (s *Server) updateLabTestHandler(w http.ResponseWriter, r *http.Request) {
	var req pxthc.UpdateLabTestRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.UpdateLabTestParams{
		TestName:                nulls.NewString(req.TestName),
		BatchCode:               nulls.NewString(req.BatchCode),
		TestIDCode:              nulls.NewString(req.TestIDCode),
		LabFacilityName:         nulls.NewString(req.LabFacilityName),
		TestPerformedDateTime:   nulls.NewTime(req.TestPerformedDateTime),
		OverallPassed:           nulls.NewBool(req.OverallPassed),
		TestTypeName:            nulls.NewString(req.TestTypeName),
		TestPassed:              nulls.NewBool(req.TestPassed),
		TestComment:             nulls.NewString(req.TestComment),
		ThcTotalPercent:         decimal.NewNullDecimal(req.ThcTotalPercent),
		ThcTotalValue:           decimal.NewNullDecimal(req.ThcTotalValue),
		CbdPercent:              decimal.NewNullDecimal(req.CbdPercent),
		CbdValue:                decimal.NewNullDecimal(req.CbdValue),
		TerpeneTotalPercent:     decimal.NewNullDecimal(req.TerpeneTotalPercent),
		TerpeneTotalValue:       decimal.NewNullDecimal(req.TerpeneTotalValue),
		ThcAPercent:             decimal.NewNullDecimal(req.ThcAPercent),
		ThcAValue:               decimal.NewNullDecimal(req.ThcAValue),
		Delta9ThcPercent:        decimal.NewNullDecimal(req.Delta9ThcPercent),
		Delta9ThcValue:          decimal.NewNullDecimal(req.Delta9ThcValue),
		Delta8ThcPercent:        decimal.NewNullDecimal(req.Delta8ThcPercent),
		Delta8ThcValue:          decimal.NewNullDecimal(req.Delta8ThcValue),
		ThcVPercent:             decimal.NewNullDecimal(req.ThcVPercent),
		ThcVValue:               decimal.NewNullDecimal(req.ThcVValue),
		CbdAPercent:             decimal.NewNullDecimal(req.CbdAPercent),
		CbdAValue:               decimal.NewNullDecimal(req.CbdAValue),
		CbnPercent:              decimal.NewNullDecimal(req.CbnPercent),
		CbnValue:                decimal.NewNullDecimal(req.CbnValue),
		CbgAPercent:             decimal.NewNullDecimal(req.CbgAPercent),
		CbgAValue:               decimal.NewNullDecimal(req.CbgAValue),
		CbgPercent:              decimal.NewNullDecimal(req.CbgPercent),
		CbgValue:                decimal.NewNullDecimal(req.CbgValue),
		CbcPercent:              decimal.NewNullDecimal(req.CbcPercent),
		CbcValue:                decimal.NewNullDecimal(req.CbcValue),
		TotalCannabinoidPercent: decimal.NewNullDecimal(req.TotalCannabinoidPercent),
		TotalCannabinoidValue:   decimal.NewNullDecimal(req.TotalCannabinoidValue),
		ID:                      nulls.NewInt64(req.ID),
	}
	labTest, err := s.labTestService.UpdateLabTest(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, labTest)
}

func (s *Server) deleteLabTestHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	i := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	if id < 1 {
		Json(w, http.StatusBadRequest, "invalid id")
		return
	}
	err = s.labTestService.DeleteLabTest(r.Context(), id)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, nil)
}
