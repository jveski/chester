package metadata

import (
	"testing"
)

func TestFromTarball(t *testing.T) {
	subject, _ := FromTarball("../test_fixtures/module.tar.gz")

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
