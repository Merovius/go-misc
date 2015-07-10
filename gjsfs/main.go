// Package gjsfs implements an http.FileSystem for gopherjs compiled files
package gjsfs // import "merovius.de/go-misc/gjsfs"

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/shurcooL/gopherjslib"
)

// New returns a http.FileSystem that wraps fs. All .js files opened are
// rewritten to .go names and - if existent in fs - compiled when read. All
// other files are passed through verbatim.
func New(fs http.FileSystem) http.FileSystem {
	return fileSystem{fs}
}

type file struct {
	r    io.ReadSeeker
	size int64
	f    http.File
}

func (f *file) compile() error {
	if f.r != nil {
		return nil
	}

	buf := new(bytes.Buffer)
	err := gopherjslib.Build(f.f, buf, &gopherjslib.Options{Minify: true})
	if err != nil {
		return err
	}
	f.r = bytes.NewReader(buf.Bytes())
	f.size = int64(len(buf.Bytes()))
	return nil
}

func (f *file) Read(buf []byte) (n int, err error) {
	if f.r == nil {
		if err := f.compile(); err != nil {
			return 0, err
		}
	}
	return f.r.Read(buf)
}

func (f *file) Close() error {
	if c, ok := f.r.(io.Closer); ok {
		return c.Close()
	}
	return nil
}

func (f *file) Readdir(count int) ([]os.FileInfo, error) {
	return f.f.Readdir(count)
}

func (f *file) Seek(offset int64, whence int) (int64, error) {
	if f.r == nil {
		if err := f.compile(); err != nil {
			return 0, err
		}
	}
	return f.r.Seek(offset, whence)
}

type rewriteInfo struct {
	os.FileInfo
	size func() int64
}

func (i rewriteInfo) Name() string {
	n := i.FileInfo.Name()
	return n[:len(n)-2] + "js"
}

func (i rewriteInfo) Size() int64 {
	return i.size()
}

func (f *file) Stat() (os.FileInfo, error) {
	i, err := f.f.Stat()
	if err != nil {
		return nil, err
	}
	return rewriteInfo{
		FileInfo: i,
		size: func() int64 {
			// TODO
			_ = f.compile()
			return f.size
		},
	}, nil
}

type fileSystem struct {
	fs http.FileSystem
}

func (fs fileSystem) Open(name string) (http.File, error) {
	if path.Ext(name) != ".js" {
		return fs.fs.Open(name)
	}
	name = name[:len(name)-2] + "go"

	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return f, nil
	}

	if fi.IsDir() {
		return f, nil
	}

	return &file{f: f}, nil
}
