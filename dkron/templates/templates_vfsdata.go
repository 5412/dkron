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
			modTime: time.Date(2019, 7, 1, 21, 22, 59, 41823471, time.UTC),
		},
		"/dashboard.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "dashboard.html.tmpl",
			modTime:          time.Date(2019, 6, 30, 22, 26, 54, 553483698, time.UTC),
			uncompressedSize: 4065,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\xcc\x12\x05\x36\x01\x2a\xf1\xd0\xdb\x56\x36\x10\x24\x0b\x74\xdb\x6e\x36\xd8\x06\x05\x8a\xa2\x08\x46\x22\x2d\x31\xa6\x48\x96\xa4\x5c\x1b\x86\xff\x7b\x41\xc9\x92\x25\x47\xde\xd8\x6e\x7a\x12\xbf\xe6\xcd\x9b\x79\xc3\x0f\x25\xef\xee\xbe\xdc\x3e\xfe\xf1\xf0\x11\x0a\x5f\xca\xd9\x24\x09\x1f\x50\x79\x84\xc6\x4c\x09\x5b\x58\xad\xc8\x6c\x32\x49\x0a\x8e\x6c\x36\x01\x48\xbc\xf0\x92\xcf\x36\x1b\x88\x6f\x75\x59\x6a\x15\xdf\x63\xc9\x61\xbb\x4d\x68\x33\x13\xd6\xb8\xcc\x0a\xe3\xc1\xd9\x6c\x4a\x7a\x2b\x6f\x9c\xe3\xde\x3d\xa0\x2f\x60\xbb\xa5\xcf\x8e\xa2\xca\x2b\x89\x36\x2e\x85\x8a\x9f\x1d\x99\x25\xb4\xb1\xbc\x04\x24\xb2\xba\xf2\x7c\x04\xea\x1c\x2c\x86\xae\x48\x35\x5a\xf6\x92\x8e\x14\x6a\x01\x96\xcb\x29\x71\x7e\x2d\xb9\x2b\x38\xf7\x04\x0a\xcb\xe7\xdf\x80\xcc\x9c\xa3\xb9\x5c\x9b\x42\x64\x5a\x45\xce\x08\xa5\xb8\x8d\x33\xe7\xc8\x79\xbc\x9e\xff\xae\xb8\x5d\x1f\x09\xee\x62\x66\xa9\xd6\xde\x79\x8b\xa6\x06\x6e\x58\x9d\x41\x6a\x68\xfe\x86\xbc\xd0\x98\x0b\x72\xc4\x7e\x18\xaf\xa3\x8b\x68\x58\x91\x2d\x5c\x81\xff\x74\x8d\x0b\x52\x34\x8e\x71\x59\x9d\x77\x45\xfe\x02\xf3\xad\xf2\xae\x19\x2f\x85\xb5\xda\x52\x29\xd2\x5e\xf7\xcc\xa0\x8f\xe3\x5c\x16\x78\x0f\xaf\xd4\x8c\xd3\x67\x5c\x62\x63\xd8\x6b\xfe\x77\x6c\x64\x4c\x2b\xca\x84\x33\x12\xd7\x14\x2b\xaf\x2d\x9f\x5b\xee\x8a\xcb\xcf\xa5\x4a\x44\x47\xa3\x3f\x90\xa9\xd0\xd6\x67\x95\x87\x70\x4a\xbc\xaa\xd4\x1c\x97\x61\x5d\x2c\x32\x4d\xc0\xaf\x0d\x9f\x12\x51\x62\xce\xe9\x2a\xaa\xed\x0f\x8a\xfe\x8d\x30\xf7\x91\x07\x7c\x80\x25\x5a\xb8\xfb\xe5\xeb\x97\xfb\xa7\x9b\x87\x4f\x4f\x0f\x37\x8f\x3f\xc1\x14\x06\x0e\x1e\x3e\xed\xd0\xc9\x8f\xb5\xc5\x77\x57\xf3\x4a\x65\x5e\x68\x05\x57\xd7\xb0\xa9\xc7\xc2\xe8\xfb\x3f\x19\x7a\x8c\xbc\xce\x73\xc9\xa7\xc4\x6b\x2d\xbd\x30\xe4\xaf\xf7\xd7\xf1\xae\x7d\x75\x5d\x2f\xde\x86\xcf\x3e\x87\x09\x6d\xee\xa4\x49\x92\x6a\xb6\xae\xa3\x66\x62\x09\x99\x44\xe7\xa6\x44\xe1\x32\x45\x0b\xcd\x27\x62\x7c\x8e\x95\xf4\x6d\x77\x2e\x56\x9c\x45\x5e\x1b\x02\x56\x07\xa7\x0a\x97\x22\xc7\xc0\x8d\x34\xe1\xf5\xa1\x32\xad\x3c\x0a\xc5\xed\x6e\x0e\x20\x79\x17\x45\x70\xab\xa5\x44\xe3\x38\x83\xbd\x35\x44\x51\xb7\xe6\x05\x99\x28\xf0\xed\xa1\xec\x70\x3e\xae\x0c\x2a\xc6\x2d\xa4\x95\xf7\x03\x08\x80\x64\x37\xd6\x48\xd2\x74\xc8\x01\x6a\x93\x38\x02\x83\x2c\x66\x3b\x72\xed\x30\xda\x9c\xfb\x29\x89\x77\x36\xdd\xf4\xde\x55\xd0\xd7\xa0\x6a\xc1\x9d\x8d\xb4\x92\x6b\x32\x7b\xac\x11\x7b\x31\x26\x34\xac\x3b\x6a\x58\xdf\x74\x29\xda\xba\xdc\xff\xe7\x85\x09\x6d\x52\x52\x97\x67\x2f\xa5\x9f\x51\x28\xa8\x9f\x23\xc3\x6c\xf6\x24\x49\x2d\x2a\x36\x0c\x1f\xdb\x39\xa9\x73\x3d\xb2\x65\xda\x72\x9e\x25\xa2\xcc\x5f\xd9\xff\xa2\xcc\x69\xfd\x78\x8a\x02\x58\x6c\x54\x4e\x80\xce\x12\x8a\x7d\xf2\x4c\x2c\xbb\x6a\x69\x3a\x93\x97\x75\x71\x7a\x79\xb5\xa2\xc2\x88\xba\xfb\xb4\x8c\xa2\x01\x24\x95\x6c\xe1\x4a\x14\x2a\x0a\x07\x88\x0b\x8b\xdb\x3d\xa3\x70\x39\x4c\x97\x14\xfd\x6e\x9d\xbf\xe3\x29\x1b\xac\x04\x48\x44\xeb\xac\x7b\x1d\xc1\xfe\x9d\xd4\xbd\xc2\x82\xe2\xe2\xd0\xf6\xae\x9d\x1d\x7a\xef\xa7\x36\x74\x87\xf4\xce\x61\x4b\x9f\x75\xea\xce\xa4\xec\xd1\x2d\xdc\x28\xdd\x9f\x75\xea\x4e\x67\x9a\xd0\x4a\x1e\xa8\xf6\x1b\x47\x9b\x15\xdf\xc3\xfd\x5e\x38\x54\x0c\xbe\x72\xa3\xa1\x51\xe9\x88\x8e\x43\xf1\xda\xa6\x15\x79\xe1\x4f\x54\xb2\xf0\xde\x7c\xa0\x4d\x21\xc7\x22\x5c\x0d\xbb\x73\xe4\x29\x95\xa8\x16\x67\xe6\x28\xd5\x7a\x31\xae\xa8\xce\x2e\x4c\x51\x6f\x0f\x75\xcd\xae\xb1\xd9\x78\x5e\x1a\x89\x9e\x43\x7d\x86\x73\xe5\x09\xc4\xdb\xed\x37\x6f\x8a\xe6\x6a\x48\xb5\xf7\xba\x6c\xc7\x84\x5a\x72\xdb\x6d\xa7\x91\x9d\x57\x3f\xeb\xc9\xd8\xd6\xec\xee\x0e\xf0\x7c\xe5\xa3\x8c\x2b\x3f\xbc\x01\x06\x72\xbd\x22\x8b\x39\x3c\xf9\xf9\xca\x93\x91\x5f\x31\xe8\x0d\xfd\xce\xad\x0b\x25\x13\x7e\xd0\xcc\x79\x3b\x64\xdc\xdd\x67\x5e\xa6\xdc\x7e\xe8\xfb\x68\x86\xba\xff\xc0\x37\x71\xf3\x6b\x7d\x57\x0e\xdc\x34\x43\xa7\xb8\x39\xb1\x44\x12\xda\xbc\x1c\xc2\x53\xa2\xfe\xf9\xfd\x37\x00\x00\xff\xff\xe7\x77\xe7\x5b\x0d\x0f\x00\x00"),
		},
		"/executions.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "executions.html.tmpl",
			modTime:          time.Date(2019, 4, 19, 19, 42, 44, 166523762, time.UTC),
			uncompressedSize: 2615,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x51\x6f\xdb\xb8\x0f\x7f\xef\xa7\x60\xf5\x0f\x1a\x07\xff\xda\x5e\x87\x3e\x6d\xb6\x81\xdd\x70\x77\x40\xb1\xdb\x01\xd7\xc3\xbd\x1e\x64\x8b\xb1\xb4\x3a\x92\x21\xd1\x4b\x03\xcf\xdf\xfd\x20\x39\x4e\x9c\x26\xc5\xd0\x3e\xb8\x12\x49\x91\x3f\xfe\x48\x51\xe9\x7b\x81\x6b\xa5\x11\x58\x65\x34\xa1\x26\x36\x0c\x57\x99\x50\xdf\xa1\x6a\xb8\x73\x79\x10\x73\xa5\xd1\xc2\x64\x00\xba\x8e\xfd\xda\x9a\xa6\x41\x9b\xb3\x5f\x9f\xb1\xea\x48\x19\xed\x3e\x93\x6d\x58\x71\x05\x30\x77\x60\xcd\x36\xc8\xf6\x52\x25\x72\xb6\x41\xe7\x78\x8d\xac\xc8\x52\xa1\xbe\x87\x03\x87\xc5\xe5\x93\xf2\xae\x38\x86\xc9\x52\x79\x37\xc9\xdf\x17\x8f\xd2\x6c\x95\xae\xe1\x0b\x77\x04\x77\xef\xde\xc1\x83\x29\x01\x0f\xc6\xb0\x36\x16\xbe\x99\x12\xfa\x3e\x79\x30\xe5\x57\xbe\xc1\x61\x08\x87\x01\x32\x0e\xd2\xe2\x3a\x67\xff\x1b\x93\x6a\x54\xf5\x94\x33\xdb\xe9\x07\x53\x46\xcb\xf9\x81\xe5\x2a\x58\x38\x69\xb6\x39\xbb\xb6\x9d\xd6\x4a\xd7\xff\x7e\x33\x25\x03\x52\xd4\x60\xce\xfe\xea\x34\x2b\x32\x35\x61\xaf\x9b\x5d\x2b\x55\x65\x34\x1c\x56\x71\xdb\xf0\x1d\xdb\x87\xf6\x7f\xdc\x2a\x1e\x4b\x25\x04\xea\x9c\x91\xed\x02\x21\xaa\xc8\x52\x5e\x9c\x03\xfc\x89\xef\xca\xd4\xa0\x8d\xdd\xf0\x26\xb6\xaa\x96\x14\xbb\x56\x69\x8d\x76\x06\x7b\x8e\xfa\x34\x50\x96\xca\xf7\xb3\x2a\x5c\x01\xf4\xbd\xe5\xba\x46\x58\xd4\x4f\xf0\x21\x87\xe4\x97\xdd\xef\xd6\x74\x6d\x60\x2e\x93\xf7\xc7\x62\x7c\x80\xbe\x5f\xd4\x4f\xc3\x90\xa5\xf2\x3e\xf8\x20\x5e\x36\x38\x41\x1d\x37\xe1\x1b\x3b\xb2\xaa\x45\x31\x95\x94\xec\x21\x4b\x92\xc5\x83\x29\xb3\x94\xe4\x5c\xf4\x48\xdc\x12\x0a\xf8\x44\x2f\x35\xbf\x29\xad\x9c\xbc\xa8\xfa\x6a\x04\xbe\x94\xfd\xd9\x51\xdb\x9d\x59\x3e\x76\x55\x85\xce\x1d\xc5\x59\x3a\x41\x3a\x64\x8f\x4f\xb7\xb0\xc0\x67\x4f\x41\xa4\xb4\xc0\x67\x58\x24\x81\x08\xe7\x99\x59\xed\x1b\xe9\x24\x15\x51\xf4\xfd\x02\x9f\x8f\xad\x93\xa5\x24\xce\xd5\xfb\xe4\x3e\xd1\x6b\x06\x53\x8e\xaf\x5b\xf8\x54\x2f\x87\x38\xf4\xd8\xb1\x7f\xfa\x1e\x66\xa8\x60\x18\x62\x2f\xa9\x9f\xa6\x15\x86\xd5\xc6\x08\xde\x30\x10\x9c\x78\x4c\xa6\xae\x7d\x63\x9f\xc8\xb8\xad\x91\xde\xe8\xae\x98\xb5\xfc\x88\x7c\x2c\x08\xfc\x00\x32\x8f\x64\xfd\xe5\xfd\x01\x64\x3b\x5d\x71\xc2\x61\x48\x92\x64\x7e\x05\x2e\xd3\x37\x16\x6f\x9e\xfa\x58\xbe\x71\x79\x1d\xc7\xf0\x87\x0f\x0e\x71\x3c\x9b\x3d\xfb\xa6\x0c\xb0\x60\xcd\x05\xb2\x30\x8e\xde\xc4\x0d\xf1\x32\x74\x42\xce\xe2\x3b\x06\xd6\x78\x8a\x84\xe2\x8d\xa9\xa7\xab\x1d\xae\x75\xc3\x4b\x6c\x1a\x14\xe5\x2e\x67\x9b\x5d\xc0\xf2\xc5\x8b\x0e\x6c\x9c\x01\x8a\xf7\x5e\x26\x9f\xa6\xea\x36\x7e\xdc\xce\xaa\x79\x76\x64\x1a\xc9\x73\x8a\xcf\xad\x24\x72\x81\xf6\xc4\x08\x20\x2b\x3b\x22\xa3\x81\x76\x2d\xe6\x6c\xdc\xb0\xc3\xd0\x6f\x8c\xc3\x7d\xcd\x85\x72\x1b\x75\x70\xc6\x66\xe9\xe5\xec\x73\xb0\x2b\x32\xd7\x72\x7d\x69\x9c\xdd\x90\xda\xa0\xfb\x98\xa5\xde\xa0\xc8\xd2\x31\xcc\x0b\x20\xf2\xfe\x14\x6e\x18\xa8\x63\x65\x4e\x99\xfb\x47\xe1\x16\xcc\xd8\x3b\x7e\xaa\x9f\x5d\xb4\x71\x02\x1d\x1c\x4f\x6f\xca\xeb\xcc\x94\x46\xec\x5e\xf2\xd2\x5a\x2c\x5e\x69\x53\x1f\xc3\xab\xdf\x16\x64\x6d\x0c\xbd\x8d\xfe\x92\x34\x94\xa4\x63\x81\x6b\xde\x35\x74\xb9\x10\x45\x20\xff\x12\xa7\x2f\x30\x9d\x6c\x67\x9b\xd9\xb2\xef\x51\x8b\x71\xb8\xa7\x61\x58\x17\x57\x47\xe1\xde\x2e\x73\x95\x55\x2d\x79\xcd\x22\x9a\xba\x73\x95\x58\xe4\x62\x17\xad\x3b\x5d\xf9\xe7\x00\xa2\x15\xf4\xc1\xa5\x5a\x43\xb4\x55\x5a\x98\x6d\xd2\x98\x8a\x7b\x65\x22\xb9\x93\x70\x73\x03\x8b\x8b\x9a\x55\xd2\xa0\xae\x49\x4e\x1e\xe0\x55\xbb\x40\x40\xb4\xf4\xcf\xda\x72\xf5\x31\x18\x8f\x93\x78\x11\x2d\x47\xe5\x72\x95\x18\x3d\x5a\x24\xa5\xdb\xcb\x6e\xe1\x08\x13\x8f\x51\x3c\x52\x5f\x06\xb3\x86\x08\x13\x8b\x0d\x27\x14\x7f\x87\x49\xb7\x82\xeb\x1c\x58\xa7\xc7\xdf\x48\x82\x1d\x0f\x01\x5c\x4c\x2e\x87\xc5\x99\x8f\x84\x13\xd9\x68\xe9\x07\xf1\x84\x76\xc2\x3b\xec\xf7\x73\x42\x3d\xee\xf1\x16\xfd\x14\xb9\x54\x8e\x8c\xdd\x25\x6d\xe7\xe4\x23\x71\xc2\x88\xb1\x5b\x98\x5c\x25\xe1\x22\xdd\x9e\x01\x6d\x39\x49\xed\x07\xdd\xff\xcf\x54\x0e\xb9\xad\xe4\x44\x69\xf8\xef\xbf\x59\x3a\xd5\x7e\xea\x89\xff\x02\x00\x00\xff\xff\x83\x7f\x99\xd3\x37\x0a\x00\x00"),
		},
		"/index.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "index.html.tmpl",
			modTime:          time.Date(2019, 4, 19, 19, 42, 44, 166523762, time.UTC),
			uncompressedSize: 1090,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4d\x8f\xda\x30\x10\xbd\xf3\x2b\x46\xbe\xd0\xaa\x85\x48\xbb\xb7\x95\x13\xa9\xed\xa9\x52\xb5\xaa\xd4\xfe\x01\xc3\x0c\xc4\xc2\xd8\xd1\x78\x60\x8b\xd2\xfc\xf7\xca\x21\x21\xa6\xb0\xab\xbd\xc4\xcf\xef\x3d\x8f\x3d\x1f\x69\x5b\xa4\x8d\xf5\x04\x6a\x1d\xbc\x90\x17\xd5\x75\x33\x8d\xf6\x08\x6b\x67\x62\x2c\x7b\xda\x58\x4f\x0c\xa3\x01\xfc\x76\x91\x30\x07\xe7\x88\x4b\xf5\xdd\x23\xfd\xf9\x26\xec\x54\x35\x03\x68\x5b\xa1\x7d\xe3\x8c\x10\xa8\x28\x46\x0e\x51\xc1\xb2\xeb\x66\x00\x79\x54\x0e\x2f\xbd\xfb\x9a\x6d\x8c\x27\x07\xfd\x77\x81\xb4\x31\x07\x27\x83\xeb\x8e\x6f\x51\x93\x41\xeb\xb7\x17\x07\x80\xae\x1f\xaf\x2d\x62\xc5\x91\xaa\x7e\xd8\x23\xe9\xa2\x7e\xbc\xc4\x2a\xd0\x1e\x5f\x0f\xbc\x0a\x78\xca\xa3\xb2\x5d\xef\x62\x6d\x5e\x60\x04\x8b\xd0\x88\x0d\x3e\x96\x6a\x00\x6a\x92\x36\x64\xe4\xc0\x14\x4b\x35\xa2\x4c\x8c\xc4\x36\x49\xe7\x55\x55\xba\x18\xa5\x3b\x4f\xbb\xc0\x01\xbc\x5e\xc2\xfa\xa1\x7a\x0e\x48\x51\x17\xf5\xc3\x40\x89\x59\x39\x1a\xad\xe7\x4d\xff\x5d\x44\x61\xdb\x10\x4e\x75\x15\xce\x52\x95\xba\x7a\x36\x7b\xd2\x85\xd4\xd7\xec\x17\x44\xa6\x18\x6f\x85\x9f\x81\xe5\x96\xfd\xd5\x77\xfe\x96\xff\x6d\xb6\x57\xac\x2e\xa6\xeb\xb5\x70\x1a\x2d\xa6\x86\x8c\x94\x6a\x4f\xfb\x15\x31\x58\x0f\x67\x14\xe1\x2f\x04\x46\xe2\xaf\xa7\xa7\x79\x7a\xe4\x3c\xef\x91\x60\x3a\xbb\xb2\x1e\xc7\x93\xcb\xe4\x49\x25\x16\x7c\xdb\x97\x52\x7b\x8f\x2f\x65\xfa\x1e\xdf\x39\xf7\x3b\xce\x69\x33\x74\x32\xcb\xf6\xc3\x8e\x4e\x9f\xe1\x68\xdc\x81\x3e\x4e\x39\x2f\x53\xbd\x54\x7e\x0e\x40\xc7\xc6\xf8\xe9\xd2\x1d\x9d\xe0\x13\xcc\x9f\x60\x9e\x6e\x4c\x5a\xf5\x9f\xa3\x0f\x7a\x11\xf3\x37\xe4\xbf\x41\xdf\x0b\xbc\xed\x8b\x2e\xfa\xb9\xc9\xe6\x70\x58\xda\x96\x3c\x76\xdd\xec\x5f\x00\x00\x00\xff\xff\x65\x81\xbc\x4b\x42\x04\x00\x00"),
		},
		"/jobs.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "jobs.html.tmpl",
			modTime:          time.Date(2019, 7, 1, 21, 22, 59, 42878051, time.UTC),
			uncompressedSize: 7483,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x8f\xdb\xb8\xf1\x7f\xbf\x9f\x62\xc2\x04\xb1\x8d\x44\xf6\x05\x38\xe0\x0f\x6c\x24\x1f\x0e\xc9\xff\x90\x2e\x76\x2f\x87\x66\xdb\x17\x2d\x8a\x1c\x25\x8e\x2d\xa6\x34\x29\x90\xd4\xee\x1a\x8e\x3f\x51\x3f\x42\xdf\xdd\x27\x2b\x48\x4a\xd6\x83\xe5\x87\xe4\xb2\x28\x5a\x34\x2f\xd6\x92\x38\x1c\xcd\xfc\xe6\x37\xc3\x19\x65\xb3\x61\xb8\xe0\x12\x81\x64\x4a\x5a\x94\x96\x6c\xb7\x17\x31\xe3\x77\x90\x09\x6a\x4c\xe2\x1f\x53\x2e\x51\x43\x2d\x00\x72\x19\xb9\x6b\xad\x84\x40\x9d\x90\x2b\x95\x5e\x73\x63\xdf\x58\x2d\xc8\xfc\x02\x60\xb3\xb1\xb8\x2a\x04\xb5\x08\xc4\x58\x6a\x4b\x43\x60\xba\xdd\x5e\x00\xb4\xf5\x6a\x75\xef\xa5\x01\xe2\x85\xd2\xab\x70\xd9\x15\x71\xcf\xa3\xa5\x56\x65\x41\xea\xe5\xae\x00\x97\x45\x69\xf7\x24\x0e\xca\x44\x94\x31\x25\xc9\x3c\xe6\xf5\xe2\x52\xac\x8b\x9c\x67\x4a\xc2\xee\x2a\x32\x48\x75\x96\x93\x79\x3c\xe3\xf3\x78\xc6\xf8\x5d\x47\xb3\x57\x07\x76\x5d\x60\x42\x2c\x3e\x58\xd2\x31\xb6\x82\x85\x40\x21\x68\x86\xb9\x12\xcc\x01\xf4\x13\x17\x16\xb5\xc7\x6d\xa5\x18\x8a\x84\x84\x77\x5c\xa9\xf4\x4b\xcc\x6e\x49\x02\xc4\x14\x72\x8d\x8b\x84\x3c\x25\xc0\xa8\xa5\x91\x55\xcb\xa5\xc0\x84\xac\x14\xa3\xa2\x7e\x46\xf5\x12\x6d\x42\x9e\x7e\x52\xe9\x6d\x15\x93\x9b\xb0\x6e\xb9\x75\xd2\x3f\xe3\xbd\xc3\xa3\xa3\xda\xfd\x3b\x82\x4f\x21\x5c\x40\xa9\xe6\x34\xca\x39\x63\x28\x13\x62\x75\x89\x35\x60\xb4\xe3\x51\x17\xbf\xce\x6d\xeb\x26\x9e\x35\x14\x88\x2d\x4d\x05\xd6\x06\x84\x1b\xff\x37\x32\x56\xf3\x02\xd9\x0e\x88\xd8\xe6\x48\x59\x4b\xbb\xd5\x9d\x77\xdb\x7c\x7e\xa5\xd2\x78\x66\xf3\xfe\xe3\x0f\x59\x8e\xac\x14\x38\xb8\x56\x66\x19\x1a\x33\xb4\xf4\xff\x5a\x2b\x3d\xb4\x70\x4d\x8d\x05\x7c\xc0\xac\xb4\x5c\xc9\x21\x89\x9f\xf1\xe1\x84\xc4\x07\x9f\x29\x43\x2b\x3f\x66\x6e\x4f\x6f\x29\x9e\x35\xde\xba\x95\x16\x12\xb1\x5d\x28\x65\xdb\xb8\x30\xc8\x94\x30\x05\x95\x09\xf9\xbf\x83\x8c\x2b\x4a\x21\x22\xcd\x97\xb9\xed\x53\xad\x14\x3b\x19\xba\xe4\x92\x3a\x73\x7a\x32\x00\xb1\xe0\xbe\x30\x04\xc1\x0d\xe3\xc6\x05\x8d\x5d\x42\x56\x6a\x8d\xd2\xfe\x42\x97\x08\x49\x02\xdf\x6d\xf7\xb6\xee\xd8\x1c\x14\xf0\xec\xef\x09\x59\x70\x6d\xfc\x9e\xf1\x84\xcc\x7f\x72\x37\x3d\x6a\x05\xc7\x05\x7f\x64\x3b\x0a\x8d\x77\xb5\x19\xbf\xfd\x03\x7e\xd1\x78\x77\xd0\x90\x61\x4b\x34\x16\x48\x6d\x42\x24\x70\x09\x9a\xca\x25\x8e\x0b\xba\x44\xf6\x07\x8b\x2b\x33\x15\x28\x97\x36\x7f\xd9\xb6\xae\x73\x03\x2f\x60\x49\x8b\x09\x90\x3d\x5b\x5b\x4e\xd2\xcc\xf2\x3b\xbc\x04\xe9\x1c\x6b\x6d\xde\x92\x96\x27\x06\x77\x78\x1e\xf3\x3b\xe5\x92\x39\x63\x5f\xc0\x2b\xd2\x4f\xe7\x53\xae\xee\x83\x3e\x6e\x59\x33\x71\xd6\xed\xb9\x0e\x11\xbc\x3a\x2f\x14\x12\x1f\x76\x1e\xf8\x7c\xfa\xed\x9f\xbf\x97\x13\xdf\xd2\x3c\x41\x1b\xc2\xba\x82\x70\x96\x6d\xf1\xac\x14\x47\x0b\xa6\x65\xad\x2c\x6f\xe7\x75\x6c\x53\xc5\xd6\x9d\xea\xd7\x66\xdb\x27\x95\x3a\xbe\x35\xfe\xfc\xb5\xe5\xea\xdf\xe0\x33\x2c\xfc\xa1\x74\x79\xe0\x24\xb2\x6c\x1e\x53\xa7\x2f\x1c\x32\x9f\x54\x6a\x66\x9b\xcd\xaf\x9b\x0d\x7c\x52\xe9\x54\xd2\x15\xc2\x76\xfb\xeb\x76\x3b\xdb\xd5\x34\x43\x1a\xf6\xd4\x32\x81\x40\x6d\x27\x2a\xe5\x5d\x51\x53\x95\x63\x72\x86\x68\xa8\xce\x1f\x33\x55\x4a\x7b\x86\x3c\xba\x92\x7d\x44\xba\x13\x8c\xdd\xce\xb1\xdb\xea\xe2\xf9\xb1\x7a\x1f\xcc\x61\xf7\xc8\xab\x84\x1f\x60\x4f\xe6\xb2\x27\x33\x81\xcf\xee\x18\xc6\xcb\xd1\x7a\xbd\x5e\x47\x37\x37\x11\x63\xf0\xee\xdd\xe5\x6a\x75\x69\x0c\xfc\x65\xd4\x45\xfc\x94\x27\x8e\xfd\x27\x15\x0e\x68\xe9\xf1\xad\x9d\x08\x8e\x23\x9f\x21\xb4\x68\x6f\xdc\x23\xb2\x77\xee\x87\x53\xfd\xb8\x9d\x7b\x8d\xc9\x8e\x36\x4f\x03\x65\x9e\x71\xc9\xf0\x21\x10\x26\xea\xf4\x27\xc7\x7a\x96\x83\x7b\x7b\x49\x55\x75\x32\x7f\xe6\x78\xbf\x7f\x28\x1d\xed\xf4\x70\x8d\x91\x2a\x50\x1e\xec\x66\x7a\xa9\xda\xcb\xe7\xe7\x32\x35\xc5\xeb\x83\x5d\x59\x53\x19\x74\x29\xaf\x54\x3a\xae\xb3\x62\xe2\xd7\x4c\xae\xee\x13\xf2\x44\x97\x52\x72\xb9\xfc\xb8\xef\xed\xae\x47\xfb\x63\x29\xbf\xbc\x47\xa3\xeb\x73\x7b\xb4\x8e\xd5\x27\x7a\xe3\x4c\x2d\x41\x2a\xbd\xa2\x55\xa7\x10\x99\x82\x4b\x89\x7a\xf0\x74\x0a\x1e\x1e\x71\x70\xd8\x9a\xb3\x61\x0d\xec\xe9\x21\x5b\x61\x76\xeb\xd7\x4e\xb9\x53\xd0\xd2\xe0\xbe\xed\xe7\xe2\x76\xc2\xd2\xaf\xa7\x78\xc4\x50\xa0\xdd\xb3\xac\x61\x8d\x5f\x3f\x41\x9b\xb7\x41\xc7\x09\x08\xac\xa6\x26\xff\x7a\x08\x76\xfe\x3e\x02\x73\x8e\x39\x39\x3c\x6c\xb4\x6b\x51\xaf\x41\x6e\x0e\xcb\x78\xe6\xc7\x89\xaa\x87\x89\x9f\x44\x11\xf8\x99\x08\x02\x60\x10\x45\x95\x9c\xeb\x8d\x7f\xd7\x89\x5a\xe3\xe1\x63\x0a\x0b\xca\x76\x11\xe5\x2c\x21\xa7\x82\xef\xa6\x1e\xbf\x98\x90\xe8\x15\x01\xad\x5c\x50\x19\xa7\x42\x2d\xab\xd4\x16\x34\x45\x21\x90\xa5\xeb\x84\xac\xd6\xde\x89\x6b\xf7\x88\x0c\x0d\xd3\x95\xee\x6a\x7f\xa5\x4d\x65\xe5\xca\x8d\xf4\xc3\xe3\x75\xd8\x52\x8f\xfd\x87\xe6\x86\x20\xe5\xf0\xed\x95\xdf\x1f\x35\xc2\x5a\x95\x60\x4a\x8d\x3f\x1c\xee\x71\x86\xb4\xb9\x46\x07\x75\x7f\x0e\x49\x4b\x6b\x95\xac\x25\x53\x2b\x21\xb5\x32\x62\xb8\xa0\xa5\xb0\xed\xca\x10\x10\xec\x55\x86\x3b\x2a\xca\x26\x2d\x7a\xb4\xf3\x79\xc9\xb8\x59\xf1\x9d\x11\x64\xfe\x46\xc9\x05\xd7\xab\x78\x16\x5e\x3c\x6c\x4d\xf8\x14\x10\x6e\xc8\x41\xdb\x86\xf5\x0b\x65\x70\x48\xfb\xd9\x63\xb3\xbf\xec\xf3\xf8\x51\x09\x7c\x94\xba\x35\xaa\xff\x51\xcc\x75\xa3\xf3\x41\xae\x0d\x46\x37\x73\x71\x1b\x8e\x69\xcb\xbd\x84\xf8\xf8\x92\x79\xec\x26\xef\x73\x2a\xec\x73\xcb\x57\x68\x5e\xc7\x33\xb7\x61\x7e\x80\x76\xf9\xf7\x5d\xf3\x7d\xb5\x0f\x61\xe9\x20\xd9\x74\x8f\x23\xd7\x1a\xc1\x08\x5e\x40\xbb\x2f\xcf\xbf\x3f\xc2\xb8\x83\x19\xee\xb4\x72\xc9\x03\x9b\xae\x8c\x92\x89\x55\xee\xc7\x25\x5a\x7f\xb0\x8c\x2d\x3e\x58\xaa\x91\x42\xc9\xa3\x4c\x31\x5c\x71\xdf\x3a\x77\xee\x22\x55\x58\x93\x10\x64\xdc\x2a\xfd\xbe\x68\x26\x89\xea\x4b\x59\xf5\x1a\xdf\xda\x56\xea\xbe\xd0\xee\xa3\xb5\xe4\x31\xb2\xf7\x4b\x2a\x55\x59\xb8\x86\xbe\xaa\x54\xce\xd1\xa6\x50\x85\xa5\xb3\x0a\xd5\x9f\xbc\xe8\x37\xac\x24\x57\x2a\x85\xfa\x83\xe1\x50\x55\x39\x50\x17\xf6\x3f\x33\xfe\x2f\xf7\xff\xcd\xb9\xff\x46\x23\xb5\x38\xfa\xfa\x8c\x7f\x9c\x9c\xae\x69\xf2\xdf\x9c\xdb\x99\x87\xbe\xca\xed\xb6\xc3\x4d\x8e\x87\xe8\x9c\xd7\x8c\x78\xd1\x6f\x92\xe3\xd5\x55\xfd\x63\x32\xcd\x0b\xff\x5d\xe9\xd9\xb8\xce\xad\xc9\x54\x23\x65\xeb\xf1\xa2\x94\xfe\xeb\x33\x8c\x27\xb0\xf1\x6a\x66\x33\xb8\xe6\xc6\xa2\x84\x85\xd2\x60\x73\x74\xa7\x0a\x68\x75\x6f\xc0\x2a\xa0\x45\x81\x54\x7b\xc1\x67\xe3\xd1\xd4\x77\xda\xa3\xc9\xd4\x91\x71\x3c\x7a\xfb\xfe\xe6\x43\x99\x5a\x8d\xae\x3a\xf0\x05\x47\x36\x7a\x09\xcd\x1b\xb0\x7e\x05\x00\x5f\xc0\x18\xa7\x61\x4e\x9a\xfa\x39\xe1\xdd\xed\xcd\x75\xfd\x49\x6e\x0e\xdf\x35\xa2\x41\xf8\x9e\x4b\xa6\xee\xa7\x42\x65\xfe\xeb\xf4\x34\xa7\x26\x87\xe7\xcf\xe1\xd9\xe0\xca\xa4\xd2\xd4\xd6\x02\x07\x65\x7d\x00\xc6\x23\x37\x94\x8c\x26\xcd\xb8\xb7\xbd\x68\xff\x6e\x27\xaf\x2f\x6a\x7c\x6e\x73\x6e\xc0\xe4\xaa\x14\x0c\x52\x74\x3d\x97\x83\xc9\xaa\x02\x04\xde\xa1\x80\x1a\x64\x07\x98\x43\xc6\xfd\x3a\x89\x45\x69\x4b\x8d\xf0\xf6\xfd\x0d\xa0\xc0\x95\xa9\xf5\xd9\x9c\x5a\xa0\x1a\x41\x2a\x0b\x54\xf8\xc8\x40\xa1\xd1\x38\x1d\x4a\x02\xf1\x4f\xc8\xb4\x82\xbd\x89\xa1\x92\xc1\xec\x69\x6a\x82\x17\xc7\xf0\x76\xe9\xa2\x3c\xee\x1a\x1d\x53\xd9\xad\x87\x7f\x02\x4f\x12\x20\xa5\x0c\xff\x4d\xc8\x48\x1b\xb3\x41\xd4\x13\x78\xb6\xa7\x63\x4a\xad\xd5\xe3\x91\x1b\x15\x1b\x08\x7b\xc0\xf5\x0d\x0f\xf5\xf1\xa4\xe9\x39\x37\x56\xe9\xf5\xb4\x28\x4d\xfe\xc1\x52\x8b\x63\x42\x5e\xee\x20\x9e\xfa\x12\xf9\x72\xcf\xd2\x82\xda\xdc\x7f\xc5\x7c\xb1\xb7\x14\xda\xdf\xca\xca\xad\xff\x75\x7f\xe3\x59\x9d\x27\x9b\x0d\x4a\xb6\xdd\x5e\xfc\x2b\x00\x00\xff\xff\x5f\x02\x9d\x71\x3b\x1d\x00\x00"),
		},
		"/status.html.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "status.html.tmpl",
			modTime:          time.Date(2019, 4, 19, 19, 42, 44, 166523762, time.UTC),
			uncompressedSize: 1053,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x93\x31\x4f\xfb\x30\x10\xc5\xf7\x7c\x8a\x93\x77\xff\x23\xfd\x67\xd7\x1b\x0c\xac\xb0\x57\x8e\x7d\x09\x46\x57\xbb\xf2\xd9\xa5\x28\xca\x77\x47\x4d\x9d\x52\x24\x50\x2b\x44\x59\x72\xd1\xbb\x77\x97\x5f\x9e\xe5\x71\x74\xd8\xfb\x80\x20\x38\x9b\x5c\x58\x4c\x53\xa3\x9c\xdf\x81\x25\xc3\xbc\x12\x29\xbe\x0a\xdd\x00\xcc\x9a\x77\x2b\xb1\x41\x66\x33\xa0\xd0\xaa\x75\x7e\xa7\x9b\xa5\x7c\x37\x53\xb5\xad\x09\x48\x30\x3f\xa5\xc3\xde\x14\xca\xb3\xe7\x0b\x97\xec\xa2\x7b\x83\x23\x8e\x9c\x95\xea\x04\x50\x85\x16\x2b\x79\xce\xd2\x07\x3a\xb0\x57\xf4\xc5\x05\xa0\xc8\x2f\xbe\x8c\xfb\x2c\x2d\x86\x8c\x09\x6c\x24\xb9\x71\xf2\xff\x99\x13\x40\x71\x4e\x31\x0c\x10\x06\xd9\xf9\xe0\x56\xe2\x25\x76\xfc\x8f\x30\x0c\xf9\xf9\xf0\x97\xc7\xf6\xa7\x89\xad\x56\xbc\x31\x44\xfa\x29\x66\x43\xf0\x10\x3b\x56\xed\x51\xf9\x40\x68\xc9\xff\x12\x10\x17\x6b\x91\xb9\x2f\xb4\x3e\xb0\x5d\x80\x5a\xbe\x53\x03\xac\xc3\x42\x3f\x9e\xb6\xdc\x1c\xb8\x37\x9e\xd0\xfd\x00\xd6\x99\x30\x60\x12\xfa\x7e\x5e\xf0\x57\xc1\xae\x6d\x2c\x21\x5f\x3a\xeb\xb3\xfc\xee\xf6\x68\x4b\xf6\x31\xdc\x94\x0e\x53\x8a\xe9\x3a\xb6\x9a\xd7\x95\x5c\xaa\x2d\xb5\x53\x2f\xef\xe9\xa5\x96\x71\xc4\xe0\xa6\xa9\x79\x0f\x00\x00\xff\xff\x33\xc9\xe6\xe6\x1d\x04\x00\x00"),
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