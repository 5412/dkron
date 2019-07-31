// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Templates contains project templates.
var Templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 7, 10, 21, 17, 17, 865450392, time.UTC),
		},
		"/dashboard.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "dashboard.html.tmpl",
			modTime:          time.Date(2019, 7, 10, 21, 14, 58, 335961438, time.UTC),
			uncompressedSize: 3853,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\xcc\x12\x05\x36\x01\x2a\xf1\xd0\xdb\x56\x36\x10\x24\x0b\x74\xdb\x6e\x36\xd8\x06\x05\x8a\xa2\x08\x46\x22\x2d\x31\xa6\x48\x96\xa4\x5c\x1b\x86\xff\x7b\x41\xc9\x92\x25\x47\xde\xd8\x6e\x7a\x12\xbf\xe6\xcd\x9b\x79\xc3\x0f\x25\xef\xee\xbe\xdc\x3e\xfe\xf1\xf0\x11\x0a\x5f\xca\xd9\x24\x09\x1f\x50\x79\x84\xc6\x4c\x09\x5b\x58\xad\xc8\x6c\x32\x49\x0a\x8e\x6c\x36\x01\x48\xbc\xf0\x92\xcf\x36\x1b\x88\x6f\x75\x59\x6a\x15\xdf\x63\xc9\x61\xbb\x4d\x68\x33\x13\xd6\xb8\xcc\x0a\xe3\xc1\xd9\x6c\x4a\x7a\x2b\x6f\x9c\xe3\xde\x3d\xa0\x2f\x60\xbb\xa5\xcf\x8e\xa2\xca\x2b\x89\x36\x2e\x85\x8a\x9f\x1d\x99\x25\xb4\xb1\xbc\x04\x24\xb2\xba\xf2\x7c\x04\xea\x1c\x2c\x86\xae\x48\x35\x5a\xf6\x92\x8e\x14\x6a\x01\x96\xcb\x29\x71\x7e\x2d\xb9\x2b\x38\xf7\x04\x0a\xcb\xe7\xdf\x80\xcc\x9c\xa3\xb9\x5c\x9b\x42\x64\x5a\x45\xce\x08\xa5\xb8\x8d\x33\xe7\xc8\x79\xbc\x9e\xff\xae\xb8\x5d\x1f\x09\xee\x62\x66\xa9\xd6\xde\x79\x8b\xa6\x06\x6e\x58\x9d\x41\x6a\x68\xfe\x86\xbc\xd0\x98\x0b\x72\xc4\x7e\x18\xaf\xa3\x8b\x68\x58\x91\x2d\x5c\x81\xff\x74\x8d\x0b\x52\x34\x8e\x71\x59\x9d\x77\x45\xfe\x02\xf3\xad\xf2\xae\x19\x2f\x85\xb5\xda\x52\x29\xd2\x5e\xf7\xcc\xa0\x8f\xe3\x5c\x16\x78\x0f\xaf\xd4\x8c\xd3\x67\x5c\x62\x63\xd8\x6b\xfe\x77\x6c\x64\x4c\x2b\xca\x84\x33\x12\xd7\x14\x2b\xaf\x2d\x9f\x5b\xee\x8a\xcb\xcf\xa5\x4a\x44\x47\xa3\x3f\x90\xa9\xd0\xd6\x67\x95\x87\x70\x4a\xbc\xaa\xd4\x1c\x97\x61\x5d\x2c\x32\x4d\xc0\xaf\x0d\x9f\x12\x51\x62\xce\xe9\x2a\xaa\xed\x0f\x8a\xfe\x8d\x30\xf7\x91\x07\x7c\x80\x25\x5a\xb8\xfb\xe5\xeb\x97\xfb\xa7\x9b\x87\x4f\x4f\x0f\x37\x8f\x3f\xc1\x14\x06\x0e\x1e\x3e\xed\xd0\xc9\x8f\xb5\xc5\x77\x57\xf3\x4a\x65\x5e\x68\x05\x57\xd7\xb0\xa9\xc7\xc2\xe8\xfb\x3f\x19\x7a\x8c\xbc\xce\x73\xc9\xa7\xc4\x6b\x2d\xbd\x30\xe4\xaf\xf7\xd7\xf1\xae\x7d\x75\x5d\x2f\xde\x86\xcf\x3e\x87\x09\x6d\xee\xa4\x49\x92\x6a\xb6\xae\xa3\x66\x62\x09\x99\x44\xe7\xa6\x44\xe1\x32\x45\x0b\xcd\x27\x62\x7c\x8e\x95\xf4\x6d\x77\x2e\x56\x9c\x45\x5e\x1b\x02\x56\x07\xa7\x0a\x97\x22\xc7\xc0\x8d\x34\xe1\xf5\xa1\x32\xad\x3c\x0a\xc5\xed\x6e\x0e\x20\x79\x17\x45\x70\xab\xa5\x44\xe3\x38\x83\xbd\x35\x44\x51\xb7\xe6\x05\x99\x28\xf0\xed\xa1\xec\x70\x3e\xae\x0c\x2a\xc6\x2d\xa4\x95\xf7\x03\x08\x80\x64\x37\xd6\x48\xd2\x74\xc8\x01\x6a\x93\x38\x02\x83\x2c\x66\x3b\x72\xed\x30\xda\x9c\xfb\x29\x89\x77\x36\xdd\xf4\xde\x55\xd0\xd7\xa0\x6a\xc1\x9d\x8d\xb4\x92\x6b\x32\x7b\xac\x11\x7b\x31\x26\x34\xac\x3b\x6a\x58\xdf\x74\x29\xda\xba\xdc\xff\xe7\x85\x09\x6d\x52\x52\x97\x67\x2f\xa5\x9f\x51\x28\xa8\x9f\x23\xc3\x6c\xf6\x24\x49\x2d\x2a\x36\x0c\x1f\xdb\x39\xa9\x73\x3d\xb2\x65\xda\x72\x9e\x25\xa2\xcc\x5f\xd9\xff\xa2\xcc\x69\xfd\x78\x8a\x02\x58\x6c\x54\x4e\x80\xce\x12\x8a\x7d\xf2\x4c\x2c\xbb\x6a\x69\x3a\x93\x97\x75\x71\x7a\x79\xb5\xa2\xc2\x88\xba\xfb\xb4\x8c\xa2\x01\x24\x95\x6c\xe1\x4a\x14\x2a\x0a\x07\x88\x0b\x8b\xdb\x3d\xa3\x70\x39\x4c\x97\x14\xfd\x6e\x9d\xbf\xe3\x29\x1b\xac\x04\x48\x44\xeb\xac\x7b\x1d\xc1\xfe\x9d\xd4\xbd\xc2\x82\xe2\xe2\xd0\xf6\xae\x9d\x1d\x7a\xef\xa7\x36\x74\x87\xf4\xce\x61\x4b\x9f\x75\xea\xce\xa4\xec\xd1\x2d\xdc\x28\xdd\x9f\x75\xea\x4e\x67\x9a\xd0\x4a\x1e\xa8\xf6\x1b\x47\x9b\x15\xdf\xc3\xfd\x5e\x38\x54\x0c\xbe\x72\xa3\xa1\x51\xe9\x88\x8e\x43\xf1\xda\xa6\x15\x79\xe1\x4f\x54\xb2\xf0\xde\x7c\xa0\x4d\x21\xc7\x22\x5c\x0d\xbb\x73\xe4\x29\x95\xa8\x16\x67\xe6\x28\xd5\x7a\x31\xae\xa8\xce\x2e\x4c\x51\x6f\x0f\x75\xcd\xae\xb1\xd9\x78\x5e\x1a\x89\x9e\x43\x7d\x86\x73\xe5\x09\xc4\xdb\xed\x37\x6f\x8a\xe6\x6a\x48\xb5\xf7\xba\x6c\xc7\x84\x5a\x72\xdb\x6d\xa7\x91\x9d\x57\x3f\xeb\xc9\xd8\xd6\xec\xee\x0e\xf0\x7c\xe5\xa3\x8c\x2b\x3f\xbc\x01\x06\x72\xbd\x22\x8b\x39\x3c\xf9\xf9\xca\x93\x91\x5f\x31\xe8\x0d\xfd\xce\xad\x0b\x25\x13\x7e\xd0\xcc\x79\x3b\x64\xdc\xdd\x67\x5e\xa6\xdc\x7e\xe8\xfb\x68\x86\xba\xff\xc0\x37\x71\xf3\x6b\x7d\x57\x0e\xdc\x34\x43\xa7\xb8\x39\xb1\x44\x12\xda\xbc\x1c\xc2\x53\xa2\xfe\xf9\xfd\x37\x00\x00\xff\xff\xe7\x77\xe7\x5b\x0d\x0f\x00\x00"),
		},
		"/executions.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "executions.html.tmpl",
			modTime:          time.Date(2019, 7, 31, 16, 26, 28, 795099963, time.UTC),
			uncompressedSize: 2655,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x51\x6f\xe3\xb6\x0f\x7f\xef\xa7\x60\xf5\x0f\x1a\x07\xff\xda\x46\x0f\x7d\xba\xb3\x0d\xdc\x8a\x6d\x40\x71\xbb\x01\xeb\xb0\xd7\x41\xb6\x18\x4b\x57\x47\x32\x24\xfa\xd2\xc0\xe7\xef\x3e\x48\x8e\x13\xa7\x49\x71\x68\x1f\x5c\x89\xa4\xc8\x1f\x7f\xa4\xa8\xf4\xbd\xc0\xb5\xd2\x08\xac\x32\x9a\x50\x13\x1b\x86\xab\x4c\xa8\xef\x50\x35\xdc\xb9\x3c\x88\xb9\xd2\x68\x61\x32\x00\x5d\xc7\x7e\x6d\x4d\xd3\xa0\xcd\xd9\xaf\x2f\x58\x75\xa4\x8c\x76\x0f\x64\x1b\x56\x5c\x01\xcc\x1d\x58\xb3\x0d\xb2\xbd\x54\x89\x9c\x6d\xd0\x39\x5e\x23\x2b\xb2\x54\xa8\xef\xe1\xc0\x61\x71\xf9\xa4\xbc\x2b\x8e\x61\xb2\x54\xde\x4d\xf2\x0f\xc5\x93\x34\x5b\xa5\x6b\xf8\xc2\x1d\x41\xdf\x27\x0f\xa6\xd3\x34\x0c\x80\x07\x7b\x58\x1b\x0b\xdf\x4c\xe9\xb5\x8f\xa6\xfc\xca\x37\x38\x0c\xe1\x3c\x40\xc6\x41\x5a\x5c\xe7\xec\x7f\x63\x5e\x8d\xaa\x9e\x73\x66\x3b\xfd\x68\xca\x68\x39\x3f\xb0\x5c\x05\x0b\x27\xcd\x36\x67\xd7\xb6\xd3\x5a\xe9\xfa\xdf\x6f\xa6\x64\x40\x8a\x1a\xcc\xd9\x5f\x9d\x66\x45\xa6\x26\xf8\x75\xb3\x6b\xa5\xaa\x8c\x86\xc3\x2a\x6e\x1b\xbe\x63\xfb\xd0\xfe\x8f\x5b\xc5\x63\xa9\x84\x40\x9d\x33\xb2\x5d\xe0\x44\x15\x59\xca\x8b\x73\x80\x3f\xf1\x5d\x99\x1a\xb4\xb1\x1b\xde\xc4\x56\xd5\x92\x62\xd7\x2a\xad\xd1\xce\x60\xcf\x51\x9f\x06\xca\x52\xf9\x61\x56\x88\x2b\x80\xbe\xb7\x5c\xd7\x08\x8b\xfa\x19\x3e\xe6\x90\xfc\xb2\xfb\xdd\x9a\xae\x0d\xcc\x65\xf2\xfe\x58\x0f\x08\xf2\x8f\xd0\xf7\x8b\xfa\x79\x18\x20\xea\xfb\x06\x35\x44\x4a\x0b\x7c\x81\x45\x12\xd4\xce\xfb\x59\x0d\xc3\x2a\x4b\xe5\x7d\x08\x44\xbc\x6c\x70\xca\x67\xdc\x84\x6f\xec\xc8\xaa\x16\xc5\x54\x7a\xb2\x07\x2a\x48\x16\x8f\xa6\xcc\x52\x92\x73\xd1\x13\x71\x4b\x28\xe0\x33\xbd\xd6\xfc\xa6\xb4\x72\xf2\xa2\xea\xab\x11\xf8\x5a\xf6\x67\x47\x6d\x77\x66\xf9\xd4\x55\x15\x3a\x77\x14\x67\xe9\x04\xe9\x40\x11\x3e\xdf\xc2\x02\x5f\x3c\x4f\x97\xd3\x3e\x4f\x45\x14\x7d\xbf\xc0\x97\x63\x7f\x65\x29\x89\x73\xf5\x3e\xb9\xcf\xf4\x96\xc1\x94\xe3\xdb\x16\x3e\xd5\xcb\x21\x0e\x8d\x78\x6c\xb2\xbe\x87\x19\x2a\x18\x86\xd8\x4b\xea\xe7\x69\x85\x61\xb5\x31\x82\x37\x0c\x04\x27\x1e\x93\xa9\x6b\xdf\xfd\x27\x32\x6e\x6b\xa4\x77\xba\x2b\x66\xf7\x62\x44\x3e\x16\x04\x7e\x00\x99\x27\xb2\xfe\x92\xff\x00\xb2\x9d\xae\x38\xe1\x30\x24\x49\x32\xbf\x27\x97\xe9\x1b\x8b\x37\x4f\x7d\x2c\xdf\xb8\xbc\x8e\x63\xf8\xc3\x07\x87\x38\x9e\xcd\xa8\x7d\x53\x06\x58\xb0\xe6\x02\x59\x18\x5b\xef\xe2\x86\x78\x19\x3a\x21\x67\xf1\x1d\x03\x6b\x3c\x45\x42\xf1\xc6\xd4\xd3\xfd\x0f\x77\xbf\xe1\x25\x36\x0d\x8a\x72\x97\xb3\xcd\x2e\x60\xf9\xe2\x45\x07\x36\xce\x00\xc5\x7b\x2f\x93\x4f\x53\x75\x1b\x3f\x96\x67\xd5\x3c\x3b\x32\x8d\xee\x39\xc5\xe7\x56\x12\xb9\x40\x7b\x62\x04\x90\x95\x1d\x91\xd1\x40\xbb\x16\x73\x36\x6e\xd8\xe1\x71\x68\x8c\xc3\x7d\xcd\x85\x72\x1b\x75\x70\xc6\x66\xe9\xe5\xec\x21\xd8\x15\x99\x6b\xb9\xbe\x34\xf3\x6e\x48\x6d\xd0\x7d\xca\x52\x6f\x50\x64\xe9\x18\xe6\x15\x10\x79\x7f\x0a\x37\x4c\xdd\xb1\x32\xa7\xcc\xfd\xa3\x70\x0b\x66\xec\x1d\x3f\xfa\xcf\x2e\xda\x38\x81\x0e\x8e\xa7\xb7\xe7\x6d\x66\x4a\x23\x76\xaf\x79\x69\x2d\x16\x6f\xb4\xa9\x8f\xe1\xd5\xef\x0b\xb2\x36\x86\xde\x47\x7f\x49\x1a\x4a\xd2\xb1\xc0\x35\xef\x1a\xba\x5c\x88\x22\x90\x7f\x89\xd3\x57\x98\x4e\xb6\xb3\xcd\x6c\xd9\xf7\xa8\xc5\xf8\x02\xa4\x61\x58\x17\x57\x47\xe1\xde\x2e\x73\x95\x55\x2d\x79\xcd\x22\x9a\xba\x73\x95\x58\xe4\x62\x17\xad\x3b\x5d\x85\x37\x23\x5a\x41\x1f\x5c\xaa\x35\x44\x5b\xa5\x85\xd9\x26\x8d\xa9\xb8\x57\x26\x92\x3b\x09\x37\x37\xb0\xb8\xa8\x59\x25\x0d\xea\x9a\xe4\xe4\x01\xde\xb4\x0b\x04\x44\x4b\xff\xf6\x2d\x57\x9f\x82\xf1\x38\x89\x17\xd1\x72\x54\x2e\x57\x89\xd1\xa3\x45\x52\xba\xbd\xec\x16\x8e\x30\xf1\x18\xc5\x23\xf5\x65\x30\x6b\x88\x30\xb1\xd8\x70\x42\xf1\x77\x98\x74\x2b\xb8\xce\x81\x75\x7a\xfc\x2d\x25\xd8\xf1\x10\xc0\xc5\xe4\x72\x58\x9c\xf9\x48\x38\x91\x8d\x96\x7e\x10\x4f\x68\x27\xbc\xc3\x7e\x3f\x27\xd4\xe3\x1e\x6f\xd1\x4f\x91\x4b\xe5\xc8\xd8\x5d\xd2\x76\x4e\x3e\x11\x27\x8c\x18\xbb\x85\xc9\x55\x12\x2e\xd2\xed\x19\xd0\x96\x93\xd4\x7e\xd0\xfd\xff\x4c\xe5\x90\xdb\x4a\x4e\x94\x86\xff\xfe\x9b\xa5\x53\xed\xa7\x9e\xf8\x2f\x00\x00\xff\xff\xa4\x38\x14\xe2\x5f\x0a\x00\x00"),
		},
		"/index.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "index.html.tmpl",
			modTime:          time.Date(2019, 7, 4, 21, 24, 27, 840596557, time.UTC),
			uncompressedSize: 1090,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4d\x8f\xda\x30\x10\xbd\xf3\x2b\x46\xbe\xd0\xaa\x85\x48\xbb\xb7\x95\x13\xa9\xed\xa9\x52\xb5\xaa\xd4\xfe\x01\xc3\x0c\xc4\xc2\xd8\xd1\x78\x60\x8b\xd2\xfc\xf7\xca\x21\x21\xa6\xb0\xab\xbd\xc4\xcf\xef\x3d\x8f\x3d\x1f\x69\x5b\xa4\x8d\xf5\x04\x6a\x1d\xbc\x90\x17\xd5\x75\x33\x8d\xf6\x08\x6b\x67\x62\x2c\x7b\xda\x58\x4f\x0c\xa3\x01\xfc\x76\x91\x30\x07\xe7\x88\x4b\xf5\xdd\x23\xfd\xf9\x26\xec\x54\x35\x03\x68\x5b\xa1\x7d\xe3\x8c\x10\xa8\x28\x46\x0e\x51\xc1\xb2\xeb\x66\x00\x79\x54\x0e\x2f\xbd\xfb\x9a\x6d\x8c\x27\x07\xfd\x77\x81\xb4\x31\x07\x27\x83\xeb\x8e\x6f\x51\x93\x41\xeb\xb7\x17\x07\x80\xae\x1f\xaf\x2d\x62\xc5\x91\xaa\x7e\xd8\x23\xe9\xa2\x7e\xbc\xc4\x2a\xd0\x1e\x5f\x0f\xbc\x0a\x78\xca\xa3\xb2\x5d\xef\x62\x6d\x5e\x60\x04\x8b\xd0\x88\x0d\x3e\x96\x6a\x00\x6a\x92\x36\x64\xe4\xc0\x14\x4b\x35\xa2\x4c\x8c\xc4\x36\x49\xe7\x55\x55\xba\x18\xa5\x3b\x4f\xbb\xc0\x01\xbc\x5e\xc2\xfa\xa1\x7a\x0e\x48\x51\x17\xf5\xc3\x40\x89\x59\x39\x1a\xad\xe7\x4d\xff\x5d\x44\x61\xdb\x10\x4e\x75\x15\xce\x52\x95\xba\x7a\x36\x7b\xd2\x85\xd4\xd7\xec\x17\x44\xa6\x18\x6f\x85\x9f\x81\xe5\x96\xfd\xd5\x77\xfe\x96\xff\x6d\xb6\x57\xac\x2e\xa6\xeb\xb5\x70\x1a\x2d\xa6\x86\x8c\x94\x6a\x4f\xfb\x15\x31\x58\x0f\x67\x14\xe1\x2f\x04\x46\xe2\xaf\xa7\xa7\x79\x7a\xe4\x3c\xef\x91\x60\x3a\xbb\xb2\x1e\xc7\x93\xcb\xe4\x49\x25\x16\x7c\xdb\x97\x52\x7b\x8f\x2f\x65\xfa\x1e\xdf\x39\xf7\x3b\xce\x69\x33\x74\x32\xcb\xf6\xc3\x8e\x4e\x9f\xe1\x68\xdc\x81\x3e\x4e\x39\x2f\x53\xbd\x54\x7e\x0e\x40\xc7\xc6\xf8\xe9\xd2\x1d\x9d\xe0\x13\xcc\x9f\x60\x9e\x6e\x4c\x5a\xf5\x9f\xa3\x0f\x7a\x11\xf3\x37\xe4\xbf\x41\xdf\x0b\xbc\xed\x8b\x2e\xfa\xb9\xc9\xe6\x70\x58\xda\x96\x3c\x76\xdd\xec\x5f\x00\x00\x00\xff\xff\x65\x81\xbc\x4b\x42\x04\x00\x00"),
		},
		"/jobs.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "jobs.html.tmpl",
			modTime:          time.Date(2019, 7, 31, 20, 20, 29, 671911222, time.UTC),
			uncompressedSize: 7537,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x8f\xdb\xb8\xf1\x7f\xbf\x9f\x62\xc2\x04\xb1\x8d\x44\xf6\x05\x38\xe0\x0f\x6c\x24\x1f\x0e\xc9\xff\x90\x2e\x76\x2f\x87\x66\xdb\x17\x2d\x8a\x1c\x25\x8e\x2d\xa6\x34\x29\x90\xd4\xee\x1a\x8e\x3f\x51\x3f\x42\xdf\xdd\x27\x2b\x48\x4a\xd6\x83\xe5\x87\xe4\xb2\x28\x5a\x34\x2f\xd6\x12\x67\x38\x9a\xf9\xcd\x03\x67\x98\xcd\x86\xe1\x82\x4b\x04\x92\x29\x69\x51\x5a\xb2\xdd\x5e\xc4\x8c\xdf\x41\x26\xa8\x31\x89\x5f\xa6\x5c\xa2\x86\x9a\x01\xe4\x32\x72\xcf\x5a\x09\x81\x3a\x21\x57\x2a\xbd\xe6\xc6\xbe\xb1\x5a\x90\xf9\x05\xc0\x66\x63\x71\x55\x08\x6a\x11\x88\xb1\xd4\x96\x86\xc0\x74\xbb\xbd\x00\x68\xcb\xd5\xea\xde\x73\x03\xc4\x0b\xa5\x57\xe1\xb1\xcb\xe2\xd6\xa3\xa5\x56\x65\x41\x6a\x72\x97\x81\xcb\xa2\xb4\x7b\x1c\x07\x79\x22\xca\x98\x92\x64\x1e\xf3\x9a\xb8\x14\xeb\x22\xe7\x99\x92\xb0\x7b\x8a\x0c\x52\x9d\xe5\x64\x1e\xcf\xf8\x3c\x9e\x31\x7e\xd7\x91\xec\xc5\x01\x67\x09\x59\x70\x61\x51\x13\xb0\xeb\x02\x13\x62\xf1\xc1\x92\x8e\xe2\x15\x44\x04\x0a\x41\x33\xcc\x95\x60\x0e\xac\x9f\xaa\x5d\x72\x19\xad\x14\x43\x91\x90\xf0\xbd\x2b\x95\x7e\x89\x09\x2d\x4e\x80\x98\x42\xae\x71\x91\x90\xa7\x04\x18\xb5\x34\xb2\x6a\xb9\x14\x98\x90\x95\x62\x54\xd4\x6b\x54\x2f\xd1\x26\xe4\xe9\x27\x95\xde\x56\xfe\xb9\x09\x74\xcb\xad\xe3\xfe\x19\xef\x1d\x36\x1d\xd1\xee\xdf\x11\xac\x0a\xe1\x9c\x4b\x35\xa7\x51\xce\x19\x43\x99\x10\xab\x4b\xac\xc1\xa3\x1d\x8b\xba\x58\x76\x5e\x5b\x2f\xf1\xac\x09\x87\xd8\xd2\x54\x60\xad\x40\x78\xf1\x7f\x23\x63\x35\x2f\x90\xed\x80\x88\x6d\x8e\x94\xb5\xa4\x5b\xdd\xf9\xb6\xcd\xe7\x57\x2a\x8d\x67\x36\xef\x2f\x7f\xc8\x72\x64\xa5\xc0\x41\x5a\x99\x65\x68\xcc\x10\xe9\xff\xb5\x56\x7a\x88\x70\x4d\x8d\x05\x7c\xc0\xac\xb4\x5c\xc9\x21\x8e\x9f\xf1\xe1\x04\xc7\x07\x9f\x35\x43\x94\x1f\x33\xb7\xa7\x47\x8a\x67\x8d\xb5\x8e\xd2\x42\x22\xb6\x0b\xa5\x6c\x1b\x17\x06\x99\x12\xa6\xa0\x32\x21\xff\x77\x30\xe2\x8a\x52\x88\x48\xf3\x65\x6e\xfb\xa1\x56\x8a\x1d\x0f\x5d\x72\x49\x9d\x3a\x3d\x1e\x80\x58\x70\x5f\x24\x02\xe3\x86\x71\xe3\x9c\xc6\x2e\x21\x2b\xb5\x46\x69\x7f\xa1\x4b\x84\x24\x81\xef\xb6\x7b\x5b\x77\xd1\x1c\x04\xf0\xec\xef\x2e\xd3\xb4\xf1\x7b\xc6\x13\x32\xff\xc9\xbd\xf4\x42\x2b\x18\x2e\xf8\x23\xeb\x51\x68\xbc\xab\xd5\xf8\xed\x1f\xf0\x8b\xc6\xbb\x83\x8a\x0c\x6b\xa2\xb1\x40\x6a\x13\x22\x81\x4b\xd0\x54\x2e\x71\x5c\xd0\x25\xb2\x3f\x58\x5c\x99\xa9\x40\xb9\xb4\xf9\xcb\xb6\x76\x9d\x17\x78\x01\x4b\x5a\x4c\x80\xec\xe9\xda\x32\x92\x66\x96\xdf\xe1\x25\x48\x67\x58\x6b\xf3\x96\xb4\x2c\x31\xb8\xc3\xf3\x98\xdd\x29\x97\xcc\x29\xfb\x02\x5e\x91\x7e\x3a\x9f\x32\x75\x1f\xf4\x71\x4b\x9b\x89\xd3\x6e\xcf\x74\x88\xe0\xd5\x79\xae\x90\xf8\xb0\xb3\xc0\xe7\xd3\x6f\xff\xfc\xbd\x31\xf1\x2d\xd5\x13\xb4\x09\x58\x57\x10\xce\xd2\x2d\x9e\x95\xe2\x68\xc1\xb4\xac\x95\xe5\xed\xbc\x8e\x6d\xaa\xd8\xba\x53\xfd\xda\xd1\xf6\x49\xa5\x2e\xde\x1a\x7b\xfe\xda\x32\xf5\x6f\xf0\x19\xc2\x51\x76\x79\xe0\x24\xb2\x6c\x1e\x53\x27\x2f\x1c\x32\x9f\x54\x6a\x66\x9b\xcd\xaf\x9b\x0d\x7c\x52\xe9\x54\xd2\x15\xc2\x76\xfb\xeb\x76\x3b\xdb\xd5\x34\x43\x9a\xe8\x71\x3c\x8c\x9b\x42\xd0\xb5\x67\x7d\x92\xc0\x68\x04\x3f\x40\x7f\xfd\x72\x27\x2d\x84\x5a\xdb\xdc\x4a\x8d\xae\x50\x53\x15\x6e\x72\x06\x6b\xa8\xe3\x1f\x33\x55\x4a\x7b\x06\x3f\xba\xe2\x7e\x84\xbb\xe3\xb6\xdd\xce\xb1\xdb\xea\x3c\xff\xb1\xfa\x1e\xcc\x61\xb7\xe4\x45\x56\x66\x77\x78\x2e\x7b\x3c\x13\xf8\xec\x0e\x6c\xbc\x1c\xad\xd7\xeb\x75\x74\x73\x13\x31\x06\xef\xde\x5d\xae\x56\x97\xc6\xc0\x5f\x46\x5d\xdf\x9c\xb2\xc4\xe5\xc9\x49\x81\x03\x52\x7a\x91\xd9\x4e\x19\x17\x4d\x9f\x21\x34\x76\x6f\xdc\x12\xd9\xeb\x10\xc2\xf9\x7f\x5c\xcf\xbd\x16\x66\x17\x60\x4f\x43\x70\x3d\xe3\x92\xe1\x43\x08\xad\xa8\xd3\xc9\x1c\xeb\x6e\x0e\xee\xed\xa5\x5f\xd5\xf3\xfc\x99\xe3\xfd\xfe\xf1\x75\xb4\x3f\xc4\x35\x46\xaa\x40\x79\xb0\xef\xe9\x25\x75\x2f\xf3\x9f\xcb\xd4\x14\xaf\x0f\xf6\x6f\x4d\x0d\xd1\xa5\xbc\x52\xe9\xb8\xce\x8a\x89\xa7\x99\x5c\xdd\x27\xe4\x89\x2e\xa5\xe4\x72\xf9\x71\xdf\xda\x5d\x37\xf7\xc7\x52\x7e\x79\x37\x47\xd7\xe7\x76\x73\x1d\xad\x4f\x74\xd4\x99\x5a\x82\x54\x7a\x45\xab\x9e\x22\x32\x05\x97\x12\xf5\xe0\x39\x16\x2c\x3c\x62\xe0\xb0\x36\x67\xc3\x1a\xa2\xa7\x87\x6c\x85\xd9\xad\xa7\x9d\x32\xa7\xa0\xa5\xc1\x7d\xdd\xcf\xc5\xed\x84\xa6\x5f\x1f\xe2\x11\x43\x81\x76\x4f\xb3\x26\x6a\x3c\xfd\x44\xd8\xbc\x0d\x32\x4e\x40\x60\x35\x35\xf9\xd7\x43\xb0\xb3\xf7\x11\x22\xe7\x98\x91\xc3\x63\x49\xbb\x16\xf5\x5a\xe9\xe6\x58\x8d\x67\x7e\xf0\xa8\xba\x9d\xf8\x49\x14\x81\x9f\x9e\x20\x00\x06\x51\x54\xf1\xb9\x2e\xfa\x77\x9d\xbd\x35\x1e\xde\xa7\xb0\xa0\x6c\xe7\x51\x37\x75\x9e\x72\xbe\x9b\x8f\x3c\x31\x21\xd1\x2b\x02\x5a\x39\xa7\x32\x4e\x85\x5a\x56\xa9\x2d\x68\x8a\x42\x20\x4b\xd7\x09\x59\xad\xbd\x11\xd7\x6e\x89\x0c\x8d\xe0\x95\xec\x6a\x7f\x25\x4d\x65\xe5\x0a\xa5\x3d\x30\x94\x87\x2d\xf5\x65\xc1\xa1\x09\x23\x70\x39\x7c\x7b\xe5\xf7\x47\x8d\xb0\x56\x25\x98\x52\xe3\x0f\x87\xbb\xa1\x21\x69\xae\x25\x42\xdd\x9f\x58\xd2\xd2\x5a\x25\x6b\xce\xd4\x4a\x48\xad\x8c\x18\x2e\x68\x29\x6c\xbb\x32\x04\x04\x7b\x95\xe1\x8e\x8a\xb2\x49\x8b\x5e\xd8\xf9\xbc\x64\xdc\xac\xf8\x4e\x09\x32\x7f\xa3\xe4\x82\xeb\x55\x3c\x0b\x1f\x1e\xd6\x26\x5c\x1a\x84\x17\x72\x50\xb7\x61\xf9\x42\x19\x1c\x92\x7e\xf6\x80\xed\x1f\xfb\x71\xfc\xa8\x01\x7c\x34\x74\x6b\x54\xff\xa3\x22\xd7\x0d\xd9\x07\x63\x6d\xd0\xbb\x99\xf3\xdb\xb0\x4f\x5b\xe6\x25\xc4\xfb\x97\xcc\x63\x37\xa3\x9f\x53\x61\x9f\x5b\xbe\x42\xf3\x3a\x9e\xb9\x0d\xf3\x03\x61\x97\x7f\xdf\x55\xdf\x57\xfb\xe0\x96\x0e\x92\x4d\xf7\x38\x72\xad\x11\x8c\xe0\x45\xa7\x2f\xcf\xbf\x3f\x12\x71\x07\x33\xdc\x49\xe5\x92\x87\x68\xba\x32\x4a\x26\x56\xb9\x1f\x97\x68\xfd\x11\x34\xb6\xf8\x60\xa9\x46\x0a\x25\x8f\x32\xc5\x70\xc5\x7d\xeb\xdc\x79\x8b\x54\x61\x4d\x42\x90\x71\xab\xf4\xfb\xa2\x99\x39\xaa\x3b\xb5\xea\x33\xbe\xb5\xad\xc4\x7d\xa1\xde\x47\x6b\xc9\x63\x64\xef\x97\x54\xaa\xb2\x70\x0d\x7d\x55\xa9\x9c\xa1\x4d\xa1\x0a\xa4\xb3\x0a\xd5\x9f\x3c\xeb\x37\xac\x24\x57\x2a\x85\xfa\x6a\x71\xa8\xaa\x1c\xa8\x0b\xfb\x17\x92\xff\xcb\xfd\x7f\x73\xee\xbf\xd1\x48\x2d\x8e\xbe\x3e\xe3\x1f\x27\xa7\xeb\x30\xf9\x6f\xce\xed\xcc\x43\x5f\xe5\x76\xdb\xe0\x26\xc7\x83\x77\xce\x6b\x46\x3c\xeb\x37\xc9\xf1\xea\xa9\xfe\x31\x99\xe6\x85\xbf\x81\x7a\x36\xae\x73\x6b\x32\xd5\x48\xd9\x7a\xbc\x28\xa5\xbf\xa7\x86\xf1\x04\x36\x5e\xcc\x6c\x06\xd7\xdc\x58\x94\xb0\x50\x1a\x6c\x8e\xee\x54\x01\xad\xee\x0d\x58\x05\xb4\x28\x90\x6a\xcf\xf8\x6c\x3c\x9a\xfa\x4e\x7b\x34\x99\xba\x60\x1c\x8f\xde\xbe\xbf\xf9\x50\xa6\x56\xa3\xab\x0e\x7c\xc1\x91\x8d\x5e\x42\xf3\x05\xac\x3f\x01\xc0\x17\x30\xc6\x69\x98\x93\xa6\x7e\x4e\x78\x77\x7b\x73\x5d\x5f\xde\xcd\xe1\xbb\x86\x35\x30\xdf\x73\xc9\xd4\xfd\x54\xa8\xcc\xdf\x63\x4f\x73\x6a\x72\x78\xfe\x1c\x9e\x0d\x52\x26\x95\xa4\xb6\x14\x38\xc8\xeb\x1d\x30\x1e\xb9\xa1\x64\x34\x69\xc6\xbd\xed\x45\xfb\x77\x3b\x79\x7d\x51\xe3\x73\x9b\x73\x03\x26\x57\xa5\x60\x90\xa2\xeb\xb9\x1c\x4c\x56\x15\x20\xf0\x0e\x05\xd4\x20\x3b\xc0\x1c\x32\xee\xd7\x71\x2c\x4a\x5b\x6a\x84\xb7\xef\x6f\x00\x05\xae\x4c\x2d\xcf\xe6\xd4\x02\xd5\x08\x52\x59\xa0\xc2\x7b\x06\x0a\x8d\xc6\xc9\x50\x12\x88\x5f\x21\xd3\x0a\xf6\xc6\x87\x4a\x06\xb5\xa7\xa9\x09\x56\x1c\xc3\xdb\xa5\x8b\xf2\xb8\x6b\x74\x91\xca\x6e\x3d\xfc\x13\x78\x92\x00\x29\x65\xf8\xcf\x45\x46\xda\x98\x0d\xa2\x9e\xc0\xb3\x3d\x19\x53\x6a\xad\x1e\x8f\xdc\xa8\xd8\x40\xd8\x03\xae\xaf\x78\xa8\x8f\x27\x55\xcf\xb9\xb1\x4a\xaf\xa7\x45\x69\xf2\x0f\x96\x5a\x1c\x13\xf2\x72\x07\xf1\xd4\x97\xc8\x97\x7b\x9a\x16\xd4\xe6\xfe\xb2\xf2\xc5\x1e\x29\xb4\xbf\x95\x96\x5b\xff\xeb\xfe\xc6\xb3\x3a\x4f\x36\x1b\x94\x6c\xbb\xbd\xf8\x57\x00\x00\x00\xff\xff\xe8\x2e\x0d\xb0\x71\x1d\x00\x00"),
		},
		"/status.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "status.html.tmpl",
			modTime:          time.Date(2019, 7, 31, 20, 55, 52, 848272784, time.UTC),
			uncompressedSize: 1179,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\xbd\xee\xd3\x30\x14\xc5\xf7\x3c\xc5\x95\x19\x0a\x83\x89\xc4\xec\x98\x09\x06\x56\xd8\x2b\xc7\xbe\x49\x0d\xb7\x76\xe5\x8f\x52\x14\xe5\xdd\x51\x12\xa7\x1f\xa2\x40\x05\xea\x7f\xa9\xab\x73\x8e\x6f\x7e\xc7\x8e\x32\x0c\x06\x3b\xeb\x10\x58\x4c\x2a\xe5\xc8\xc6\xb1\x12\xc6\x1e\x41\x93\x8a\xb1\x61\xc1\x7f\x67\xb2\x02\x98\x35\x6b\x1a\xb6\xc7\x18\x55\x8f\x4c\x8a\xda\xd8\xa3\xac\xd6\xe5\x77\x7b\x8a\x76\x50\x0e\x09\xe6\x5f\x6e\xb0\x53\x99\xd2\x9c\xb9\x93\xe2\xad\x37\x3f\x60\xc1\xe1\xb3\x52\x92\x00\x22\xd3\x1a\x25\x1b\x13\xb7\x8e\x26\xf6\x82\xbe\xa6\x00\x04\xd9\x35\x97\xf0\x94\xb8\x46\x97\x30\x80\xf6\xc4\xf7\x86\xbf\xbb\x4a\x02\x88\x98\x82\x77\x3d\xb8\x9e\xb7\xd6\x99\x86\x7d\xf5\x6d\x7c\x4b\xe8\xfa\xb4\x9b\x5a\x2e\xf6\xcd\x8e\x83\x14\x71\xaf\x88\xe4\x17\x9f\x14\xc1\x27\xdf\x46\x51\x2f\xca\x05\xa1\x26\xfb\x5f\x40\x52\x28\xd8\x05\xec\x16\xa0\x57\xef\x3b\x4b\x09\x43\x13\xb3\xd6\x18\x23\x9b\x78\x35\x59\xfd\xad\x61\x8b\xf3\x7a\x53\xac\xcd\x1b\x76\x29\x53\xb4\x2e\xd3\x76\x1a\x33\x15\x52\x7f\x2e\xb5\x72\x96\x0b\x58\x9f\x27\x3f\x9f\x27\xbd\x6c\xe1\x4e\x59\x42\x73\xaf\xef\xe2\xdc\xd4\x5d\xa4\x7f\xac\x6a\x94\xeb\x31\x30\xf9\x71\x1e\xf2\xa4\x96\xbf\x5c\xcd\x56\xfb\xec\xd2\xdf\xde\xb4\xab\xd3\xff\x70\x42\x9d\x93\xf5\xee\xa9\x74\x18\x82\x0f\x8f\xb1\x95\xf3\x7a\x90\x4b\xd4\xb9\x38\xe5\xd3\x71\xfe\x53\x96\x61\x40\x67\xc6\xb1\xfa\x19\x00\x00\xff\xff\xe0\x82\x82\x52\x9b\x04\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/dashboard.html.tmpl"].(os.FileInfo),
		fs["/executions.html.tmpl"].(os.FileInfo),
		fs["/index.html.tmpl"].(os.FileInfo),
		fs["/jobs.html.tmpl"].(os.FileInfo),
		fs["/status.html.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
