// Code generated by go-bindata.
// sources:
// html/index.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _htmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x7b\x77\xdb\xb6\x92\xff\x5b\x3d\xa7\xdf\x01\xcb\xf8\x26\xd2\xad\x5e\x7e\x28\x9b\xca\xa2\xee\x75\x6d\x67\xeb\xbd\x4d\xec\xe3\xc7\xd9\xee\x71\x7c\x6f\x20\x12\x12\xd1\x40\x04\x0b\x80\x96\x5d\xc7\xfb\xd9\xf7\x00\x20\x48\x50\xa4\x64\xca\x8f\x3a\x37\x6d\xda\xc4\x26\x08\x0c\x66\x06\xbf\x19\xcc\x0c\x41\x0e\x02\x31\x25\xc3\x6f\xbf\xf9\xf6\x9b\x41\x80\xa0\x3f\xfc\xf6\x9b\xda\x60\x8a\x04\x04\x5e\x00\x19\x47\xc2\x75\x62\x31\x6e\xbd\x71\xd4\x0d\xee\x31\x1c\x09\xc0\x99\xe7\x3a\x81\x10\x11\xef\x77\x3a\x1e\xf5\x51\xfb\x97\x5f\x63\xc4\xae\xdb\x1e\x9d\x76\xf4\xaf\xad\xcd\xf6\x46\x7b\xbd\x3d\xc5\x61\xfb\x17\xee\x0c\x07\x1d\x3d\x74\x21\x95\x98\xa3\xf6\x98\x86\x02\xce\x10\xa7\x53\xa4\x28\x31\x44\x10\xe4\x88\x77\x2e\x7b\xed\x6e\x7b\xab\xf3\x0b\xef\x40\x42\xe6\xe9\x29\x8a\xe2\x9a\x20\x49\xbb\xf6\x77\x3c\x8d\x28\x13\x20\x66\xa4\x6e\x68\x4b\xba\xbc\x3d\xa1\x74\x42\x10\x8c\x30\x57\xc4\x3d\xce\xff\x36\x86\x53\x4c\xae\xdd\x63\x3a\xa2\x82\xf6\x37\xbb\xdd\xc6\xb6\xa2\x57\xfb\xbb\x1c\xd2\x1a\x43\x0f\x81\x1b\x79\x5d\x4b\xae\x65\xf7\x3e\x78\xb5\xc3\x2e\xe9\xab\x6d\x75\x83\x33\xaf\xaf\x26\x93\x6d\x3f\x50\xe2\x1f\x08\x48\xb0\xd7\x16\x62\xdc\x50\x3d\x6e\x35\xc5\xb6\xc0\x82\x18\x6a\x11\xe5\x58\x60\x1a\xf6\x01\x43\x04\x0a\x7c\x89\xb6\x17\xcd\xd2\x04\x1c\x31\x3c\xd6\xf7\x05\xba\x12\x2d\x48\xf0\x24\xec\x03\x0f\x85\x02\x31\x6b\x1c\xc7\xbf\xa1\x3e\xd8\xec\x75\xff\xa2\x1b\x23\xe8\xfb\x38\x9c\xf4\xc1\xc6\x5f\x40\x17\x74\x75\xa3\x47\x09\x65\x7d\x30\x0b\xb0\x40\x39\xf6\x08\x9d\xe0\xb0\x15\xc1\x89\xe1\x71\x86\x7d\x11\xf4\xc1\xe6\xeb\x6e\x74\x35\x47\x70\xd3\x22\x38\x85\x6c\x82\xc3\x3e\x80\xb1\xa0\x39\x82\x6a\x8a\xd6\x88\x5e\x21\xff\x0e\xa9\x7f\x6b\xe1\xd0\x47\x57\x7d\xb0\xae\xaf\x47\xd0\xfb\x34\x61\x34\x0e\xfd\x3e\x78\xf1\x56\xfd\x31\x73\x5d\xb5\x8a\x6c\x19\x0e\xba\x8a\x07\xd0\x2b\xf2\xbb\xfe\xc6\x34\x2d\x52\xe0\x88\x5e\xb5\x78\x00\x7d\x3a\x93\x74\xba\x60\xa3\x1b\x5d\x81\x2e\x60\x93\x11\xac\x77\x9b\x20\xf9\xbf\xbd\xd1\x68\x82\x2e\xe8\x45\x57\xea\x6f\xc9\xfd\xad\xfc\x9a\x4f\x18\xba\x6e\x8d\x31\x22\x7e\x19\x8c\x1c\x8d\x3b\xa7\x09\x38\x0c\x79\xcb\x5a\x67\x1a\x0b\x82\x43\xd4\x37\x4a\xce\x69\x64\xbc\x21\xff\x33\x7c\x33\x1f\xb1\xb4\x5f\xa2\x9c\xf5\xae\xc1\x40\xa6\x9b\x2e\x58\xef\x15\x15\x93\xea\x4a\x29\x00\xff\xa6\x5a\x35\x51\xb9\x74\x45\xa5\x11\x34\x16\x05\xcc\xad\x6f\x24\x64\x8c\xdc\x21\x9c\xa2\xdf\x53\xee\x47\x17\x73\x81\x71\x15\x04\x45\x8c\x51\xf6\x0e\x71\x9e\xd9\xcd\x3d\x45\x2d\x2c\x5d\x62\xa9\x2f\x76\x7e\xe8\x76\xbb\xdd\x22\x2b\x5b\x79\x56\xc6\x94\x4d\x9f\xdb\xce\x36\xba\xcf\x63\x67\x38\x8c\x62\xf1\x44\x48\x2b\xac\xcb\x8a\xd0\xeb\x55\x81\xde\x92\x75\x55\xb2\xf5\x71\x78\x09\x09\x36\xd6\x34\xaf\x45\x4b\x49\x1b\xbd\x5e\xa6\xa6\xef\x1b\x16\x21\x8e\x08\xf2\xbe\x56\x2d\x8d\x62\x21\x68\xb8\x9a\x70\x0a\xa3\x82\xc1\x90\x4b\xdb\xe9\x83\x38\x8a\x10\xf3\x20\x47\x15\x64\xdf\xda\xdd\x79\xdb\x5b\x64\xb9\x73\xb2\x97\x48\x69\x6c\xdb\x36\xba\x32\xe9\x6a\xad\x19\x1a\x7d\xc2\x09\x9b\x89\x51\x43\x42\x40\xb7\xbd\x09\x50\xca\xea\xf2\xbb\x5e\xcc\xb8\x9c\x2d\xa2\x38\x35\x44\xe3\xbf\xf8\x14\x12\xd2\xfa\x12\xb4\xf7\x26\x43\x40\x09\x72\x5a\x72\xdf\xe9\x83\x8d\x02\x72\xde\x7c\x31\x2a\xad\xf9\x98\x47\x04\x5e\xf7\x01\x0e\xa5\xe8\xad\x11\xa1\xde\xa7\xa5\xee\xd0\x80\x97\xfa\xd7\xc6\xb2\x6d\x35\xfd\xe7\xeb\xd1\x9b\x5e\x62\x60\x9d\xbf\x82\x31\x24\x44\xde\x07\x63\xca\x00\x25\x3e\x18\x31\x3a\xe3\x88\x71\xf0\xd7\x4e\x61\xb0\x91\x52\x72\x02\x59\x6b\xc2\xa0\x8f\x51\x28\xea\x0c\x4f\x02\xd1\x34\xb4\x9b\xe0\xc5\x9b\xbd\xdd\x8d\xd7\x6f\x1b\xc5\x55\x6a\x4d\xe9\x6f\x0f\x19\x4e\x1f\x30\x78\x7e\xa4\xa0\x2a\xf0\x58\x38\xb6\x22\x66\x8d\x52\x34\x2e\xa6\x94\x8a\x40\x81\x08\x86\x02\x43\x82\x21\x47\x7e\xd2\x51\xca\x4e\xf9\x55\xa1\xe7\x84\xc1\x6b\xee\x41\x92\xc6\xce\xb5\x41\xc7\xe4\x1e\x83\x4e\x92\x3f\x7d\xfb\xcd\x40\x2e\xa9\x9d\xea\x88\xeb\x08\xb9\x8e\x84\x41\xe7\x17\x78\x09\x75\xab\x4a\xa9\x6a\x97\x90\x81\xb3\x93\xfd\xe3\xf7\x3b\xef\xf6\x81\x0b\x1c\x67\xdb\xb4\x1e\xed\x9c\x9c\xfc\xcf\xe1\xf1\x5e\xbe\xf5\xf8\xf0\xec\x74\xff\x04\xb8\xe0\xfc\x22\x6b\xdb\xdf\x3d\x3c\xde\x9b\x6b\x94\x44\xd3\x26\xd3\x38\x85\x11\x70\x13\xb0\x39\xfb\xc7\xc7\x87\xc7\x60\x67\x77\x77\xff\xe4\x04\xec\xed\xbf\x3f\xd8\xdf\x73\xfa\xc0\xd9\xf1\x3c\xc4\x39\xf0\x51\x88\x91\xdf\x06\x47\x2a\x09\x03\x9e\x4c\xce\x3c\x01\xae\x69\xcc\x00\xf4\xa7\x38\xc4\x5c\x30\x28\x28\x6b\x83\xc1\x88\x75\x86\xea\x1f\xa7\x69\x93\xde\xff\xf9\xe0\xe4\xf4\xc4\x51\x6b\x12\x0b\xc4\x01\xba\xc2\x5c\xf0\x8c\x66\x80\xbc\x4f\x9a\xa2\xde\xc3\x7d\x28\xe0\x62\x72\x47\x87\xc7\xa7\xe0\x70\x77\xf7\xec\x28\xe1\xf4\x48\x66\x7a\xd4\xf3\xe2\xc8\xe6\x54\xb0\x6b\x00\x43\x2a\x02\xc4\x80\xcc\x05\x17\x13\x3c\x78\x7f\x74\x76\x2a\x09\x1d\xa8\xd9\x55\x34\x77\x5f\xe6\x76\xce\x4e\x7f\xd4\xa4\x3c\xca\x98\xdc\x6a\x63\x8e\x18\x38\xd8\x03\x94\x81\x08\x72\x3e\xa3\xcc\x5f\x3c\xfc\xfd\xe1\xa9\x56\x97\xa4\x71\x1a\x20\xc0\xa4\xc6\xe4\xf4\x4a\x1e\x41\x81\x8f\x08\x12\x08\x60\x0e\x42\x2a\xb4\x26\xdb\x8b\xa8\x69\x40\x48\x52\xbb\x34\x0c\x91\x27\x1d\x99\xa4\xb1\x07\x05\x1c\x49\xe1\x94\xa8\x4d\x45\x1a\x0b\x00\x27\x10\x87\x80\x40\x81\x96\x2c\xa6\xa6\x69\xaf\x29\xf2\x28\xf3\x81\x0c\xf6\x25\x3a\xc6\x04\x7b\xa2\x99\xd3\xbe\xbc\x95\x23\xa8\xac\x26\x81\xe3\x38\x0e\x35\x5f\x1e\x43\x50\xa0\x7a\x20\xa6\xe4\x44\xb0\x46\x82\x4e\x89\xd6\x31\x83\x13\xe0\x02\x9f\x7a\xf1\x14\x85\xa2\xad\x7b\xee\x25\x97\x6f\x19\x9c\xc8\x9f\xf5\x86\xe6\xb3\x26\xd0\x34\x2a\x76\xdf\x27\x48\xf5\x7a\xe5\xe3\xcb\x57\x0d\xe3\x90\xa7\x51\x1b\x87\x21\x62\x3f\x9e\xbe\xfb\x09\xb8\x20\x99\x3c\xd9\x93\x02\x4c\x10\xa8\xab\x4e\x63\xcc\xb8\xd8\x0d\x30\xf1\x0d\x63\x35\xc9\x55\x1b\x46\x11\x0a\x7d\x75\xa3\xd0\x51\x53\xb9\x55\xff\x32\x24\x62\x16\x2a\x49\x72\xe1\xba\x11\x5e\xaf\xaa\x32\x8f\xba\x0a\xcf\x6d\xf9\x35\x06\xdc\xc4\xe6\xcf\xd5\x7d\x6d\xe0\xea\x76\x04\x19\x9c\x72\xe0\x82\x8f\x12\x6a\x52\xd9\xee\xda\x8d\xf1\x25\xb7\x2f\x0d\xe8\xdc\xb5\x1b\xe3\x4a\x6e\x5f\x7a\x44\x7a\x54\x77\xed\x46\xd1\x6e\xef\xaa\xcb\xdb\x97\x9c\xc6\xcc\x43\x38\x4a\x6f\x9c\xa8\x86\x83\xa3\xdb\x97\x6a\x8b\x0b\x21\xb1\x6e\xee\x21\x2e\x70\x08\xa5\x00\xb2\x07\xba\xd2\x3d\xa4\xb1\xa5\x7d\xf6\x93\x46\x69\xa3\x19\x91\x5c\x97\x83\xd0\xea\xf2\x31\x93\xeb\x2a\x60\xc0\x05\x21\x9a\x81\x9f\xdf\xfd\xf4\xa3\x10\xd1\x31\xfa\x35\x46\x5c\xd4\x1b\x59\x1f\x1f\x91\x7d\x89\xe1\x64\x75\xed\x65\x9f\x20\x91\xb4\xfe\x70\x7d\xe0\xd7\x1d\x4b\xc5\x6a\x88\x63\xaa\x3c\xb5\xab\x80\xb5\x69\x84\xc2\xfa\xab\xa3\xc3\x93\xd3\x57\x4d\xe0\x74\x74\x67\xa7\x29\xb7\x5b\x8e\x92\x09\x65\x3f\x8e\x44\xc2\xc6\x8f\x08\xfa\x88\xd5\xa5\x69\x09\x14\x8a\x96\x74\xee\x4e\x13\x38\x30\x8a\x08\xf6\x94\x56\x3a\x57\xad\xd9\x6c\xd6\x92\xa1\x51\x2b\x66\x04\x85\x1e\xf5\x91\x3f\x37\x71\xc8\x10\xf4\xaf\xb9\x80\x02\x79\x01\x0c\x27\x72\xa5\x53\x5c\xd4\x53\xb8\xe1\x31\xa8\xcb\xfe\xaa\xf7\x89\xec\x0d\x5c\x17\x6c\x81\x97\x2f\x81\xe2\x4b\x40\x11\x73\xd9\xb4\xd1\xed\xa6\x83\x6a\x0b\xb5\x01\x7d\xdf\x56\x85\x94\x6b\x47\x08\x86\x47\x12\x81\x4e\x80\x7d\x1f\x85\x52\x1a\xc1\x62\x94\x31\x6c\xb3\xc1\x23\x1a\x72\x74\x8a\xae\x84\x9c\xd5\x39\xfc\x87\x93\x4d\x5b\x9b\x5b\x97\xbb\xc9\xeb\x61\xda\x4c\xcc\xe5\x6d\x6e\xd6\xff\x28\xce\x8b\x43\xb9\x93\x35\x96\xcc\xcb\xd0\x94\x5e\xa2\xe2\xd4\xd9\x94\xf3\x23\x6c\x77\xa0\xf7\x2b\xe3\x72\xc7\x10\x93\xf9\xfd\x45\x7a\xcc\x9c\x6b\xbb\x53\x94\x95\x39\x5c\xc6\xe0\x14\x46\xe7\xf3\x5a\x49\x3c\x43\xe2\x78\x6e\x2d\xa8\x71\x14\xfa\x75\xed\x2d\x12\xea\x0c\x8d\x19\xe2\x41\x3d\xbd\xd4\x5e\x4a\x62\x7e\x89\x9b\x3a\xe3\x88\x15\xbd\x94\xda\xe7\x5c\x1d\x6f\x3c\x8a\x8f\x92\x3d\xdd\xb5\x1b\xf9\xe3\xc9\xdd\x82\x14\xa9\xaa\x57\x90\x7d\x9f\xdd\x33\x3c\xa5\x63\xb0\x94\xf1\x85\xfb\x85\x95\x8d\xe9\xde\x8e\x24\xe7\x16\xa4\x82\x9e\xd6\x2b\x3c\xa7\x91\xf3\x80\xce\x74\x24\xa2\xf6\xe7\x26\x28\x98\xba\x0c\x93\x80\xab\x39\xf8\x38\x10\x6c\x98\x48\x31\x10\x3e\xf0\x08\xe4\xdc\x75\xb2\x9a\xb7\x63\xee\xd6\x0a\x41\x05\x68\x0d\x41\x69\xa0\xd0\x2f\x0f\x31\xfa\xa5\x21\x83\x99\xbc\x23\x7c\x8b\x91\xe1\xc0\xc7\x97\x86\x1b\xbb\xc0\xe1\x00\x1a\x7a\x04\x7b\x9f\x5c\x3b\x24\xa8\xaf\xdd\x28\x31\x6f\x1b\xdb\x19\xc3\x03\x6c\x08\x8c\x21\x07\x63\xd8\x12\x0c\xf2\xa0\x05\x89\x70\x86\x83\x0e\x1e\x0e\x3a\x3e\xbe\x1c\xde\x7f\x7e\xb9\x09\xab\xe8\xf9\xce\xd9\x99\x9c\x3d\x40\x90\xdd\x31\xf3\xa0\x23\xd8\xf0\xa3\xfa\x7d\xad\xee\xbc\x50\xda\xe2\x4e\x23\x89\x53\x55\x70\xdd\x58\xbc\xe8\x9a\x17\xa6\x7e\xdc\x6f\xd9\xb3\x92\x7f\x6e\xd9\x15\xc5\xb6\xa6\xff\x1e\x4e\x51\xc9\x9a\x81\x0a\xe8\xd1\x64\x72\x48\x49\x20\xa4\xef\x14\xd1\xa2\xdb\x1f\x17\x2e\xd5\x56\xec\x29\xf0\x12\xf3\x95\x66\x8f\x48\xcc\x5b\x1e\x66\x1e\x41\xab\xa1\x46\xcd\x51\x19\x36\x2a\x1c\x90\xbb\xf4\xc3\x3d\x45\xb2\xdb\xdf\x7b\x6d\x14\x2b\x8f\xb9\x32\x73\x9a\x91\xdc\x55\xd2\x4b\x0e\x2a\xc5\x84\x4e\xa7\xcc\xae\xa9\xd8\x2c\x09\x97\x74\xd7\x24\x60\x2a\xda\xd1\xcb\xea\xe1\x54\x3e\x81\x2b\x33\x98\x42\x06\x57\x62\x6f\x85\x14\xae\xc4\xc2\x9e\x3c\x87\x53\x73\x56\x0d\xd7\x74\xef\x67\x0f\xd8\x1e\x9c\xca\xfd\xc1\x32\x2f\x6d\x22\x7f\xa6\x5e\x06\x3e\x26\x61\xaf\x4b\x08\xda\xce\x24\xb3\x6b\x89\x36\xca\xa6\xed\xac\xa5\x7d\x09\x49\x8c\x32\x53\xb3\x0d\xdc\xf4\xb6\xdb\xe6\xfb\xdb\xc6\x3e\x4f\xbd\xac\xbf\xaa\x27\x9a\x42\xb5\x6a\x95\xf8\xb1\xf8\x21\x28\x9c\x88\x40\xc2\x36\x43\xb6\x1e\xf4\x9d\x0b\x5e\x19\x2f\x02\x0e\x8e\x00\xe6\x80\xa1\x5f\x63\xcc\x90\xff\x21\x7c\x65\xd7\xd2\x6c\x92\x8a\x89\x8a\x44\x55\x99\x78\x39\x59\x35\x2a\xa5\x01\x09\x62\x22\x69\x4b\x56\xb6\xb0\x4c\xe9\x2a\xdf\x37\xcb\xcd\xb9\xe5\xec\xa2\xe0\x8c\xed\xcb\x82\x17\xb6\x2f\x57\x75\xbf\xd0\xf7\xab\xba\xdf\xb9\xa2\xd1\x12\xd7\x0b\xfd\xaf\xc0\xe3\xae\x50\x4a\x7c\xec\x3c\x79\x6e\x51\x7e\x37\x2f\x3e\x3f\xef\xdd\x5e\x7c\x7e\x44\x49\xfd\x0c\xfa\xfe\xbc\x07\xd7\x4f\x55\x44\x80\x16\x3d\x54\xb9\x53\xa4\x95\x39\x5d\xc6\xe8\xef\xed\xcc\xb3\x40\x7e\x85\xd8\xf0\xf1\xdc\xcc\x97\x14\xfd\xfd\xe9\x7e\xbe\x7e\xf7\xb3\xb2\xad\x2e\x14\x7d\x0a\xa3\x7f\x65\x50\x76\x1a\x3a\x00\x01\x2e\x28\x83\xf4\x4a\xd4\x24\xa2\x8b\xf4\x6c\x64\x6f\x3f\xdc\x9d\xe6\x9c\xe3\x8e\x27\xf0\xa5\x8c\x6d\x2f\x29\xc3\xc2\x3c\xea\x7d\x9c\x50\xf7\x4b\x72\x76\x59\x95\xab\xf2\x93\xcd\xb4\x83\x76\x82\xc9\xa1\x88\x88\xd1\x69\x24\xea\x4e\xa2\x1a\x75\x82\x06\x40\xfd\xe4\x79\x4c\x99\x7e\x48\x9f\x57\x67\x1f\x48\x98\x4e\xaf\x95\xb1\x18\x7c\xc9\xe5\xb3\x09\xbb\x20\x8c\x09\x49\x97\xcd\xd6\x69\x79\x74\x97\x4b\xca\x33\x4a\xf7\xcf\xc6\xff\x7d\x1f\xa7\x56\x4c\xc5\x53\x0c\xfc\x5b\xbb\xe5\x3f\x66\x26\xbe\x28\x86\xfb\x83\x66\xe1\xaa\xc6\x38\x9f\x84\x87\x68\x66\x6c\xdf\xe4\xc9\x56\xd3\x7c\x9a\x1c\xa2\x99\x71\x0a\x56\xef\xf4\x74\x50\xa5\xa4\xda\x26\xbf\x3c\x01\x3e\x33\x8c\x59\xb9\x6f\x21\xf3\xb5\xe7\x5f\x4e\xee\xc8\x70\x5e\x4a\xee\x39\x53\x69\x4b\x25\xee\xda\x8d\x75\xa5\x6e\x59\x43\xac\xab\x07\x87\xab\x15\x1f\xa4\x2e\xf7\x8d\x5f\xf7\x23\xe5\xc2\x23\xf6\xaf\x2d\x5c\x7d\xa4\x78\xd0\xf7\x15\xf8\xd4\x41\x8a\xf9\x7c\x39\x39\x76\xf9\x98\x19\xf3\x73\xfa\x51\x42\x27\x34\x16\x6f\x29\x9b\xe6\x5c\x69\xf1\x04\x6c\xad\x78\xfc\x75\x31\xd0\x54\x00\xd4\x1a\x13\x0a\x45\x15\x90\x2d\x25\xa5\xa2\x95\xc7\xa1\xa5\x9e\x1a\x3d\x1e\x29\xa9\xb0\x87\xd2\x99\xc2\xe8\x31\xc8\xe8\xd7\x05\x13\x4a\x77\x98\xc9\x32\x22\x34\x16\xad\x15\xf8\xb9\x13\x5e\x63\x24\xbc\x40\x1f\x36\xae\xdb\xdb\xf4\x3d\x76\x95\x8f\x76\x3a\x72\xd7\x0e\x51\xe2\xdd\x09\xe6\xc2\x76\xed\x8f\xee\xdc\x17\xda\xa4\x47\x43\x4e\x09\x6a\x13\x3a\x29\xf8\x26\x6b\x47\xd3\x67\x04\x80\x0b\xfe\xfb\xe4\xf0\x7d\x3b\x82\x8c\xa3\xd2\xee\x6a\x40\x7a\xec\x5c\x8f\xb2\x75\x13\x52\x1f\x2d\x4b\x18\xcc\x59\x84\xdc\xe1\x5a\x39\xa8\x1d\x40\xae\x4e\xcc\xbe\xa7\xbe\x5c\xaf\xd4\x5f\xaa\x9b\x1a\x55\xfa\x84\xad\x6a\x20\xb0\x70\xc0\x56\xfd\x94\x19\x60\x5d\x3d\xb4\x00\x2e\xe8\x6e\x03\x0c\x06\x09\x97\x49\x40\xb3\x0d\xf0\x77\xdf\xa5\xc4\xe7\x4e\xb9\xf0\x73\x7c\xd1\x04\xd8\x3e\xb4\x5b\x0e\x2a\xfd\x6c\xfc\x4b\x40\x55\xc2\xca\xf3\x80\x2b\x2b\x55\x56\x86\x4e\xfa\x76\x42\x32\x6e\x25\xf0\x98\x23\x09\x4f\x85\x9e\xac\x26\x70\x92\x9d\xf0\x2e\x43\x94\x66\xa4\x14\x52\xb5\xc2\x21\x9a\x3c\xaa\xb2\x9d\x74\x01\xb8\x64\x88\xf4\x45\x40\x4b\x31\x92\x02\xeb\x0b\x74\x5a\x6a\x4b\xad\x08\x3c\xf3\x06\x8c\x1a\xb3\x0a\xe8\x92\xd3\x1e\x4f\x09\x39\xc5\x5b\x09\xe0\x5c\x09\xb7\x81\x62\x60\x31\xd4\xd2\x83\x37\xd5\x61\x36\x41\xe2\x88\x32\xf1\x28\x20\xbb\x3f\xc4\x22\xc9\xc2\x13\xe6\x3b\xe5\xe8\xaa\x84\xaf\x4c\x29\x92\xc7\x8a\x08\xab\x02\x26\x49\xaf\xa5\xdf\xbe\x75\x8c\x45\xfd\x3e\x9b\xa0\x92\xa4\x14\x45\xd6\xcb\x35\xc9\xdb\x37\x1f\x07\x34\x52\x40\x51\x35\x08\xd7\x59\xbb\x51\xa3\xcf\xf1\xc5\xad\x33\xb4\x2e\x06\x1d\xdd\x6f\xf8\xb1\x61\xb1\x69\xbf\x0d\x23\xe9\x2e\xdb\x4d\xd3\x8c\x21\x61\x26\x43\xa6\x1e\x94\x0b\xe1\xec\x26\xb3\x01\x5b\x6d\x89\xdf\x5c\x94\x69\xe0\xb0\x90\x68\x14\xea\x2a\x96\x5f\xb1\x8b\x38\x8b\x2a\x38\xf3\xe5\x9b\xb2\xda\x8d\xb4\xef\x07\x94\x68\x4a\x8e\x27\x3c\xa0\x42\xf3\x68\x87\x1d\xd2\x83\x7d\x22\xba\xb3\x58\x12\x33\x22\xf5\xdb\x81\xb1\x08\x9c\xb2\xd3\x6d\x96\xbf\xc9\x6a\x35\x96\xbf\x99\xaf\xd2\xc8\x59\xb5\x33\x71\xa4\x33\x71\x9a\x72\x8a\x26\x90\xb9\x41\xc3\xea\x71\x97\x1f\xd1\xe2\x55\x77\x26\x7a\xda\x15\xaa\xca\x6a\x40\x49\xf5\x44\x73\x57\x5e\x3e\xe1\x33\x2c\xbc\x20\x1d\x6b\xb9\x99\xac\x7e\xe0\x41\x8e\x80\xb3\xb3\xf7\xee\xe0\xbd\xd3\x37\x8d\x55\x13\xcf\xbb\xeb\xc0\x95\xf2\xce\xc5\x64\x12\x22\x79\x66\x0f\xff\x61\x71\x6a\xe5\xf7\x66\xbd\xb3\xc9\xad\x4c\xdf\xac\x7b\x05\xce\x72\x29\x68\xc5\xc2\x4f\xd5\x0a\xc2\x43\x34\x36\x57\x40\x78\x08\xa9\x7c\x82\xfc\x10\x4a\x59\xda\x5f\x81\x8a\xbd\x55\x3a\x3b\xb1\x08\x00\x8f\xd5\x3b\xc1\x76\xa7\xb9\xc2\x8f\xfc\x33\x62\x08\x7e\xda\xce\xa3\xc0\x7e\x39\x76\xc9\x04\xba\xea\x55\x49\x14\xe9\x53\xde\x42\x4c\x1e\xa8\x11\x8b\xcc\x7c\x3d\xac\x60\x88\x17\x8b\x64\xf4\xd1\x18\xc6\x44\x2c\x93\x4c\x7b\xe7\x38\xe4\x11\xf2\xf0\x18\x4b\x21\x9f\x5b\x46\xe7\x2c\xe3\x46\xf3\xe7\x2c\x10\xf0\xb6\x10\x57\x1a\x3f\x5b\xac\x03\x96\x56\x66\xe6\xbe\xd3\x65\x1d\xe7\x56\xdf\xc3\xd2\xe7\xb4\x4f\xe3\x30\x44\xe4\x07\x04\xb9\x50\x43\xf4\x89\xec\xb9\xfe\xd9\x07\xaa\x1c\x7d\xb3\x36\x50\xdf\xbc\x81\xca\x0d\xbb\xce\x0b\x07\xd0\x90\xc7\xa3\x29\x16\xae\x93\x70\x93\x85\x00\x22\xc0\xbc\xb1\xed\x00\xec\xbb\xb6\xe3\xd0\x87\xbd\x07\x04\x8e\x10\x01\x5a\xb1\xae\xf6\x19\x66\x5a\xfb\x0b\x3f\x7a\x78\xaa\x54\x19\x00\x0c\x07\x1d\x35\x38\x21\xa4\x2b\xb4\xd9\xdb\xfe\x8e\x7a\x46\xad\xce\xef\x2b\x9f\xe7\x80\x88\x40\x0f\x05\x94\xf8\x88\xb9\xce\x59\xda\xdc\x29\x21\x60\x7c\xa1\x21\x92\x5d\xe7\x88\x1c\xa5\xcd\x86\x48\xf2\x3d\x0f\x4d\x45\xab\xc4\x19\x2a\xa9\x07\x1d\x7d\x2f\xe9\x18\x0d\x8f\xe8\x0c\x31\xe4\x83\xd1\x35\x18\x40\x10\x30\x34\xce\x3e\xc9\x36\xc1\x22\x88\x47\xea\x63\x69\x23\x18\xc4\xfc\x12\x91\x8e\xb5\x54\xce\xd0\xba\x18\x74\xe0\x70\xd0\x89\x14\xdd\x41\x47\xea\xd6\x2c\x92\xb5\x84\xd6\x27\xc1\xb4\x2a\x6d\xa7\x9b\x57\xbf\x66\x70\x37\x66\x0c\x85\x02\xe8\x68\xb0\x9f\x48\xc7\x52\x41\x05\x1c\x25\x9f\x7d\x93\x17\xea\x3b\x1a\x29\x59\xee\x64\x27\xf9\xcd\xe7\x18\xd4\x45\x36\x66\xb5\x75\x2f\x1c\xbe\x29\xac\xbf\x85\xdc\x0a\xe0\x4c\xcf\xf5\xda\xd8\x34\x7e\xba\x4c\x1b\x2b\xc2\x34\x77\x66\xaa\x12\x56\x13\x06\xec\x73\x36\x09\xf4\xec\x96\x1c\xf8\xac\x83\xbb\x8e\x56\xb7\xb2\xeb\xdc\x0d\xb0\xe3\xfb\x0c\x71\xde\x07\x3f\x83\x76\xf6\x37\xe9\x1f\x41\x21\x7b\xba\x4e\xbd\xfe\xcf\xcf\x1f\xda\x8d\x7a\x7d\xa3\x77\xde\x6d\xf5\x2e\x1a\x9f\xeb\x1b\xe7\xdd\xd6\xd6\xc5\x07\xbf\xf1\xb9\xbe\xfe\xc1\x57\x3f\xcf\xd7\x5b\xdf\x5f\xfc\xed\x83\xdf\x68\x34\x6e\xb6\x6e\xd7\xca\x2d\xa7\x5c\x1c\x75\xd0\x67\x4e\x20\xdd\x56\x2e\x92\xcc\x4a\x16\x08\xa5\x4e\x14\x33\x19\x0a\xf6\x41\x17\xfc\x1f\x78\xdd\xeb\x6d\xf6\xe6\x05\xfa\x67\xfd\xbc\xdb\xfa\xfe\xe2\x66\xbd\xb9\x75\xfb\xf9\x7c\xbd\xd5\xbb\xd0\xd7\x5b\xb7\x9f\x5f\x2b\xc1\xf4\xe5\xe6\xed\xe7\xd7\x3d\xeb\x7a\x43\x5e\xcb\x86\x0d\xdd\x20\xaf\x36\xb5\x46\x32\x71\x4f\xf4\xd7\x9c\xcc\x51\x11\x95\xea\x25\xf6\x91\x7c\xe8\x49\xcb\x69\x9f\x25\x51\x20\x98\x5a\x18\x93\x0a\xb2\xb3\x53\x63\x22\xfa\x72\x99\x33\xd9\xf1\xfd\x9c\x2b\x59\xc1\xe4\xed\xe8\xa8\x0c\xe5\xff\x4b\x63\x06\xde\x9a\xb3\x3c\xab\x5a\x7e\x52\x2d\x7c\x1a\xd3\xb7\x8e\xbc\x3c\xdc\xf6\x55\x0d\xc7\x36\xfd\x34\xd2\x7e\x14\xdb\xcf\x1e\xba\x16\x4d\xdf\x3c\x6f\xb4\xfc\xb7\x8a\xca\xfb\xcb\x2c\x29\x44\xb3\x7f\x65\xdb\x98\x46\x97\xf5\xa4\x7d\xd1\xce\xb6\xdc\x3c\x25\xd1\xf9\x6d\xce\x7a\x46\xbf\x68\xa7\x5b\xba\xd1\xdd\x1f\x9b\x76\xd6\xb4\x6c\x3b\x52\x55\x07\x80\x43\x5b\x81\x29\x44\xef\x40\xa8\x2e\x2d\x3e\x09\x3e\x97\x2c\xf9\x6a\xe8\xb4\x9e\xd1\x1a\x80\xe6\x19\x49\xa2\xa8\x34\x37\x59\xb6\x1c\x3f\x1d\xfe\xd7\xe1\xd9\xe9\xc2\x15\x31\x9c\x0d\x3a\x89\x26\xd4\x87\x94\xc4\x94\x0c\xff\x3f\x00\x00\xff\xff\x78\x8f\xc2\x2d\x99\x56\x00\x00")

func htmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlIndexHtml,
		"html/index.html",
	)
}

func htmlIndexHtml() (*asset, error) {
	bytes, err := htmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/index.html", size: 22169, mode: os.FileMode(438), modTime: time.Unix(1589151347, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"html/index.html": htmlIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"html": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{htmlIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

