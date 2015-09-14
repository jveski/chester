package release

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jolshevski/chester/metadata"
	"io/ioutil"
	"strings"
)

// Release represents a specific Puppet
// module release on disk.
type Release struct {
	localPath string
	Metadata  metadata.Metadata
	File_uri  string
	File_md5  string
}

// New instantiates a new release object
// given the path to the tarball on disk.
func New(path string) *Release {
	return &Release{localPath: path}
}

// Tarball returns the path to the
// release's tarball on disk.
func (r *Release) Tarball() string {
	return r.localPath
}

// FromDisk populates the applicable
// properties given the tarball on
// disk.
func (r *Release) FromDisk() (err error) {
	// Get the metadata object from metadata.json
	m, err := metadata.FromRelease(r)
	r.Metadata = *m

	// Get the tarball's MD5 checksum
	raw, _ := ioutil.ReadFile(r.localPath)
	checker := md5.New()
	checker.Write(raw)
	r.File_md5 = hex.EncodeToString(checker.Sum(nil))

	return
}

// Slug will return the author-module-version
// formatted slug. Any slashes will be replaced
// with dashes.
func (r *Release) Slug() string {
	name := strings.Replace(r.Metadata.Name, "/", "-", -1)
	return name + "-" + r.Metadata.Version
}
