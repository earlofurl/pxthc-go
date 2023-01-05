package main

import (
	"github.com/earlofurl/pxthc/http"
	"github.com/rs/zerolog/log"
)

// TODO: Inject version dynamically
var Version = "v0.1.0"

func main() {
	log.Printf("Starting API version: %s\n", Version)
	s := http.NewServer()
	s.Init(Version)
	s.Run()
}
