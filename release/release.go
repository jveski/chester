package release

import (
	"github.com/jolshevski/chester/metadata"
	"path/filepath"
)

// Release represents a specific Puppet
// module release on disk.
type Release struct {
	LocalPath string
	Metadata  metadata.Metadata
}

// New instantiates a new release object
// given the module's name, version and
// the path containing module tarballs.
func New(q string, v string, modulepath string) *Release {
	path, _ := filepath.Glob(modulepath + "/" + q + "-" + v + ".tar.gz")

	return &Release{LocalPath: path[0]}
}

// Tarball returns the path to the
// release's tarball on disk.
func (r *Release) Tarball() string {
	return r.LocalPath
}

// FromDisk populates the applicable
// properties given the tarball on
// disk.
func (r *Release) FromDisk() (err error) {
	m, err := metadata.FromRelease(r)
	r.Metadata = *m
	return
}
