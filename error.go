package pxthc

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"net/http"
)

var (
	ErrBadRequest    = errors.New("error bad request")
	ErrInternalError = errors.New("error internal")

	ErrFormingResponse = errors.New("error forming response")

	ErrNoRecord = errors.New("no record found")

	ErrFetchingBook = errors.New("error fetching books")

	ErrDuplicateUsername = errors.New("error duplicate username")
)

func Errors(w http.ResponseWriter, statusCode int, errors []string) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(statusCode)

	if errors == nil {
		write(w, nil)
		return
	}

	p := map[string][]string{
		"message": errors,
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Err(err).Msg("at respond.Errors")
	}

	if string(data) == "null" {
		return
	}

	write(w, data)
}

func Error(w http.ResponseWriter, statusCode int, message error) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(statusCode)

	var p map[string]string
	if message == nil {
		write(w, nil)
		return
	}

	p = map[string]string{
		"message": message.Error(),
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Err(err).Msg("at respond.Error")
	}

	if string(data) == "null" {
		return
	}

	write(w, data)
}

func write(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		log.Err(err).Msg("at respond.write")
	}
}
