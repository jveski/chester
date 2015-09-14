package api

import (
	"encoding/json"
	"github.com/jolshevski/chester/release"
	"net/http"
)

type releaseReponse struct {
	Pagination struct{}           `json:"pagination"`
	Results    []*release.Release `json:"results"`
}

func (a *API) getReleases(w http.ResponseWriter, r *http.Request) {
	response := &releaseReponse{}
	factory := release.NewFactory(a.Config["modulepath"], a.Config["fileurl"])

	if q := r.URL.Query()["module"]; q != nil {
		var err error

		a.Logger.Printf("Querying for all releases of %v", q[0])
		response.Results, err = factory.AllForModule(q[0])

		if err != nil {
			a.Logger.Printf("Error encountered while querying for all releases of %v: %v", q, err.Error())
			a.returnError(err.Error(), w)
			return
		}

		a.Logger.Printf("Returning %v releases for %v", len(response.Results), q[0])

		body, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	} else {
		a.returnError("Invalid query", w)
	}
}
