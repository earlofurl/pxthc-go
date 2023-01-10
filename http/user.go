package http

import (
	"database/sql"
	"encoding/json"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/postgres"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("User Create handler called")
	var req pxthc.CreateUserRequest
	log.Debug().Msgf("req: %+v", req)
	//err := s.readJSON(w, r, &pxthc.CreateUserRequest{})
	err := req.Bind(r.Body)
	if err != nil {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	// TODO: implement validation
	//errs := validate.Validate(h.validate, req)
	//if errs != nil {
	//	respond.Errors(w, http.StatusBadRequest, errs)
	//	return
	//}

	user := &pxthc.CreateUserRequest{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Role:      req.Role,
	}

	log.Debug().Msgf("user: %+v", user)

	createdUser, err := s.userService.CreateUser(r.Context(), user)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusCreated, createdUser)
}

func (s *Server) getUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("User Get By Email handler called")
	var req pxthc.GetUserByEmailRequest
	//err := s.readJSON(w, r, &req)
	err := req.Bind(r.Body)
	if err != nil {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	// TODO: implement validation
	//errs := validate.Validate(h.validate, req)
	//if errs != nil {
	//	respond.Errors(w, http.StatusBadRequest, errs)
	//	return
	//}

	userRequestData := &pxthc.GetUserByEmailRequest{
		Email: req.Email,
	}

	log.Debug().Msgf("userRequestData: %+v", userRequestData)

	userData, err := s.userService.FindUserByEmail(r.Context(), userRequestData.Email)
	if err != nil {
		log.Err(err).Msg("error getting user")
		if err == sql.ErrNoRows {
			pxthc.Error(w, http.StatusNotFound, err)
			return
		}
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	if userData == nil {
		pxthc.Error(w, http.StatusNotFound, err)
		return
	}
	Json(w, http.StatusOK, userData)
}

func (s *Server) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("User Login handler called")
	var req pxthc.LoginUserRequest
	//err := s.readJSON(w, r, &req)
	err := req.Bind(r.Body)
	if err != nil {
		pxthc.Error(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Msgf("req: %+v", req)

	// TODO: Validate

	log.Debug().Msg("Calling FindUserByEmail")
	u, err := s.userService.FindUserByEmail(r.Context(), req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			pxthc.Error(w, http.StatusNotFound, err)
			return
		}
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	log.Debug().Msg("Calling CheckPassword")
	err = postgres.CheckPassword(req.Password, u.HashedPassword)
	if err != nil {
		pxthc.Error(w, http.StatusUnauthorized, err)
		return
	}

	log.Debug().Msg("Calling CreateToken for Accesstoken")
	accessToken, accessPayload, err := s.tokenMaker.CreateToken(u.Username, s.cfg.AccessTokenDuration)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	log.Debug().Msg("Calling CreateToken for Refreshtoken")
	refreshToken, refreshPayload, err := s.tokenMaker.CreateToken(u.Username, s.cfg.RefreshTokenDuration)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	log.Debug().Msg("Calling CreateSession")
	session, err := s.store.CreateSession(r.Context(), &sqlc.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     u.Username,
		RefreshToken: refreshToken,
		UserAgent:    r.UserAgent(),
		ClientIp:     r.RemoteAddr,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	log.Debug().Msg("Setting LoginUserResponse")
	rsp := pxthc.LoginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  u,
	}

	log.Debug().Msgf("rsp: %+v", rsp)

	rspMarshalled, err := json.Marshal(rsp)
	if err != nil {
		pxthc.Error(w, http.StatusInternalServerError, err)
		return
	}

	log.Debug().Msg("Setting Content-Type header and writing response")
	w.Header().Set("Content-Type", "application/json")
	Json(w, http.StatusOK, rspMarshalled)

}
