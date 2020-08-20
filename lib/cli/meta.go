package cli

import (
	"io"
	"os"

	"github.com/spf13/afero"
)

// Meta defination
type Meta struct {
	// InReader define
	InReader io.Reader
	// OutWriter define
	OutWriter io.Writer
	// ErrWriter define
	ErrWriter io.Writer
	// FS define
	FS afero.Fs
}

// DefaultMeta define
func DefaultMeta() Meta {
	return Meta{
		InReader:  os.Stdin,
		OutWriter: os.Stdout,
		ErrWriter: os.Stderr,
		FS:        afero.NewOsFs(),
	}
}
