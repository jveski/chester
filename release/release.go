package release

import (
	"path/filepath"
)

// Release represents a specific Puppet
// module release on disk.
type Release struct {
	LocalPath string
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
