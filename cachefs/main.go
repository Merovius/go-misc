// Package cachefs provides an in-memory cache for http.FileSystems
package cachefs

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"syscall"
	"time"
)

// Log is a log.Logger, where logs are sent.
var Log = log.New(os.Stderr, "[cache]", log.LstdFlags)

type inMemoryFile struct {
	*bytes.Reader
	fi os.FileInfo
}

func (f inMemoryFile) Close() error {
	return nil
}

func (f inMemoryFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, syscall.ENOTDIR
}

func (f inMemoryFile) Stat() (os.FileInfo, error) {
	return f.fi, nil
}

type entry struct {
	content []byte
	t       time.Time
}

type cache struct {
	mtx sync.Mutex
	m   map[string]entry
	fs  http.FileSystem
}

// Options contain configuration for the cache.
type Options struct {
}

// New returns a cache that wraps fs.
func New(fs http.FileSystem, o Options) http.FileSystem {
	return &cache{fs: fs}
}

// Open implements http.FileSystem.
func (c *cache) Open(name string) (http.File, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	f, err := c.fs.Open(name)
	if err != nil {
		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		Log.Println(err)
		return f, nil
	}

	if fi.IsDir() {
		return f, nil
	}

	e, ok := c.m[name]
	if !ok {
		return f, nil
	}

	if e.t.After(fi.ModTime()) {
		return inMemoryFile{bytes.NewReader(e.content), fi}, nil
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	c.m[name] = entry{content, fi.ModTime()}
	return inMemoryFile{bytes.NewReader(content), fi}, nil
}
