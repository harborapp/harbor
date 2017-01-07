package template

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileIndexHTML is a file
var FileIndexHTML = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\xc1\x4e\xc4\x20\x10\x40\xef\xfd\x8a\x91\x0f\x28\xd9\xfb\xb4\x89\x51\x0f\x9e\xdc\x18\x3d\x78\x9c\x85\xd9\x65\x94\x96\xa6\x43\x9a\x6c\x36\xfb\xef\xa6\x82\xd1\xed\x89\x09\x8f\x17\x1e\xe0\xdd\xe3\xcb\xc3\xdb\xc7\xfe\x09\x42\x1e\x62\xdf\x34\x58\x56\x00\x0c\x4c\x7e\x1d\x00\xf0\x40\xca\x10\x66\x3e\x76\xe6\x72\x81\xf6\x35\xa5\x0c\xd7\xab\x01\x5b\x79\x94\xf1\x0b\x66\x8e\x9d\x11\x97\x46\x53\x8f\x92\x2a\x67\xb5\x32\xd0\x89\xd5\x1e\x69\x59\x61\x2b\x2e\x99\xbe\xd9\x7a\x9a\xcf\x91\x35\x30\xe7\x8d\x5d\x80\xa5\x69\x8a\xe2\x28\x4b\x1a\x5b\xa7\x5a\xaf\x46\x5b\x22\xd7\xf1\x90\xfc\xb9\xe6\x78\x59\x40\x7c\x67\xfe\x49\xa6\xa0\xf5\x59\xbb\xfe\x7d\x50\x17\x22\x9d\xe0\x7e\xff\x8c\x36\xec\xaa\x66\xbd\x2c\xbf\x65\xea\x66\x99\x32\xe8\xec\xfe\x4a\x7e\xb6\x6e\x53\x3e\xd5\xf4\x58\x49\x09\x2a\x19\x68\xcb\x37\x7e\x07\x00\x00\xff\xff\x7b\xf0\x5a\x26\x5f\x01\x00\x00")

func init() {
	if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}

	var err error

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileIndexHTML)
	r, err = gzip.NewReader(rb)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "./index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}
}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// FileNames is a list of files included in this filebox
var FileNames = []string{
	"./index.html",
}
