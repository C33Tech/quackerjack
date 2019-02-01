// Code generated by go-bindata.
// sources:
// static/gui/index.html
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

var _staticGuiIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\xff\x6f\xdb\xb8\x92\xff\x3d\x7f\xc5\x94\xb9\xae\x24\xd8\x92\x9c\x26\x01\xde\x3a\x96\xf7\xda\x6c\xdf\xbe\xde\xeb\x6d\x73\x49\xfa\x70\x8b\x9c\xaf\xa0\x25\xda\x62\x22\x91\x5a\x92\x92\xed\x4d\xf2\xbf\x1f\x48\x7d\xb1\x24\x3b\x4d\x7a\xd8\x7b\xd8\x03\xb6\x40\x63\x72\x38\x33\xfc\x70\x66\x38\x1c\xd2\x9e\xbc\xfa\xf1\xd3\xf9\xf5\x2f\x17\xef\x21\x56\x69\x32\x3d\x98\xe8\x0f\x48\x30\x5b\x06\x88\x30\xa4\x09\x04\x47\xd3\x03\x80\x89\xa2\x2a\x21\xd3\x2b\x1e\x52\x9c\xc0\x39\x4f\x53\xc2\x14\xbc\x65\x38\xd9\xfc\x46\xc4\xc4\x2f\x87\x0f\x34\x67\x4a\x14\x86\x30\xc6\x42\x12\x15\xa0\x5c\x2d\xdc\xbf\xa0\x69\x33\x10\x2b\x95\xb9\xe4\xd7\x9c\x16\x01\xfa\x4f\xf7\xf3\x5b\xf7\x9c\xa7\x19\x56\x74\x9e\x10\x04\x21\x67\x8a\x30\x15\xa0\x0f\xef\x03\x12\x2d\x49\x4b\x8e\xe1\x94\x04\xa8\xa0\x64\x95\x71\xa1\x5a\xac\x2b\x1a\xa9\x38\x88\x48\x41\x43\xe2\x9a\xce\x10\x28\xa3\x8a\xe2\xc4\x95\x21\x4e\x48\x70\x84\x4a\x60\x09\x65\x77\x20\x48\x12\x20\xa9\x36\x09\x91\x31\x21\x0a\x41\x2c\xc8\x22\x40\x1a\x96\x1c\xfb\x7e\x8a\xd7\x61\xc4\xbc\x39\xe7\x4a\x2a\x81\x33\xdd\x09\x79\xea\x37\x04\xff\xd8\x3b\xf6\xde\xf8\xa1\x94\x5b\x9a\x97\x52\xe6\x85\x52\x96\x70\x8d\x72\x50\x9b\x8c\x04\x48\x91\xb5\xf2\xeb\x11\x00\x6f\x4e\x97\x2c\x4f\x25\xdc\x43\x8a\xc5\x92\x32\x77\xce\x95\xe2\xe9\x18\xde\x8c\xb2\xf5\x19\x2c\x38\x53\xae\xa4\xbf\x91\x31\x9c\x18\x82\x96\x77\x71\x42\x97\x6c\x0c\x21\x61\x8a\x88\x33\x78\xec\xaa\x92\x19\x66\x5e\x44\x64\x08\xf7\x6d\xf9\xa3\x13\x2d\x1f\x51\x99\x25\x78\x33\x86\x79\xc2\xc3\xbb\x5a\xf6\x30\xe1\x38\xa2\x6c\x09\xf7\x5b\x06\xc6\x19\x39\x03\x63\xbf\x31\xfc\xe5\xf4\xf5\x59\x85\x70\x0c\xc7\xa3\x6c\x0d\x38\x57\xbc\x83\xef\xf8\x49\x7c\x21\x4f\xb8\x18\xc3\xe1\xf7\xdf\x7f\xdf\x4c\x28\x88\xf6\xda\xee\x7c\xdd\x61\x2f\x13\x7c\x29\x88\x6c\xd9\x47\xf1\xac\x9e\xab\xe2\xd5\x11\x49\x04\xdc\xc3\x1c\x87\x77\x4b\xc1\x73\x16\xd5\x4b\x1e\x8d\x5e\x9f\xb5\xc9\x19\x97\x54\x51\xde\xc2\x96\xe1\x48\x2f\x5c\xf3\x6a\x95\x73\x2e\x22\x22\x2a\x27\xb8\x82\x2e\x63\xe5\x0a\x1c\xd1\x5c\xee\xe7\x48\xc8\xa2\xcf\x50\x79\x43\x2a\xc1\xef\x08\xdc\x37\xab\x5f\x2c\x16\x95\x79\x64\x8c\x23\xbe\x1a\xc3\xe1\x68\x34\x02\x6d\x4b\xfd\xff\x54\xcb\xba\x2b\x32\xbf\xa3\xca\x2d\xcd\x9a\x72\xae\x62\x03\x0e\x33\x1d\xbd\x14\x4b\x12\x35\x13\x84\xd5\x96\xbb\xaf\x20\x8d\xe1\x28\x5b\x83\xe4\x09\x8d\xe0\x30\x0c\xc3\x06\x6a\x17\x5e\x6f\xc1\xb5\x4f\x6b\x18\x47\xa7\x65\x63\x67\x9a\xf8\xa4\xf1\x41\x8b\x7b\x54\x73\x37\x6a\x77\x87\x3a\x16\xdb\x45\xd9\x9b\x27\xab\xa3\x76\x81\x53\x9a\x6c\xc6\x90\x72\xc6\x65\x86\xc3\x2a\x38\x26\xbe\xd9\x4e\xd3\x83\x89\x5f\xa6\xa2\xc9\x9c\x47\x1b\xb3\xcf\x22\x5a\x40\x98\x60\x29\x03\xa4\x93\x01\xa6\x8c\x88\x6a\x9f\x4d\x5e\xb9\x2e\x5c\x29\xac\x68\x08\x0c\x17\x73\x2c\xc0\x75\xab\x21\x86\x1b\xb1\x6a\xa8\xfc\x70\x23\xb2\xc0\x79\xa2\x2a\x15\x4f\x4c\xe0\x2e\x92\x9c\x46\x0d\x4f\x97\xab\x52\x54\x46\x68\x8b\x07\x60\x32\xcf\x95\xe2\xac\xca\x0a\x65\x07\xf5\xc4\x14\x5f\x2e\x13\xa2\x03\x28\xc1\x99\x24\x11\x82\x08\x2b\x5c\x91\x35\x84\x92\x5e\x93\xb1\x58\xea\xfc\x7a\x58\x4a\x23\xc0\x82\x62\x97\xac\x33\xcc\x22\x12\x05\x68\x81\x13\xcd\x6b\xa8\x1a\xbd\xe0\x49\x33\x55\x07\x9a\xce\x58\x19\x66\x35\x18\x29\x5c\xce\x92\x0d\x9a\x5e\x97\x70\x18\x2e\xe8\x12\xeb\x5d\x34\xf1\x35\xdf\x57\x44\x69\xc8\x99\x6b\xd4\xff\xb3\x58\x27\x7e\x69\xca\x0e\x0d\xf7\xec\x3a\x17\x98\x45\x75\x9e\x3f\x44\xfd\x33\xec\xd2\xe4\x9e\x89\x8f\x5b\x3e\xf5\x23\x5a\xf4\x5c\x4c\xa3\xc6\x7a\x3d\xfd\xb5\x63\x1a\xcf\x75\x3d\xbf\xe0\x22\xed\x49\x18\x52\xd5\xd6\x29\x05\x81\xe0\xda\xc5\x92\x60\x11\xc6\x08\x52\xa2\x62\x1e\x05\x68\x49\x54\xdf\x55\xad\x68\xd3\x5a\x5c\x9d\xe9\xb2\x1e\x13\xc0\x84\xb2\x2c\x57\xad\x33\x08\x75\x84\xaa\x78\x40\x90\x25\x38\x24\x31\x4f\x22\x22\x02\x74\xc1\xa5\x82\xcf\x97\x1f\x51\x73\xd6\x46\xfd\xd9\xbb\x76\xd9\x0d\x6c\x99\xcf\x53\xba\x9d\x6c\xae\x18\xcc\x15\xdb\xee\xac\xcb\x9c\xed\x75\x99\xaf\x61\x75\x28\x79\xd2\xb2\x59\x6d\x2b\x86\x8b\x3e\xa2\x57\xae\xdb\x25\x24\xb4\x16\xc4\xa1\xa2\x05\x41\xd3\x09\xde\xfa\xfe\x6f\x3c\x25\xda\xd5\x13\x3f\xa1\xd3\xbe\x60\x87\xf3\x5a\x60\xca\x28\x5b\xee\xe7\xae\xb3\x49\x8d\x3f\x4f\x5e\x84\xbe\x6e\x9a\x93\xa6\xbf\x94\x36\x00\x5d\x8e\x8c\x7d\x7f\x49\x55\x9c\xcf\x4d\x01\x92\xd2\x3b\xb2\x48\x36\x8c\xf9\xbf\xe6\x38\xbc\x23\xe2\x16\x87\x77\x3a\x98\x73\x11\x12\xe0\x0c\x7e\xa2\xea\x6f\xf9\x7c\x1f\xdc\x2e\xbc\xd2\x89\xda\x70\xbe\xc7\x70\xb1\x0d\xdf\xed\x9a\xda\x2c\xbd\xcc\xb7\xcd\xa3\x3e\xc3\x85\xa9\xab\x5a\x1b\xa4\xaa\x2b\x9a\x95\x5d\x92\xb2\xce\xa8\xb2\xbd\xf4\x3c\x0f\xec\x5f\xe8\x1d\x91\x43\x90\x3c\x25\xc0\x17\xa0\x62\x22\x09\x60\x41\x20\x14\xf8\xb7\xcd\x2b\x67\x4f\xfa\xad\xeb\x82\xee\xce\xda\xc3\xa0\x13\x05\xb4\x3b\xae\x54\x82\x66\x24\x82\x2a\x1a\xaa\x9d\x56\xb3\x6c\x33\x67\x81\x93\x9c\x30\xbe\x0a\xd0\xd1\x68\xd4\xa6\xa5\x94\x05\xa8\x4b\xc1\xeb\x8a\xcb\x9c\x4d\x55\x09\x5a\xd6\x1f\x5d\x84\xbd\x3c\xb2\xed\x56\xcd\xae\xf5\xca\x2a\xa8\x7b\xfe\x68\x7a\x75\x9e\x34\x19\xda\x94\x1a\xed\x23\x28\x3e\x32\x7c\x05\x8d\x08\xff\x62\x2a\x71\x9d\x2c\xe3\xa3\x36\xcb\x49\x07\x98\x49\xb0\x5a\x26\x8c\x31\x63\x24\xd9\x4a\x99\x14\xab\xe3\x69\xcb\xc3\x88\x5a\x71\x71\xd7\xe3\x69\x2d\x4b\x2b\x6f\x4d\x25\xc0\x6f\xf7\x5b\x6e\x12\x7c\x05\x55\xdd\xfa\xa4\x2b\x43\x9e\xb8\x6b\xe9\x9e\x40\xd5\xe0\x8b\x85\x24\xca\x3d\xd9\x7b\x68\x69\x78\x8a\x2b\x9c\x7c\xa9\x43\x0c\x4d\x47\xcf\x1e\x29\xba\x5a\xd6\x07\x9c\xda\x1e\x03\x72\xdf\xe1\xd2\xcb\x76\xfb\x50\x3e\x89\xaa\xc6\xf3\x25\x23\xe2\x4b\x84\x37\x2f\xc7\x55\x23\x82\x0b\x22\xe0\x47\xbc\x79\x16\x59\xa7\xdb\x8e\xac\x67\x37\xd1\x0b\xb7\x50\x1e\x86\x5a\xee\x05\x7b\xe7\x7f\xb9\x73\x46\xaf\xcf\x3a\x86\x6c\x4c\x50\x5d\x30\x23\xcd\xf2\xad\xeb\x15\x7c\xf5\xc4\x52\xb5\xf3\xd2\xc8\x3d\xee\xc6\x60\x7c\x3a\xbd\xe6\x19\xfc\x9d\x6c\x56\x5c\x44\x72\xe2\xc7\xa7\x9d\x71\x85\xe7\x09\x31\xbe\xbd\xab\x58\x9a\x2d\x59\x0e\x99\xbf\x75\xc2\xe9\x25\x02\x33\xb6\xe3\xb2\xaf\xa1\x3b\xdd\x41\x57\x57\x2c\x57\x21\x17\x64\x17\x5f\x88\x59\x81\xa5\x01\x98\x6e\xce\x63\xac\x6f\xc9\xe5\xe5\x18\x99\xcc\x04\x31\xd1\xc7\x4e\x80\x8e\x47\x23\xbd\x8b\x4b\xfe\x6f\xc2\x74\xd2\xc7\x64\x66\x93\x38\xcd\x12\xca\x96\x5f\xea\xc2\xf7\xfd\x5a\x53\x48\x6b\x6f\xf5\xa0\xd6\x99\xad\x96\xd4\x70\x9e\x0d\xe9\x4e\x73\x0b\x76\x22\x43\x41\x33\x05\x52\x84\xdb\x6b\x3c\xbe\xc5\x6b\x6f\xc9\xf9\x32\x21\x38\xa3\xd2\x9c\xa0\x9a\xe6\x27\x74\x2e\xfd\xdb\x5f\x73\x22\x36\xfe\x1b\xef\xc8\x3b\xae\x3a\xe6\x0a\x7f\x2b\x4d\x76\x33\xfa\xa6\x4f\xa9\x7e\xe9\x0b\xc1\x6d\xff\x81\xe0\x45\xda\xc3\x88\xdd\x4a\x2f\x4c\x78\x1e\x2d\x12\x2c\x48\x0f\xb9\x71\xab\x77\x2b\xfd\x23\x6f\xe4\x1d\x55\xdd\x5d\xe5\x2d\xed\xad\x77\x88\x5b\x5c\xe0\x92\x6a\xdc\xb8\xc8\x59\xa8\xeb\x7a\x58\x12\xf5\xf9\xf2\xe3\x05\x16\x38\x25\x8a\x08\x5b\x97\x7f\x0e\xdc\x1b\x83\x0b\xa2\x72\xc1\x20\x22\x21\x8f\xc8\xe7\xcb\x0f\xe7\x3c\xcd\x38\x23\x4c\xd9\x36\x23\x2b\xb8\x24\xcb\xf7\xeb\xcc\xb6\x6e\x7e\x78\xf8\x6e\x66\xc1\xc0\x94\x8e\x30\x00\x2b\xd0\x1d\xcb\xbe\xf9\xef\xef\xce\x66\x83\x1f\x1c\xfb\xbb\x87\xc3\x87\xb3\x87\x7f\x71\x2c\xc7\x23\x6b\x12\xda\x09\x0f\xcd\xa5\xc2\x2b\x6b\x5e\xe7\xe1\xe1\x66\x88\xd0\xcc\xb9\x39\x9a\x79\x82\x98\xba\xd4\xf6\xff\x6b\xe0\x2f\x87\x60\xbd\x7e\x33\xb2\x1c\xe7\xe1\x81\xe5\x49\x72\x00\xf0\x78\xd0\x06\x1f\x15\xb6\xbe\x0f\x0d\xa1\x2a\x32\x7f\xc4\x0a\xf7\xc0\x1b\x06\xf8\xc1\xdc\x9b\x60\xdc\x61\x3c\xdb\xd1\xa7\x44\xce\x42\xac\x88\x2d\x95\x18\x42\x8a\xd7\xb5\x32\xba\xd0\x24\x2f\x21\x6c\xa9\x62\x98\xb6\x87\x00\xa4\x12\x10\xe8\xbf\x9e\xcc\xe7\x3a\x03\xb0\xa5\x3d\xaa\xc4\x07\x80\x3c\xcf\x43\x67\x86\xf5\xf1\xa0\x8d\x4c\x2a\xb1\x8b\x20\x97\x44\x7c\xa4\xec\xce\xd6\x0d\x6d\xcf\xa1\x71\x62\x0b\x87\xee\x42\x10\x80\xf5\x0b\xcf\xaf\xf3\x39\xf9\x87\x3e\xf8\xad\x2d\x9a\x4a\x7b\x5d\x45\x6e\x78\xae\xf2\x79\x19\x4a\x5a\xa7\x8f\x60\x00\xb5\xf2\x0a\x16\x90\x44\x92\x8e\xee\x0f\x4c\x2a\xbc\x14\x38\xbd\xa0\xe1\xd3\xba\x69\xcd\x65\xb4\xef\x51\xdc\x59\x2f\x3a\x44\xbb\xeb\x8d\x04\x5e\x95\x97\x30\xfb\x56\x72\x56\x4f\x75\xfb\x1f\x7a\x5f\xda\xa6\x82\x36\xd7\x19\xe4\x78\x05\x4e\x0c\x8f\xf7\xf9\xf2\xa3\x73\xd6\x61\xab\xde\x90\x90\xe3\xc9\x98\xaf\xec\x6a\x94\x2e\xa0\x14\xb8\x7e\xc6\x64\x8d\x9a\x76\x11\xe5\x78\x7a\xdf\xd8\x4d\x48\xec\xd9\x05\x46\xf9\xbf\x13\x85\x75\x70\x79\xd7\x5a\xcc\x19\xc2\xe9\xc8\xa9\x00\xb4\x34\x77\x4b\xad\x4a\xf7\xb3\x2a\xcf\x4b\xa9\x52\xf3\xae\xd2\x6e\x6d\x56\x29\x45\xd5\x22\xd1\x2e\x7f\xaf\x58\x7a\x29\x0a\x53\x2b\xd5\xe9\x7c\x0f\x8c\x2a\xf3\x3b\x5e\x28\xa5\x7d\xdf\xa4\x71\xd4\x7a\x8c\xa3\x29\x5e\x12\x34\x06\x94\x8b\xc4\xb6\xd0\xe0\xf9\x39\xe3\x3c\x9d\x33\x4c\x13\x67\x80\x2c\x07\x55\x4a\x1f\x9d\x5e\xc4\xf6\x1d\xfc\x44\xdc\xfe\x3e\x0e\x3e\xc7\x99\x8e\xd9\xdf\xd7\xc5\x9f\x25\x11\x3f\xeb\x9c\xfb\x52\xf7\x36\x4b\xfc\xe7\x39\xb8\xb4\x5a\x41\xc9\xea\x1b\x75\x7e\xd4\xb7\xbd\x3f\x76\xc4\xfc\x15\x87\x64\xce\xf9\xdd\x05\x97\xea\xff\x47\xc4\x5c\xe0\x25\xf9\xa6\x88\xa9\x97\xf8\x67\xc0\xfc\x1e\x01\xf3\x0f\xca\xfe\x8f\x0e\x90\x3f\x46\x7e\xd1\xeb\xfb\x33\x52\xbe\x25\x52\x0e\xba\xb5\xc8\xce\xe5\xbf\x5a\x97\xd1\x59\x99\xe2\x6d\xb1\xbc\x20\xe2\x47\xbc\xf1\x14\xff\x2b\x5d\x93\xc8\x7e\xe3\xf4\x6a\x1a\xaf\x7d\x0d\x47\x8e\x87\x95\x12\xb6\xd5\xb9\x72\x5b\x43\x68\x2b\x3d\xe7\x05\x11\x78\x49\x2e\x88\x08\x09\x53\x4e\xef\x5d\xb8\xf3\xaf\x34\x8f\x65\xae\x88\xd6\xf8\x6b\x6a\x06\xd6\x6b\xeb\xf1\xeb\xba\xcc\xf2\xac\x3d\xf7\x76\x6b\xf0\x8c\x62\xe7\xac\x34\x9e\xef\xc3\x15\x17\x0a\x54\x4c\xa0\xba\x60\x43\x42\xa5\x32\x83\x92\x8b\xf2\x8e\x1d\xc0\xcd\xcc\x50\x16\x5c\x80\x5d\x60\x01\x6b\xa0\xac\x04\x5f\xdf\xdc\x5b\x65\x79\x25\xe6\x65\xb9\x8c\xed\x9b\xf5\xb0\xcb\x78\xb3\x9e\xcd\xba\x1e\x6c\x04\x74\xc3\xae\x4b\x54\x1b\x0f\x61\xee\xc0\x7d\x55\xc5\xce\x6f\x8e\x66\xe0\x02\xbe\x39\x9a\x3d\x3a\xa5\xdc\x06\x02\x18\xed\xc1\x55\xeb\xdb\x42\xda\x0c\x06\x3b\x01\xdb\xbc\x27\x38\x1e\xce\x32\xc2\x22\xdb\x9a\x28\x31\x9d\xa8\x68\x6a\x0d\x36\x03\xcb\x9b\xf8\x2a\xaa\xba\xb5\xc6\x9b\xf5\xec\x66\x34\x1b\x58\x60\x77\x69\x47\xb3\x81\xe5\x94\xfc\xbe\x12\x53\xab\xbb\xbe\x6d\x2e\xe3\x59\xb3\x25\x77\x73\x58\xff\x46\x5f\x27\x86\x6b\x9e\x35\x97\x7a\x54\x07\x44\x5f\x9f\xb7\xe0\xe2\x3d\x0e\xe3\xad\xf9\x8a\x21\xd0\x21\xe0\xed\x44\x7b\xa6\x6a\xaf\xbd\xf3\xf0\x60\x94\xa2\xe9\x24\x3e\xd9\x3e\x8f\x5b\x83\xe6\x92\x54\x78\x6f\x73\x15\x73\x93\xe0\x2a\xf7\xea\x3c\xed\x0c\x2c\x34\xfd\x57\x6b\xd0\x1e\x1e\x58\x13\x1f\x4f\xb5\xc1\x0a\xcf\xe4\x8e\x81\x05\x89\xc9\x21\xe6\x09\x73\x92\x4d\xf5\xc8\x79\xf9\xc3\x02\xcd\x9c\x55\xcf\x12\x96\xf3\xcc\xc9\x70\x65\xde\x3b\xf6\x18\x74\x77\xf0\x0f\x6e\x9d\x6f\xb6\x44\x19\x57\x3a\xe0\xa3\x0c\x02\xb8\x7f\x3c\x7b\x6a\x7b\x5e\x11\xa6\xa8\x86\xbb\x5d\xa9\x66\x60\x10\xf4\xc6\x75\x18\x23\x0d\x09\xcd\xce\x5a\x8c\xd9\x7e\xc6\x2a\x95\x68\xde\x8a\x39\xca\x6e\xd8\x0c\x02\xc8\x76\x20\x86\x6a\x0d\x01\x44\x3c\xcc\xb5\xb8\xb7\x24\xea\x7d\x42\x74\xf3\xdd\xe6\x43\x64\x37\xcf\x66\x8e\x1e\x31\xab\xd7\x41\xff\x26\x42\xad\x1c\xf5\x0e\x0b\x30\x5c\xdb\x75\x63\x85\xf5\xca\xab\xd9\x13\x3c\x27\x89\x1c\xc3\x0d\xba\x26\x42\xd0\x79\x42\x5e\xa1\x21\xa0\xab\x3c\xbc\x93\xba\xf1\x0e\x47\xfa\xe3\x67\xae\xe0\x27\xce\x4d\xfb\x7d\x6c\x28\x24\x57\x02\x27\xba\xf9\xe9\xef\xfa\x6f\x3d\xac\x63\x15\x3e\x28\xd3\xe4\x05\x89\xaa\xf6\xdb\x15\x91\x3c\x25\xaf\xd0\x6c\x58\x2f\x1c\x2b\x2c\x89\xd2\x93\x37\xa1\x74\xdf\x4a\xdb\x06\xda\x18\x1a\x93\x0d\x5b\x63\x5a\x76\x0c\x37\x51\x61\x47\xd9\x0d\x72\x4f\xd1\x6c\x08\x23\x67\x08\x35\xe1\xa4\x4f\x38\xee\x13\xde\xf4\x09\x47\x3d\xc2\xa8\xd7\xef\x8f\xf7\x15\xf4\x67\xe8\x43\xa8\x30\xce\x9a\x55\x3c\x56\xad\xd9\xc1\xb6\xa7\x3d\x94\x6e\xde\x61\x61\x9c\x06\x01\x30\xb2\x2a\x1d\x68\x87\x6a\xed\x78\xef\xb0\xa8\x1e\x90\xee\xcd\xaf\x85\xae\x62\xbe\xfa\x49\xd0\xe8\x23\x65\x44\x8e\xc1\x7c\x8d\x3e\x04\x41\x64\xc6\x99\xa4\x05\x19\x83\x12\x39\x19\x42\x8a\xa9\xf9\x62\xec\xad\xcc\x48\xa8\x2e\xb1\xa2\xbc\xe2\x7e\xdc\xbe\x2a\xad\x28\x8b\xf8\xca\xe3\x2c\xe1\x38\x82\xa0\x79\xf3\xb0\xeb\x2d\xa0\xd1\x15\x54\x0f\xf5\x1f\xe1\xac\x82\x46\x75\xf2\xd6\x5c\x3a\xf2\xaf\xcc\x2b\xd3\x0e\xa7\x1e\xb2\x9a\x07\x0f\xbb\xa0\xd1\x6e\x4a\x37\x5f\x39\xb7\x9f\x51\x34\x57\xbd\xbb\xfa\x42\xad\x04\x54\x7f\xa3\xe7\x78\x0b\x1c\x91\x0f\xcc\x3e\x1d\x8d\x1a\xb9\x9a\x4f\x6f\x97\x7f\xbb\xfa\xf4\xb3\x8d\x7c\x9c\xd1\x1f\x0a\x1a\x05\x68\x50\xd0\x68\xb8\x5d\xb0\x36\xa0\xd3\x09\x46\xba\xb0\x5f\x69\xaa\x47\x84\xe0\xa2\x3b\xf6\x24\x80\x4f\xb9\xd2\x08\x86\xbb\x96\xdc\xfe\xab\x8c\xae\xad\x02\x81\xf1\xdc\x59\x8f\xa3\xf5\xe2\x64\x70\x75\xc7\x1f\x3b\xfd\xc7\x83\x5d\xfa\x63\xef\xcd\xac\xf2\x4d\x37\xaf\x69\xe2\xa7\xf9\x2d\x04\xa0\x4d\xe3\x65\x58\x48\xd2\x70\xb6\x2c\x5f\xf1\xb5\xd7\xd1\x7b\x11\xd3\xa3\x7d\x93\x9b\x77\x31\xe4\x78\x31\x8d\x88\xdd\x07\x76\x00\xf5\x6f\x67\xaa\xf7\xe0\x89\x5f\xfe\x6a\x66\xe2\x9b\xdf\xf9\xfd\x4f\x00\x00\x00\xff\xff\x56\x74\xef\x99\xf7\x27\x00\x00")

func staticGuiIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticGuiIndexHtml,
		"static/gui/index.html",
	)
}

func staticGuiIndexHtml() (*asset, error) {
	bytes, err := staticGuiIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/gui/index.html", size: 10231, mode: os.FileMode(420), modTime: time.Unix(1549002733, 0)}
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
	"static/gui/index.html": staticGuiIndexHtml,
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
	"static": &bintree{nil, map[string]*bintree{
		"gui": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{staticGuiIndexHtml, map[string]*bintree{}},
		}},
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

