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

func TestFromDisk(t *testing.T) {
	subject := New("../test_fixtures/module-1.2.3.tar.gz")
	subject.FromDisk()

	if subject.File_md5 != "70ec46a9b9eb0a2d4983ec4ef834b14f" {
		t.Errorf("Expected FromDisk to return the correct tarball checksum")
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
