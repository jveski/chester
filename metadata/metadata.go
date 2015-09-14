package metadata

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Metadata is the parent object contained within
// a Puppet module's metadata.json.
type Metadata struct {
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	Dependencies []Dependency `json:"dependencies"`
}

// Dependency represents an entry under the
// dependencies key of the Metadata type.
type Dependency struct {
	Name                string `json:"name"`
	Version_requirement string `json:"version_requirement"`
}

type release interface {
	Tarball() string
}

// FromRelease takes a release object and returns
// an instance of Metadata containing the values
// found in the module's metadata.json.
// An error will be returned if an issue is encountered
// while loading the file, but an empty Metadata
// object will be returned if no metadata.json
// was found.
func FromRelease(r release) (*Metadata, error) {
	path := r.Tarball()
	rawFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	tarFile, err := gzip.NewReader(rawFile)

	if err != nil {
		return nil, err
	}

	file := tar.NewReader(tarFile)

	for {
		header, err := file.Next()

		if err == io.EOF {
			break
		}

		if strings.HasSuffix(header.Name, "metadata.json") {
			contents, _ := ioutil.ReadAll(file)

			metadata := &Metadata{}
			err := json.Unmarshal(contents, metadata)

			return metadata, nil

			if err == io.EOF {
				return nil, err
			}
		}
	}

	return nil, nil
}
