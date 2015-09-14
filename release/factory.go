package release

import (
	"path/filepath"
)

// Factory is responsible for instantiating slices
// of Releases based upon queries given from the API.
type Factory struct {
	modulepath string
	fileurl    string
}

// NewFactory returns a new instance of Factory
// with the given modulepath.
func NewFactory(modulepath string, fileurl string) *Factory {
	return &Factory{
		modulepath: modulepath,
		fileurl:    fileurl,
	}
}

// AllForModule returns an instance of Release for each
// available version of a given module. Each instance will
// have had .FromDisk() called on it already prior to returning.
// An error will be returned if an error is encountered during
// the process of loading each release from disk.
func (f *Factory) AllForModule(slug string) (releases []*Release, err error) {
	tarballs, err := filepath.Glob(f.modulepath + "/" + slug + "-*.tar.gz")

	if err != nil {
		return nil, err
	}

	for _, tarball := range tarballs {
		release := New(tarball)
		release.FromDisk()
		release.File_uri = f.fileurl + "/" + release.Slug() + ".tar.gz"
		releases = append(releases, release)
	}

	return
}
