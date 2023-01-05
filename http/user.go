package http

import (
	"database/sql"
	"github.com/earlofurl/pxthc"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("User Create handler called")
	var req pxthc.CreateUserRequest
	err := s.readJSON(w, r, &pxthc.CreateUserRequest{})
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
	err := s.readJSON(w, r, &req)
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
