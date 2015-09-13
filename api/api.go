package api

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

type API struct {
	Logger *log.Logger
	Config map[string]string
}

func New() *API {
	return &API{
		Logger: log.New(os.Stdout, "", 3),
		Config: make(map[string]string),
	}
}

func (a *API) Listen() {
	a.validateConfig()
	router := httprouter.New()

	router.GET("/v3/releases/:module", a.getRelease)

	log.Fatal(http.ListenAndServe(a.Binding(), router))
}

// Binding returns the configured binding
// if set, or logs a message and returns
// :8080 otherwise.
func (a *API) Binding() string {
	if b := a.Config["binding"]; b == "" {
		a.Logger.Print("No binding has been configured, defaulting to :8080")
		return ":8080"
	} else {
		return b
	}
}

func (a *API) validateConfig() error {
	if a.Config["modulepath"] == "" {
		return errors.New("Modulepath must be set before starting the API server")
	}
	return nil
}
