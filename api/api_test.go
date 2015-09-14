package api

import (
	"testing"
)

func TestValidateConfig(t *testing.T) {

	if New().validateConfig().Error() != "Modulepath must be set before starting the API server" {
		t.Error("validateConfig should return an error when the modulepath hasn't been set")
	}

	subject1 := New()
	subject1.Config["modulepath"] = "stub modulepath"
	subject1.Config["fileurl"] = "stub fileurl"
	if subject1.validateConfig() != nil {
		t.Error("validateConfig should return nil when the modulepath and fileurl have been set")
	}

	subject2 := New()
	subject2.Config["modulepath"] = "stub modulepath"
	if subject2.validateConfig().Error() != "Fileurl must be set before starting the API server" {
		t.Error("validateConfig should return an error when the fileurl hasn't been set")
	}

}
