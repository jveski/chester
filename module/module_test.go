package module

import (
	"testing"
)

func TestNew(t *testing.T) {
	subject := New("module", "../test_fixtures")

	if subject.Path != "../test_fixtures/module-1.2.3.tar.gz" {
		t.Errorf("Expected FromQuery to return the correct path to the tarball")
	}
}
