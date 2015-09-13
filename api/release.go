package api

import (
	"encoding/json"
	"github.com/jolshevski/chester/release"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func (a *API) getRelease(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a.Logger.Printf("Request for %v received from %v", r.RequestURI, r.Host)

	// Parse the module's slug and version
	slug, ver := parseModuleName(ps.ByName("module"))

	// Instantiate the release
	result := release.New(slug, ver, a.Config["modulepath"])

	// Return 404 if the release was not found
	if result == nil {
		a.Logger.Printf("%v-%v failed to load from disk", slug, ver)
		http.NotFound(w, r)
		return
	}

	// If we've made it this far, load the release
	// from disk and return it as JSON
	result.FromDisk()
	a.Logger.Printf("%v-%v was successfully loaded from disk", slug, ver)

	body, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// parseModuleName takes a string with
// dash-delimited author, module, and
// versions, and returns a parsed user
// / module name slug, and version.
func parseModuleName(in string) (slug string, version string) {
	split := strings.Split(in, "-")

	slug = split[0] + "-" + split[1]
	version = split[2]

	return
}
