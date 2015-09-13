package api

import (
	"testing"
)

func TestParseModuleName(t *testing.T) {

	if slug, _ := parseModuleName("user-module-version"); slug != "user-module" {
		t.Error("parseModuleName should return the correct module name / author slug")
	}

	if _, ver := parseModuleName("user-module-version"); ver != "version" {
		t.Error("parseModuleName should return the correct module version")
	}

}
