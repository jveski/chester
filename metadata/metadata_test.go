package metadata

import (
	"testing"
)

type stubRelease struct{}

func (s *stubRelease) Tarball() string {
	return "../test_fixtures/stub-module-1.2.3.tar.gz"
}

func TestFromRelease(t *testing.T) {
	release := &stubRelease{}
	subject, _ := FromRelease(release)

	if subject.Name != "stubuser-stubmodule" {
		t.Errorf("Expected FromTarball() to load the name from metadata.json")
	}

	if subject.Version != "1.2.3" {
		t.Errorf("Expected FromTarball() to load the version from metadata.json")
	}

	if subject.Dependencies[0].Name != "stub/dep1" {
		t.Errorf("Expected FromTarball() to load the correct module dependencies from metadata.json")
	}

	if subject.Dependencies[0].VersionRequirement != "stub_version1" {
		t.Errorf("Expected FromTarball() to load the correct module dependencies from metadata.json")
	}

	if subject.Dependencies[1].Name != "stub/dep2" {
		t.Errorf("Expected FromTarball() to load the correct module dependencies from metadata.json")
	}

	if subject.Dependencies[1].VersionRequirement != "stub_version2" {
		t.Errorf("Expected FromTarball() to load the correct module dependencies from metadata.json")
	}
}
