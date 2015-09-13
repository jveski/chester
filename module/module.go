package module

import (
	"path/filepath"
)

// Module represents a built Puppet
// module on disk.
type Module struct {
	Path string
}

// New instantiates a new module object
// given the module's name and the path
// containing module tarballs.
func New(q string, modulepath string) *Module {
	path, _ := filepath.Glob(modulepath + "/" + q + ".tar.gz")

	return &Module{Path: path[0]}
}

// Tarball returns the path to the module's
// tarball on disk.
func (m *Module) Tarball() string {
	return m.Path
}
