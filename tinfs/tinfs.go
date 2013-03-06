package tinfs

import (
	"path"
)

// Writer creates and builds a new TinFS, and encodes it to an io.Writer.
type Writer struct {
	all []*file
	n   int64
}

type file struct {
	tinpath   string
	localpath string
	size      int64
}

func NewWriter() *Writer {
	return &Writer{}
}

func (w *Writer) Add(tinPath, localPath string) error {
	?
}

func (w *Writer) Write(u io.Writer) error {
}

// Reader mounts a TinFS from an io.Reader in memory and serves the file system
type Reader struct {
	…
}
