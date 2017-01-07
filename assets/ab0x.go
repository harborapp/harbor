package assets

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

// FileImagesFaviconIco is a file
var FileImagesFaviconIco = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x09\x54\x14\x57\xba\xbe\xdd\x8d\xa2\xd1\x08\x09\x1a\x41\x76\x69\xba\x1b\x7b\x61\x6f\x28\xb6\x06\x44\x94\xa5\x58\x0b\x44\x14\x45\x31\x46\x8c\x2b\x2a\x08\x34\xfb\xbe\xef\xfb\x26\x0a\xa2\xc6\x7d\x57\x10\x97\x18\x51\x93\x99\x68\xd6\x39\xc6\x18\x9d\x38\x2f\xc9\x9b\x89\x63\xcc\xe4\x1d\x27\x51\xbe\x77\xaa\x00\x43\x12\x31\xcc\xc4\x4c\x4e\xe6\xf0\x9d\xf3\x9d\xbf\xeb\xd6\x5f\xf7\x7e\xf7\xaf\x5b\xb7\xaa\xef\x7f\x09\xe1\x11\x01\xd1\xd6\x66\xad\x09\x59\xa9\x41\x88\x92\x10\x62\x62\x32\x70\xbc\x53\x9b\x90\x56\x0d\x42\xac\xac\x06\x8f\xcd\x09\xf1\x9e\x46\x88\x84\x10\xa2\xcd\xfa\x91\x81\x72\x0e\x1a\xe4\xbf\x1a\x13\xf8\x84\xb7\x25\x6e\xa1\x49\x5d\x65\x72\x50\x43\x61\xd4\xb2\xfc\x35\x26\x0b\xd5\x2b\x84\x6e\x4b\xc3\xec\x5e\x22\x5a\x19\xbc\xa7\x5d\x2b\x36\x17\x8d\xaf\x2a\x49\x88\x3d\x77\xba\xf3\xa3\x6b\x17\x1b\xbf\x3b\xd3\x24\xc5\x91\x02\x8d\xfe\x7d\x39\x9a\xff\xd8\x9e\xfa\xfc\xb5\x92\x0d\x33\x12\x97\xcf\xb7\x79\x69\xa4\xeb\x33\x57\x99\x2c\xd8\x5d\x2c\xfd\xfa\x74\xbb\x3b\x7a\xb7\xa9\x70\xbe\x69\x1a\x2e\x36\xf2\xf0\x7a\x1d\x0f\x67\x2a\x09\x8e\x14\x08\x1e\x35\x27\xbd\x70\x62\xc3\x92\x59\x66\xc3\xaf\x63\x68\x6a\x2a\xed\xe3\x3a\xab\x26\x5e\xb7\xe7\x58\xe9\x44\x9c\xa9\xd5\x44\x77\xcd\x54\x1c\xa9\xd4\x47\x6f\xed\x73\xe8\xa9\x99\x84\xe3\x15\x93\xd1\x53\xf9\x1c\x0e\xe6\x4f\x42\xc1\x5a\xa3\xfd\x01\x3e\x2e\xc2\x50\x9a\x9a\xc9\xd0\x94\x16\x43\x53\x7a\x61\x01\x8e\x73\x3b\x72\x8c\xae\x9d\xa8\x16\xf5\xf7\xb5\x99\x83\xe5\xa1\x72\x31\xde\x68\x11\xe1\x5c\x93\x18\x87\x2b\x24\xb8\xbc\x55\xf4\x6d\x5f\xab\xf9\x83\xa3\x15\x06\xdf\xe6\xac\x33\xcf\x0e\xf2\x75\xf6\xf2\xf6\x72\x63\x87\x08\x09\x0b\x50\x8e\x4b\x59\x25\x77\x28\x4d\x10\x9d\x6d\x48\x37\x43\x47\x81\x0c\xf5\x69\x42\x34\xa4\x8b\xd0\x9e\x27\x43\x63\x86\x29\x8a\xe3\xc5\x3b\xab\x93\x6c\x3c\x0f\x96\x89\xf7\x5c\x6a\x9d\x76\xff\x58\xc5\x0b\x97\x0a\xe2\x84\x21\x8f\x3b\xf2\xfc\x62\x7e\x62\xac\x64\x59\x6e\x9c\x79\x77\x7d\x86\xb0\xbf\x26\x55\x04\x96\xd5\x29\xe6\xdf\xe6\x6c\x10\xee\x4f\x78\xc5\x22\x88\x75\x3b\x56\x25\x6b\xec\x6b\x13\xde\xda\x9a\x39\x23\x3b\x7b\xad\xc8\xfb\xc7\x71\x4c\x78\xc5\xc2\xa7\x64\x8b\xf1\x8d\x2a\xb5\xf4\x66\x55\x8a\xc5\xbd\xf2\x24\xd3\xaf\xe3\x62\x14\x8e\x43\xe7\x1b\xd2\x2c\x57\x36\xa6\x5b\x46\x8f\x74\x1f\x96\x84\xdb\x3d\x17\x1b\x65\x65\xb8\x74\xbe\xbd\x7e\xde\x26\xab\xea\x86\x0c\xf9\xff\x6d\x59\x69\xfb\xf8\x7a\x2b\x1b\xd5\x38\x4b\x4b\x4f\xc1\xd3\xc6\xc2\x10\x92\x56\xd9\xda\xe6\x6e\xb4\x8e\x5a\xbd\xc4\x7e\xda\x68\xfc\x7f\x8f\x00\x7e\xc8\x7b\x2a\x42\x6e\x69\x11\x72\x56\x93\x90\xb3\x02\x42\x6e\xbd\x40\xc8\x3d\x6b\x42\x1e\xa4\x11\xd2\x9f\xf6\x53\x7f\xc9\xe0\x1c\xe3\x3e\x7c\x9e\xd1\xfe\xad\x7b\x35\x86\x31\xfc\xae\xc1\x27\x84\x8c\x23\xbc\x99\xe3\x08\x5f\xad\x41\x48\x33\xff\xd7\x6e\x70\x96\xd9\x8b\x93\x17\x84\xb8\xba\x25\xac\x5f\x90\x54\x90\x19\xbb\xbd\xae\x22\xf1\x50\x69\x5a\xc8\xd1\x8c\xd8\x99\x7b\xe3\x97\x89\x5b\xd7\x46\xc9\x92\xa2\x42\xed\x03\xbc\xbd\x3c\x4c\x0d\x24\x51\xe3\x9e\x61\xd3\x82\xc8\x50\x37\xaf\xf2\xc2\xb8\xa3\xdd\x47\x9b\xbe\xba\xf6\xd6\x41\x5c\x7f\xff\x04\xfe\x70\x7c\x23\x4e\x55\xe8\xe1\x58\x01\x1f\x87\xf3\x35\xb0\x2f\x47\x13\xdb\x53\xa7\xfc\xb3\x62\xa3\xde\xcd\x84\x18\xf1\xd6\x88\x20\x07\x6f\x85\x5d\xd0\xc4\x5f\xd2\xb0\x86\xc6\x0b\x1a\xab\x5f\x09\x7f\xf5\xe0\x9e\xea\xff\xfd\xe8\xc3\xd3\xf8\xcb\xed\x4b\xf8\xfc\xce\x9b\xf8\xe0\x6c\x02\x7a\x2b\x35\x71\xba\x9c\xa0\xbb\x8c\x87\x93\x25\x04\x27\x8b\x08\x4e\x16\x0f\xf0\x70\xfe\x38\x34\x27\xe9\xdc\x4f\x88\x11\x77\x04\xf8\xb8\xca\xff\xdd\xf6\x17\x06\x5b\x2f\xc8\xde\xe4\x79\x6f\x47\xcd\x22\x74\xd5\x2f\xc6\xbe\x6d\x1b\xd1\x73\x30\x17\x6f\xb4\x99\xe1\x52\xcb\x78\x5c\x6c\x12\x70\xbc\xd0\x20\xc0\xb9\x1a\x3e\x4e\x97\xf3\xb8\xf6\x4f\x14\x11\x9c\x2a\x26\x38\x98\x37\x1e\xc5\xeb\x0d\xae\x2f\x0c\x56\x06\x13\x7e\xfc\xbf\x34\x46\xe6\x78\x79\x18\x17\xac\x35\x7c\x87\x8d\x6f\x6f\x15\xc1\xb9\x3a\x1e\x7a\xab\xf9\x38\x5e\x3e\x01\x07\x6a\x9c\xb0\xaf\x3e\x00\x97\x5a\x26\x3c\xd6\x30\xc4\xd7\xeb\xf9\xe8\xad\xfc\x5e\xc7\x89\x22\x3e\xaa\x37\x4f\xff\x22\x2a\xd4\x3e\xe2\x5f\x69\xff\xe5\x08\xab\xd8\x5d\x99\x93\xfa\x5f\xaf\x21\xb8\xd4\xc2\xc7\xa5\xe6\x01\x5e\x6e\xe1\xe1\x7c\xab\x10\xe7\xb6\x39\xe0\x72\x8b\x06\x57\xd6\x37\x48\xce\x67\xd0\xf7\x8d\x06\x3e\x77\x7f\x58\x1d\xa7\x8a\x79\xa8\xda\xa4\xfb\xe9\x82\x20\x07\xd5\x48\xed\x31\x34\xa5\xc3\xd0\x94\x0d\x43\x53\xae\x01\x3e\x2e\x4e\xeb\x17\x4b\x0f\x6d\x4f\xd5\xc6\xee\x3c\x6d\xec\x29\xd0\xc6\xc1\x62\x2d\x1c\x2d\xd3\xc2\x91\x32\x2d\x1c\x28\xd2\xc2\x9e\x02\x2d\xae\x7c\x7f\x91\x36\x57\x76\xb8\x54\x0b\x7b\x0b\x07\x7c\x39\xff\x12\xd6\x77\x0a\xf6\xe5\x4d\x41\x47\xda\x14\xb4\x24\xbd\x88\x35\x51\xf2\x0b\xf4\x3c\x17\x15\x43\x53\xf6\x0c\x4d\xb9\x30\x34\xe5\xcc\xd0\x94\x84\xa1\x29\x4d\x86\xa6\x26\x33\x34\xe5\xc6\xd0\x54\x43\x80\x8f\xf3\x8a\xec\xb5\xe6\x57\x7b\xeb\x5e\x42\x67\x9e\x0c\xad\x99\x0a\x5c\xde\x2a\xc1\x7b\x3b\x45\xb8\xb6\x43\x84\x0b\xcd\x12\x34\xa5\x2b\x50\xb8\xc9\x0a\x67\x1a\x2c\xf0\x6e\x97\x08\x57\x3b\xc5\xd8\x5f\x2a\x45\xd6\x3a\x6b\xec\x29\x96\xe2\x4a\xbb\x18\xef\x75\x89\xf0\xce\x0e\x31\x7a\xea\x24\x28\x8b\x17\xa1\x39\xcd\x00\xaf\x44\xda\x34\x06\xf9\x3a\x39\x31\x34\x95\xc0\xd0\x54\x36\x43\x53\x42\x86\xa6\xc6\x0d\xc6\x80\x1f\x46\x53\x93\x8c\x2d\x02\x26\x17\x6d\x9c\x79\xe1\xad\x76\x1d\x7c\xb0\x5b\xcc\xf1\x43\xd6\xee\x1a\xfa\x2d\xc2\xfb\xbb\xc4\x78\xa7\x6b\xb0\x7c\x18\xaf\xee\x90\x7c\xef\x3f\x68\x3f\x39\x20\xc2\x5f\x8e\x18\xe2\xf6\x7e\x4d\x34\xa6\x18\x5e\x53\x3a\xce\xb1\x54\xd8\xcc\xd3\x51\x3a\xce\x9d\xf4\xa4\x7b\xe1\xe3\xed\x3a\x6e\x31\x63\x1f\xfc\x6a\x94\x4d\xec\xab\x51\x76\x9d\xc5\xf1\x66\xa8\x48\x36\x41\xe1\x66\x09\x9a\xb3\x14\x68\xc9\x56\xa0\x30\x7e\x16\xca\x12\x4d\x51\x92\x60\x86\xf2\x24\x39\x5a\xb2\xe5\xa8\x49\x55\xa0\x24\x41\x88\x6a\xb5\x09\xb2\x37\x88\xbf\x5e\xb9\x50\x99\xb3\x2e\xda\x3e\x36\x39\xd6\x36\xa5\x3a\xd9\xfa\xc8\xe9\x7a\xb3\xfb\x97\x5b\xb5\xfb\x77\xe4\x4c\xff\x4b\x6b\x86\xfe\x1f\x53\x56\x59\x54\xd3\xf3\x5c\x24\x4f\x1b\x87\x31\x11\x36\xaf\x54\x24\x1b\xa1\x3a\xc5\x00\x95\x6a\x21\xba\x8a\x65\xd8\x51\x24\x43\xa5\x5a\x84\xf2\x24\x43\x94\x25\x1a\xa1\x21\x63\x16\xb6\xe7\xcb\xd1\x96\x23\x43\x75\x0a\x7b\x3c\x03\x05\x9b\x4d\xff\x1a\xe2\xef\xa4\x18\xaa\x47\x6e\xa9\x32\xaa\x51\x5b\xbd\x73\xe7\xb0\x01\x6e\xee\xd5\xc0\xc7\x7b\x08\xae\x6e\xd7\x44\x79\xc2\xcc\x8b\x73\xbd\xdc\x84\x23\xb5\x1f\x1d\x6e\xb7\xbc\x70\xb3\xe9\x3f\xca\x93\x8c\xfe\x91\xbf\x49\xfc\x6d\x7b\x9e\x14\xdb\xf2\xa5\x28\x4e\xb0\x40\x75\x8a\x31\xca\x93\x4c\x50\xa9\x1e\x28\xab\x4f\x97\xa1\x30\x7e\xe6\x83\x2a\xb5\xe1\x37\x19\xeb\xcc\x6f\x87\xf8\x3b\x3d\x9e\x7b\x16\x04\x3b\x7a\xbd\xd1\x62\xf1\xcd\xdd\x53\x42\xec\x2b\x16\x7e\xb5\x2d\x4b\xf7\xc6\xee\x82\x97\x6e\xe6\xae\x17\x1d\x09\xf2\x75\xb6\x1d\xa9\xfd\xb0\x00\x4a\x3f\x8a\xb1\x77\x5f\x18\xa2\x54\xad\x8d\x56\xec\xaf\x4b\x33\x40\x5d\x9a\x3e\x4a\x13\x2d\xb8\x6f\xff\x96\x1c\x39\xca\x12\x85\x60\xcb\x4b\xb6\x98\x3c\x8a\x0e\xb7\xcb\x58\x12\x66\xe7\x1a\x19\xa2\x74\x0e\xa5\xa9\xc9\x43\xf5\x2c\x0d\x77\x08\x79\x6b\xbb\xe4\xbb\x4f\x0f\x9b\x23\x63\x8d\x6d\xb3\xcc\x7a\x8e\x91\xbb\xca\xc3\xd8\xc4\x82\x7e\xee\x69\xf1\x1f\x8e\x4d\x2f\x4b\x6b\x1b\xd2\x67\xa0\x2e\x6d\x06\xaa\x53\xc4\x68\xcb\x95\x73\xed\x57\xa9\x67\xa2\x4a\x6d\x80\xd2\x44\xa3\x47\xcb\x17\xd8\xcc\x7f\xd2\xb5\xa1\xfe\x94\xf2\x44\x8d\xf4\xef\x7b\x4b\x65\x7f\x0e\xa5\x29\xc7\xd1\xb6\x39\x1c\xb1\x51\x56\x59\x05\x9b\x67\xde\x2a\xd8\x6c\x7a\x23\x63\x9d\xec\x46\x55\x8a\xe5\x8d\xd2\x44\xab\xdb\x85\xf1\x92\x87\x0d\x19\x26\x28\xd9\x62\xfe\x68\x49\xd8\x93\xe7\x39\x37\x57\x97\x89\x6b\xa3\xed\x73\x62\x22\x1c\x98\x7f\xa7\x6d\x16\xf4\x3c\x97\x17\xe6\x78\xba\x1b\xb0\xf4\xf6\x74\x35\xf0\x74\x77\x99\xe1\xe3\xe5\xec\x5a\xb2\xc5\xea\x8b\x8e\x42\x19\x1a\x33\x15\x8f\x5e\x8e\x54\x8e\x38\xcf\x1a\x99\x7b\x8c\x7b\xd1\x60\xf6\x33\xfd\x56\x08\xf4\x71\x12\x97\x27\x59\x7e\xd6\x59\x28\x43\x73\x96\xfc\xd1\x8a\x48\xe5\x82\x67\x59\xff\xcf\x21\x60\x9e\x93\x7e\x62\xac\x6d\x5b\x71\x82\xd5\x81\x8c\x75\x36\xfb\x16\x85\x3a\xba\xfe\x27\xdb\x5f\x1c\xee\xc8\xd3\xd5\x9d\xa7\x41\x48\xa0\x86\x96\x96\x9f\xc6\xdc\xd9\x2e\xbf\xfa\xb7\xd8\x18\xc6\xf0\xdf\x06\x3c\x0d\xb7\x00\x10\x2d\xf4\x13\x01\x1e\x12\x1e\x1e\x10\xc2\x32\xed\x1e\x21\xaa\xab\x84\xe8\xdf\x23\xc4\xf8\x21\x11\x08\xfa\x89\x66\x06\x4b\xa4\x1b\x97\x21\x5d\xd5\x86\x74\xd5\x19\xf4\xab\xf0\x98\x4f\x01\xfb\x11\x60\xc5\xbe\x6f\x87\xaf\x53\x8c\xf8\x56\x1e\xc3\x18\xc6\x30\x86\x31\x8c\x61\x0c\xbf\x17\x4c\x10\x10\x81\x95\x85\xee\x34\x67\x5b\x53\x73\x27\x1b\x53\x39\x65\x6b\x6a\xed\xe6\xa2\xb4\xf6\xf0\x9c\x2d\xf5\xf4\x9c\x6d\xee\xa6\x9a\xa3\x6f\xae\x60\x26\xfc\xd6\x3a\x87\x43\x62\xa2\x2d\x70\xb4\x32\x12\x06\x78\xdb\xc6\xac\x5c\xea\xd7\x9a\x18\x17\xf9\x87\xe2\x9c\x55\xb7\x9a\x6b\x92\xbf\x68\xad\x4f\xbf\x5b\x99\xee\x7f\x37\xf7\x55\xa3\xcf\x52\x56\x98\xdd\x5a\xbf\x58\xfa\x7e\x6c\xa4\xe2\xf8\xc2\x10\xfb\x92\x00\x5f\x97\x25\x9e\x1e\xb3\x15\x4a\xa7\x79\x4f\x5c\x63\xf9\x4f\xc0\xc9\xd6\xd4\x2c\x22\xd8\x39\x37\x37\x7d\xe5\xf5\x43\x7b\xab\xbe\x7b\xfb\xca\x3e\x7c\xfc\xa7\x1e\x7c\xfa\xc9\x05\xdc\xfa\xf0\x30\xae\xec\x61\xd0\x53\x36\x85\x5b\xab\x3c\x55\xcc\xc3\xf1\x42\x01\x0e\xe5\x8d\x47\x57\xfa\xf3\xa8\x8b\x7f\xe9\xbb\xb4\x95\x66\x9f\xbd\x1c\x61\x7d\x34\xd0\xd7\x65\xa5\x9b\xbb\x97\xa1\x81\x24\xfc\xa9\x39\xda\x67\x05\x89\xe9\x8b\x02\x6f\x95\x34\x20\x69\xe3\xc2\xab\x3d\xc7\x9a\xfa\x6f\xdd\x38\x83\x3b\x9f\xbc\x8e\xdb\x37\xce\xe2\xd6\x8d\xf3\xb8\x7e\x75\x27\xce\x34\x28\x70\x2c\x9f\xe0\x78\x21\xc1\xf1\xa2\xa1\xb5\xce\x81\x75\x60\xb6\x3f\xdd\x25\x04\xa7\x4a\x78\xd8\x9f\x3b\x01\x35\x9b\x75\xbf\xdb\xb0\x44\xfa\x6e\x88\xbf\xd3\x26\x95\xfb\xec\x19\xbf\xa6\x76\xa9\x70\xba\x20\x2c\xc0\x69\x45\x75\x69\xfc\xdf\xd8\x58\x7f\xfe\x69\x1f\xee\xdc\xba\x80\xff\xb9\xdd\x87\xcf\xef\x5c\xc1\x67\xb7\xcf\xe3\xca\x4e\x17\x9c\xad\x22\x38\x5b\xc3\xc7\x99\xaa\x81\x75\xe3\xee\x52\x1e\x4e\x14\x7f\xdf\x8f\xc7\xfd\x19\xec\xcb\xb1\x42\x0d\xd4\x25\xbc\xf4\x30\x36\xd2\xb2\xcf\x77\xae\x9b\xbf\xab\x6a\xce\xb3\x5c\xe3\x7f\x0c\xdf\x39\xca\xd0\x95\x4b\xfd\xff\xd6\xd6\x90\x8a\x6d\x4d\x2c\xd3\xb0\xbb\xa3\x18\xdd\xc7\x3b\xf0\xde\xb5\xf3\xf8\xe0\x5c\x12\x2e\x36\x08\x70\xb1\x89\x70\xeb\xcd\x7d\xcd\x3f\x5c\x7b\x3e\x5b\xcd\x47\x4f\xd9\x60\x5f\x0a\x7f\xd8\x17\xb6\x1f\x07\x72\x27\x20\x23\xd6\xec\x6e\xa8\xbf\x53\xaa\x9b\xca\xeb\xf9\x67\xa9\xdd\x6b\xb6\xa7\x78\xfd\x12\xe9\xfb\xdb\x52\xb4\xb0\x5d\x3d\x01\xbb\xb3\x27\x60\x5f\xc1\x44\x74\x65\x4f\x42\x6b\xda\x54\xd4\xa8\xc5\xa8\x4e\x9a\x89\xbe\x76\x53\xf4\x75\xc8\xd1\xd7\xfc\xd3\xf5\xf3\xa1\xfe\x0c\x5f\x47\x3f\x5e\xf8\xc3\xfb\x71\xaa\x98\x8f\xf2\xb8\x19\xdf\x46\x04\x39\xd4\x7b\x78\xce\xd6\x79\x16\xda\x4d\x15\x0c\x7f\x7e\xa0\x63\xf1\x56\xf5\x0b\xe8\x29\x21\x38\x5b\xc5\x7b\xbc\x9e\x7e\xb9\x85\x8f\x2b\xad\xac\x25\xb8\xdc\x3a\x19\xdb\x2a\xa3\x90\x93\xbe\x06\x57\x3a\xa5\xdc\x1e\x86\x1f\xf7\xe1\x31\x9b\x05\x78\xa3\x41\x80\xde\x4a\xfe\xb0\x7c\xc0\xd0\xbd\xe0\xa1\x36\x7e\xfa\xa3\x88\x20\x87\x06\x95\xbb\xd7\x94\x5f\xaa\xdf\xc3\x63\xb6\x79\xfc\x32\xc9\xc7\x27\x8a\x04\xdc\x58\xbe\xd8\xf8\x64\x4d\x7d\xcd\xe3\xd1\xb3\xd5\x13\xbd\xaf\xc5\xe2\x72\xfb\x74\x5c\x6c\xe4\x8f\xac\x7f\x58\x3f\x2e\xd4\x0b\xb8\xe7\xe4\xc7\x7d\xa8\xdc\xa8\xf7\x6d\x28\x4d\x65\xfe\xd2\xe7\x21\xc0\xc7\x75\x11\x3b\x4f\xf4\x96\x11\x5c\xa8\xe7\x0d\xc6\xfb\xc9\x7c\xb3\x55\x03\x6f\xb6\x6a\x72\xb9\x90\x91\x7c\x7e\xcc\xa1\xfa\x2e\xd4\xf1\xd0\x53\x46\xd0\x3d\xf8\x5c\x9f\x2e\xe5\x21\x77\xb5\xc9\xbd\x40\x5f\x97\x51\xad\x4b\x32\x34\xa5\xcd\xd0\x14\xc5\xd0\x54\x00\x43\x53\xfe\x0c\x4d\xd1\xa1\xfe\x4e\x7e\xf3\x03\x1d\xba\xca\x36\xe8\xa3\x39\x49\x07\xdb\x32\xa7\xa2\x2b\x57\x07\x9d\xd9\x53\x39\xee\xcc\xd5\xc1\xae\x3c\x1d\xec\xcc\xd3\xc1\x8e\x1c\x1d\x74\x64\x4d\x45\x47\x96\x0e\x57\x3e\x74\xdc\x95\x33\xe8\x33\x78\x5d\xc7\xe0\x75\xec\xf9\xed\x59\x6c\x3d\x3a\x5c\x9d\xbb\xf3\x75\x38\xdf\x96\x64\x1d\x34\x24\xe8\x70\x79\x95\x9a\xcd\x7a\x88\x0c\xb6\x7f\x37\xc8\xcf\x39\x82\xd5\x33\xc8\xc7\xfa\x18\x9a\x9a\xc7\x0c\xec\x87\xe1\x33\x34\x35\x81\xa1\x29\x31\x43\x53\x2f\x33\x34\xf5\x01\x43\x53\xe7\x42\xfc\x9d\x12\x97\x86\xdb\xbd\xd9\x99\xa3\x8b\x5d\x05\xc6\x28\x8e\xb7\xc1\xb2\xf9\x0e\x48\x5f\x63\x8b\x03\x65\x32\x9c\x6f\x9a\x85\x2b\xed\x16\xb8\xd8\x6a\x81\xee\x5a\x29\xba\xf2\x15\xdc\xb9\xf9\x41\x14\x56\x44\x3a\xa0\x25\xd3\x12\xa7\xeb\xa5\xb8\xbc\xd5\x02\x17\x5a\x2c\x70\xb0\x4c\x86\xac\x75\xb6\x08\x0b\xa0\x10\xbb\x48\x89\x1a\xb5\x15\x0e\x57\xc8\x70\xae\x71\x16\xae\x6c\xb5\xc0\xa5\x36\x0b\xf4\xd4\x49\xd1\x9c\x21\xc3\xf2\x08\x1b\xc4\x2d\x93\xa3\x70\xa3\x39\x22\x02\x1d\x4f\x86\xfa\x3b\x85\x32\x34\x15\xcd\xd0\x54\x3d\x43\x53\x37\x19\x9a\x6a\x67\x68\x4a\xc5\xd0\xd4\x54\x86\xa6\x78\xc3\xee\xc3\x38\x86\xa6\xe6\x32\x34\x65\x30\x6e\x7a\xe4\xf8\xd5\x8b\x2d\x8f\xbc\xdb\x39\x1e\x5f\x1c\x9b\x82\xaf\x4e\x0b\x71\x73\xbf\x08\x7f\x3d\x29\xc4\xd7\xbd\x66\xf8\xea\xb4\x19\xbe\xea\x19\xe0\xfd\xd3\x66\x5c\xd9\xdd\x1e\x33\x5c\x69\x97\xe0\xc6\x3e\x11\x77\x7c\x7f\xc8\x67\xf0\xfc\x97\xdd\x66\xb8\xd8\x66\x81\x9b\x07\x44\xb8\xdf\x6b\xf6\xc4\x7a\xee\xf5\x98\xe1\xc3\xdd\xfa\xf8\xa0\x4b\x13\x37\xf6\xf0\x91\xb9\x46\xfc\xb1\x8b\xcb\x6c\x8b\x61\xfa\x9c\x19\x9a\x32\xfc\xb9\x31\xa5\x67\x1e\x32\x7e\x45\xa4\xcd\x91\x13\x95\x2f\xe2\x4a\x9b\x2e\xae\x75\x49\xb8\xdc\xd2\xdb\x9d\x12\xbc\xdd\x69\xc1\xf1\xea\x0e\x31\xfe\xd8\xf1\x3d\xaf\x76\x8a\xf1\x4e\x97\x04\xd7\x76\x48\x38\xbf\xab\x9c\xef\xf7\xe7\xd9\xdf\xef\xee\x94\x70\xf9\xa8\x81\x3a\x24\xb8\xd6\x25\xc6\xb5\x1d\x62\xae\xae\xf7\x76\x8a\x71\xfb\x90\x39\xee\x76\x0b\xf1\xc5\x71\x5d\x7c\xb2\x57\x80\xa3\xe5\x2f\x22\xd0\xc7\xb9\x54\xa2\xf0\x33\x90\x59\xfb\xe8\x98\x89\x03\x47\xf5\x4c\x07\xfa\x3a\x0b\xc2\x02\x1c\x5f\x59\x18\xa2\xac\x5f\xcc\x28\x6b\x16\x85\x38\xee\x5c\x13\x6d\xf5\x4d\xfe\x26\x21\x72\xe3\x06\x18\xbf\xc2\x12\x15\xc9\x96\x28\x4f\xb2\x44\xa5\xda\x12\x49\xab\xac\x91\xb9\x4e\xf4\xf8\x7c\xf2\x2a\x29\xf2\x37\x59\x73\x3e\x2c\xb3\x37\x58\x43\xbd\xca\x02\xf9\x9b\xcc\x90\xb7\x51\x88\xac\xf5\x22\x44\x87\x2b\x3f\x5c\x12\xe6\x50\xcb\x72\x79\x84\xb2\x79\xd3\x72\xbb\xe3\x25\xf1\x56\xb7\x4f\xd6\x4a\x1f\x7e\x72\x40\x0f\x1f\xee\xd4\x44\x79\x82\xe9\x37\x79\xeb\xcd\x6f\xa6\xae\x92\xbc\xb5\x3c\xc2\xb6\xcb\xd7\xdb\x35\xc2\xcd\xd5\x73\x54\xf3\xeb\x73\x13\xd5\xdc\xf8\x52\xb9\x7a\x5a\x26\xaf\x12\x7f\xd6\x92\xad\x87\x9a\x54\x03\x34\x66\xe8\xa1\x38\x41\x82\xc3\x35\xb3\xb0\xb3\x58\x86\x3d\x65\x52\x54\xa7\x48\x51\x93\xaa\x8f\x2a\xb5\xe1\x60\x6e\xcc\x84\xcb\xc7\xec\x2a\x91\x72\xf9\xb1\xf6\x3c\x19\x2a\x93\x8d\xd0\x96\xa3\xcb\xe5\x8b\x9a\xb2\xf4\x10\x1d\x6e\xdb\x46\x48\x22\x9f\x90\x28\x1e\x21\x0b\x78\x06\xa6\x5e\x13\x95\x8e\xae\x62\x3f\x6f\xa7\x94\xbc\x38\x9b\x6f\x6e\x1f\x9c\x8e\x8f\xf7\xf0\xf1\xd1\x6b\x3c\xfc\x69\x97\x00\x6f\x6e\x7d\x0e\x2d\xe9\x86\x0f\xa2\x42\x95\xad\xee\x2a\x8f\x51\xef\xb1\x9a\xed\xee\x61\x99\x18\x2b\xf9\xbc\x36\xd5\x00\x65\x89\xc6\xa8\x56\x1b\xa0\x60\xb3\x05\xa7\x7b\x6b\xae\x1c\x1d\x05\x72\x54\x24\xcb\x50\x96\x64\x88\xe2\x04\x13\x94\x24\x18\x73\xb9\xc4\x86\x0c\x05\xb6\xe6\xc9\xd0\x9a\x2d\x47\x73\x96\x1c\xa5\x5b\x4c\x50\x9f\xae\x8f\xf2\x24\x23\xd4\xa6\xe9\x63\xe9\x7c\x56\x7f\xce\x4f\xbe\x41\x83\xfd\xa8\x35\x07\xcb\xa5\x0f\xef\x9e\x12\xe2\xd3\x83\x53\xf1\xf1\x1e\x1e\x6e\xee\x25\x1c\xff\x7c\x80\xe0\x40\xc9\xb4\xfe\xc8\x60\x87\x22\x07\x47\xef\xf1\xa3\xd1\xef\xe9\xee\x61\x95\x14\x2b\xf9\x9c\x8d\x5d\x43\xc6\x0c\xb4\x64\xeb\xa2\x30\xde\x02\x87\x6a\xa4\xd8\x5d\x2a\xc3\xbe\x0a\x19\xa7\xbf\x3a\x65\x06\x6a\x52\xf4\x51\x9b\x3a\x03\xe5\x89\x6c\xfc\x15\x78\xad\x4c\x86\x5d\x25\x32\xee\x5e\xb0\xf1\x6f\xcf\x9b\x8e\xa6\x4c\x3d\x34\x73\xf1\xb7\x6b\x27\x44\xfd\x03\xfd\x86\x22\x8f\x49\x71\x31\x76\x3d\xec\xb3\xf0\xe5\x49\x21\xee\x1c\x12\x62\x7b\xb6\x11\x8e\x96\xeb\xe0\x7c\xc3\x64\x74\xd7\x68\xa3\x25\xcd\xf0\x9f\x8b\x19\xfb\xbd\x3e\xde\x6e\x23\xee\x6f\x1c\x8e\x40\x1f\x67\x61\x74\xb8\xdd\x9e\xb8\x18\x59\xcf\x86\x18\x59\xf7\xba\xa5\xf2\x73\xab\x16\x5b\xdf\xaf\x4a\x31\x47\x95\x5a\x88\xea\x54\x21\x12\x63\x15\x28\x4b\x92\xa2\x36\x4d\x82\xba\x74\x09\x32\xd6\x29\x90\xbf\x49\x82\xea\x14\x21\xc7\xac\x0d\xb3\xb0\x22\xd2\xe6\xa3\xb8\x18\xd9\x29\xb6\x8e\x35\x4b\x14\xbd\xf3\x03\x1d\x13\xed\xed\xe7\xfd\x40\xbf\xdc\x5a\xa5\x97\xb9\xd6\xe6\xfd\x2f\x4e\x08\xf1\xe5\x29\x21\x3e\xd8\x2d\xc2\x62\x86\x3a\x16\xec\x47\xad\x8f\x0e\xb3\xcb\x5f\x10\xe4\xb0\x65\xae\x97\x9b\xbf\xdf\x5c\xd7\xa9\xa3\xd1\xce\x62\x61\x88\x92\x3f\xc7\x43\x35\x49\x61\xed\x3b\x59\x66\xe9\x37\xc9\xc1\xd1\xdb\x74\xfd\x32\xd9\xd5\x8e\x82\x69\x68\xe4\x62\x39\x1d\x25\x5b\x84\xd8\x51\x34\x30\x96\x3a\x8b\x64\x68\xce\x66\xc7\xbb\x01\x17\xeb\x86\x0c\x3d\x54\xaa\x0d\x11\x1e\xe8\x98\x3e\xcd\x38\x6c\xa2\xa5\x8d\xef\x64\x3b\xfb\xb9\x93\x83\x7c\x9d\x7f\xf2\xdf\xf2\xb1\xfe\xe3\x42\x6e\x3e\x3d\x50\x2e\xfd\xce\x6f\xae\xd3\x88\xfb\x17\xff\x1d\xb8\xbb\x79\x4e\x5f\xbf\x4c\xf6\x36\x3b\x8e\x2a\x93\x07\x9e\xd7\xc2\x78\x73\x2e\x8f\x5e\x9f\xae\x40\x43\xa6\x82\xcb\x1d\x17\x27\x18\x73\xba\xd9\xf1\x5e\xba\xc5\x18\x91\x21\xca\xe4\x9f\xab\xdb\xc0\xdc\x73\xd2\xc6\x18\xbb\x9e\x3b\x47\xcc\xf1\xfe\x6b\x22\xac\x59\x62\x7f\xc6\xc1\x71\x74\xe3\x64\xb4\x50\xb9\x7a\x4e\x8f\x8b\x91\xbd\xdd\x59\x30\x0d\xcd\x99\x7a\x68\xcd\xd6\x45\xe9\x16\x21\xba\x8a\x64\xe8\x28\x1c\xc8\xc7\xb7\xe4\x48\x51\x91\x64\xc8\xc5\xbe\x2e\x7d\x06\x2a\x92\x0d\xb1\x20\xd8\xe1\x67\xf5\x93\x81\x1c\xf7\xc6\xce\x7c\x39\xd2\x56\xdb\xfc\xc9\x43\xe5\xe2\xf4\x2c\xb5\xb3\xf0\xf3\x76\x9d\x1a\x1d\x6e\x77\x28\xe5\x55\xf1\x0d\xf5\xab\xe2\xeb\x29\xab\xc5\xd7\x57\x2f\xb1\xba\x9e\xbe\xd6\xe6\x7a\x4e\x9c\x35\xc7\x8d\x2f\xdb\xfe\xbd\x70\xb3\x04\xb5\xa9\x42\xd4\xa6\x09\x91\xbd\x61\x16\x18\x9a\x1a\x95\xfe\x40\x1f\x27\xdd\x10\x3f\xaa\xc4\xd7\xdb\x69\xc4\xfd\x2d\xbf\x04\x8b\x42\x94\x7c\x5f\x6f\xd7\xe9\x76\xf6\x73\xf5\x59\xda\xda\xcd\xd3\x77\x77\x73\xd7\xb7\xb6\x75\xd3\x97\x5b\xa9\xf4\x4d\x44\x1e\x7a\x8b\x42\x1d\x1a\xd9\xb9\x75\x77\x89\x8c\x7b\x47\xb4\xe6\x28\x10\x19\xec\x38\x2a\xfd\x84\x8b\x91\xf3\xb8\xb9\x5e\x2e\xff\x91\xff\xf7\x4f\xc2\xf2\x08\x65\x79\x43\x86\x9c\x7b\x6f\xb1\xf3\x7f\x53\xa6\x1c\x8b\x42\x46\xaf\xff\xb7\x46\x4c\x84\xb2\x62\x6b\xae\x6c\xf0\xfd\x3b\xf0\x9e\x5b\x18\xe2\xa8\xfe\xad\x75\x8d\x16\x0b\x82\x1d\x73\xf2\x37\x5b\x7f\x59\xa9\xb6\xfc\x9c\x65\xf6\x06\xeb\xbf\x86\x05\x50\x71\xbf\xb5\xae\xd1\x22\xc4\xcf\x49\x6f\xb6\xbb\x8b\xdc\xdd\xcd\x45\xc6\xd2\xcb\xc3\x45\x1e\x4a\x53\xcf\x74\x1e\x1c\xc3\x18\xc6\x30\x86\x31\x8c\x61\x0c\xbf\x0d\x06\x76\x0b\xfd\x0a\xb6\x9f\xa8\x38\x7b\x8f\x68\x72\xf6\x16\xe1\x71\xf6\x2c\x21\x69\xac\x4d\x27\x44\xc5\x5a\x42\x88\x16\x08\xe9\x27\x84\x68\x0e\x5a\xc1\x30\xfb\x90\x10\xc2\x03\x21\x77\x58\xb1\x7d\x03\xc7\x64\xb0\x9c\xa4\x0d\xf8\x11\xad\x81\x7a\x1f\x0a\x06\x2d\x6f\xc0\x3e\x18\xb0\xbc\x07\x96\x84\xe0\x21\xd1\x7a\x90\xc6\x59\xe3\x21\xfb\x10\x03\xf6\x11\xa7\xfb\xa1\x31\x9e\x6a\x7f\x9d\x38\xfd\x7f\x00\x00\x00\xff\xff\x3a\x44\xf1\x3f\xee\x3a\x00\x00")

// FileScriptsApplicationJs is a file
var FileScriptsApplicationJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

// FileStylesApplicationCSS is a file
var FileStylesApplicationCSS = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func init() {
	if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}

	var err error

	err = FS.Mkdir(CTX, "images/", 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = FS.Mkdir(CTX, "scripts/", 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = FS.Mkdir(CTX, "styles/", 0777)
	if err != nil {
		log.Fatal(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileImagesFaviconIco)
	r, err = gzip.NewReader(rb)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "./images/favicon.ico", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileScriptsApplicationJs)
	r, err = gzip.NewReader(rb)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "./scripts/application.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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

	rb = bytes.NewReader(FileStylesApplicationCSS)
	r, err = gzip.NewReader(rb)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "./styles/application.css", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
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
	"./images/favicon.ico",
	"./scripts/application.js",
	"./styles/application.css",
}
