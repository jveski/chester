package api

import (
	"testing"
)

func TestValidateConfig(t *testing.T) {

	if New().validateConfig().Error() != "Modulepath must be set before starting the API server" {
		t.Error("validateConfig should return an error when the modulepath hasn't been set")
	}

	subject := New()
	subject.Config["modulepath"] = "stub modulepath"
	if subject.validateConfig() != nil {
		t.Error("validateConfig should return nil when the modulepath has been set")
	}
}
