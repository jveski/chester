package release

import (
	"github.com/jolshevski/chester/metadata"
	"testing"
)

func TestNew(t *testing.T) {
	if New("stub path").localPath != "stub path" {
		t.Errorf("Expected New to set the path")
	}
}

func TestTarball(t *testing.T) {
	subject := New("../test_fixtures/stub-module-1.2.3.tar.gz")

	if subject.Tarball() != "../test_fixtures/stub-module-1.2.3.tar.gz" {
		t.Errorf("Expected Tarball to return the correct archive path")
	}
}

func TestFromDisk(t *testing.T) {
	subject := New("../test_fixtures/stub-module-1.2.3.tar.gz")
	subject.FromDisk()

	if subject.FileMd5 != "37a31eea4a43669c82cd216209cb395e" {
		t.Errorf("Expected FromDisk to return the correct tarball checksum")
	}

	if subject.Metadata.Name != "stubuser-stubmodule" {
		t.Errorf("Expected FromTarball() to load the name from metadata.json")
	}

	if subject.Metadata.Version != "1.2.3" {
		t.Errorf("Expected FromTarball() to load the version from metadata.json")
	}

	if subject.Metadata.Dependencies[0].Name != "stub/dep1" {
		t.Errorf("Expected FromTarball() to load the correct module dependencies from metadata.json")
	}

	if subject.Metadata.Dependencies[0].VersionRequirement != "stub_version1" {
		t.Errorf("Expected FromTarball() to load the correct module dependencies from metadata.json")
	}

	if subject.Metadata.Dependencies[1].Name != "stub/dep2" {
		t.Errorf("Expected FromTarball() to load the correct module dependencies from metadata.json")
	}

	if subject.Metadata.Dependencies[1].VersionRequirement != "stub_version2" {
		t.Errorf("Expected FromTarball() to load the correct module dependencies from metadata.json")
	}
}

func TestSlug(t *testing.T) {
	subject1 := New("")
	subject1.Metadata = metadata.Metadata{
		Name:    "stub-name",
		Version: "stub.version",
	}

	if res := subject1.Slug(); res != "stub-name-stub.version" {
		t.Errorf("Expected Slug to return 'stub-name-stub.version' not '%v'", res)
	}

	subject2 := New("")
	subject2.Metadata = metadata.Metadata{
		Name:    "stub/name",
		Version: "stub.version",
	}

	if res := subject2.Slug(); res != "stub-name-stub.version" {
		t.Errorf("Expected Slug to return 'stub-name-stub.version' not '%v'", res)
	}
}
