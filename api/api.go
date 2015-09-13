package api

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

type API struct {
	Logger *log.Logger
}

func New() *API {
	return &API{
		Logger: log.New(os.Stdout, "", 3),
	}
}

func (a *API) Listen() {
	router := httprouter.New()

	router.GET("/v3/releases/:module", a.getRelease)

	log.Fatal(http.ListenAndServe(":8080", router))
}
