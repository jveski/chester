package metadata

import (
	"testing"
)

type stubModule struct{}

func (s *stubModule) Tarball() string {
	return "../test_fixtures/module-1.2.3.tar.gz"
}

func TestFromModule(t *testing.T) {
	module := &stubModule{}
	subject, _ := FromModule(module)

	expectation := &Metadata{
		Name:    "stubuser-stubmodule",
		Version: "1.2.3",
		Author:  "stubuser",
		Summary: "stub summary",
		License: "stub license",
		Source:  "stub source",
	}

	if expectation.Name != subject.Name {
		t.Errorf("Expected FromTarball() to load the name from metadata.json")
	}

	if expectation.Version != subject.Version {
		t.Errorf("Expected FromTarball() to load the version from metadata.json")
	}

	if expectation.Author != subject.Author {
		t.Errorf("Expected FromTarball() to load the author from metadata.json")
	}
}
