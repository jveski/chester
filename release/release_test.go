package release

import (
	"testing"
)

func TestNew(t *testing.T) {
	if New("stub path").LocalPath != "stub path" {
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
