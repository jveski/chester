package module

import (
	"testing"
)

func TestNew(t *testing.T) {
	subject := New("module", "../test_fixtures")

	if subject.Path != "../test_fixtures/module.tar.gz" {
		t.Errorf("Expected FromQuery to return the correct path to the tarball")
	}
}
