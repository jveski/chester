package main

import (
	"github.com/jolshevski/chester/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNoParams(t *testing.T) {
	app := api.New()

	app.Config["modulepath"] = "test_fixtures"
	app.Config["fileurl"] = "/stub/filepaath"

	req, _ := http.NewRequest("GET", "/v3/releases", nil)
	theServer := httptest.NewRecorder()
	app.GetReleases(theServer, req)

	if theServer.Body.String() != "{\"errors\":[\"Invalid query\"]}" {
		t.Fatal("API.GetReleases returned ", theServer.Body, " wanted {\"errors\":[\"Invalid query\"]} when given no params")
	}

	if theServer.Code != 500 {
		t.Fatal("API.GetReleases returned status ", string(theServer.Code), " wanted HTTP 500")
	}
}

func TestNonExistentModule(t *testing.T) {
	app := api.New()

	app.Config["modulepath"] = "test_fixtures"
	app.Config["fileurl"] = "/stub/filepaath"

	req, _ := http.NewRequest("GET", "/v3/releases?module=testuser-testmod", nil)
	theServer := httptest.NewRecorder()
	app.GetReleases(theServer, req)

	if theServer.Body.String() != "{\"pagination\":{},\"results\":[]}" {
		t.Fatal("API.GetReleases returned ", theServer.Body, " wanted an empty array")
	}

	if theServer.Code != 200 {
		t.Fatal("API.GetReleases returned status", theServer.Code, " wanted HTTP 200")
	}
}

func TestValidModule(t *testing.T) {
	app := api.New()

	app.Config["modulepath"] = "test_fixtures"
	app.Config["fileurl"] = "/stub/filepaath"

	req, _ := http.NewRequest("GET", "/v3/releases?module=stub-module", nil)
	theServer := httptest.NewRecorder()
	app.GetReleases(theServer, req)

	if theServer.Body.String() != "{\"pagination\":{},\"results\":[{\"metadata\":{\"name\":\"stubuser-stubmodule\",\"version\":\"1.2.3\",\"dependencies\":[{\"name\":\"stub/dep1\",\"version_requirement\":\"stub_version1\"},{\"name\":\"stub/dep2\",\"version_requirement\":\"stub_version2\"}]},\"file_uri\":\"/stub/filepaath/stubuser-stubmodule-1.2.3.tar.gz\",\"file_md5\":\"37a31eea4a43669c82cd216209cb395e\"}]}" {
		t.Fatal("API.GetReleases returned ", theServer.Body, " wanted an empty array")
	}

	if theServer.Code != 200 {
		t.Fatal("API.GetReleases returned status", theServer.Code, " wanted HTTP 200")
	}
}
