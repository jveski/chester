package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

// API is a simple API server for the module
// types and factories.
type API struct {
	Logger *log.Logger
	Config map[string]string
}

// Error models a forge error response object.
type Error struct {
	Errors []string
}

// New returns a new instance of the API.
func New() *API {
	return &API{
		Logger: log.New(os.Stdout, "", 3),
		Config: make(map[string]string),
	}
}

// Listen starts the API server on the configured port.
func (a *API) Listen() {
	a.validateConfig()

	http.HandleFunc("/v3/releases", a.getReleases)

	log.Fatal(http.ListenAndServe(a.Config["binding"], nil))
}

func (a *API) validateConfig() error {
	if a.Config["modulepath"] == "" {
		return errors.New("Modulepath must be set before starting the API server")
	}
	return nil
}

func (a *API) returnError(msg string, w http.ResponseWriter) {
	err := &Error{Errors: []string{msg}}
	body, _ := json.Marshal(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	w.Write(body)
}
