package release

import (
	"testing"
)

func TestNew(t *testing.T) {
	if New("module", "1.2.3", "../test_fixtures").LocalPath != "../test_fixtures/module-1.2.3.tar.gz" {
		t.Errorf("Expected New to return the correct path to the tarball")
	}

	if New("module", "1.2.3", "../test_fixtures").Metadata.Name != "module" {
		t.Errorf("Expected New to return the given module name")
	}

	if New("module", "1.2.3", "../test_fixtures").Metadata.Version != "1.2.3" {
		t.Errorf("Expected New to return the given module version")
	}

	if New("doesn", "exist", "../test_fixtures") != nil {
		t.Errorf("Expected New to return nil for a non-existent module")
	}
}

func TestFromDisk(t *testing.T) {
	subject := New("module", "1.2.3", "../test_fixtures")
	subject.FromDisk()

	if subject.File_md5 != "70ec46a9b9eb0a2d4983ec4ef834b14f" {
		t.Errorf("Expected FromDisk to return the correct tarball checksum")
	}
}
