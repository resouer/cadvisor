// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// generated by build/assets.sh; DO NOT EDIT

// Code generated by go-bindata.
// sources:
// pages/assets/html/containers.html
// DO NOT EDIT!

package pages

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

var _pagesAssetsHtmlContainersHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x5a\x5f\x73\xdb\xb8\x11\x7f\x96\x3f\xc5\x96\xd3\x87\xeb\x8c\x49\xc5\x49\x5e\x9a\xca\x9a\xd1\xc9\x49\xa3\x9e\x63\x7b\x2c\xfb\x6e\xee\x11\x24\x21\x09\x31\x44\xf0\x00\xd0\xb2\xea\xf1\x77\xef\x02\x20\x25\xfe\x95\x2c\xdb\x93\x54\x37\x39\x8b\x04\xf6\xb7\xbf\x5d\xec\x2e\x16\xa4\x06\x7f\xf3\xfd\x23\x80\xb1\x48\xd7\x92\xcd\x17\x1a\xde\xbf\x3b\xf9\x08\xff\x16\x62\xce\x29\x4c\x92\x28\x80\x11\xe7\x70\x6d\x86\x14\x5c\x53\x45\xe5\x3d\x8d\x83\x23\x14\x39\x67\x11\x4d\x14\x8d\x21\x4b\x62\x2a\x41\x2f\x28\x8c\x52\x12\xe1\x9f\x7c\xe4\x18\x7e\xa7\x52\x31\x91\xc0\xfb\xe0\x1d\xfc\x62\x26\x78\xf9\x90\xf7\x8f\x7f\x21\xc2\x5a\x64\xb0\x24\x6b\x48\x84\x86\x4c\x51\x84\x60\x0a\x66\x0c\x15\xd3\x87\x88\xa6\x1a\x58\x02\x91\x58\xa6\x9c\x91\x24\xa2\xb0\x62\x7a\x61\xd5\xe4\x20\x01\x42\xfc\x99\x43\x88\x50\x13\x9c\x4d\x70\x7e\x8a\x57\xb3\xf2\x3c\x20\xda\xf0\x35\x9f\x85\xd6\xe9\xa7\x7e\x7f\xb5\x5a\x05\xc4\x72\x0d\x84\x9c\xf7\xb9\x9b\xa7\xfa\xe7\x93\xf1\xe7\x8b\xe9\x67\x1f\xf9\x1a\x89\xdb\x84\x53\xa5\x40\xd2\xbf\x32\x26\xd1\xd0\x70\x0d\x24\x45\x36\x11\x09\x91\x23\x27\x2b\x10\x12\xc8\x5c\x52\x1c\xd3\xc2\xb0\x5d\x49\xa6\x59\x32\x3f\x06\x25\x66\x7a\x45\x24\x45\x94\x98\x29\x2d\x59\x98\xe9\x8a\xab\x0a\x6e\x68\x71\x79\x02\x3a\x8b\x24\xe0\x8d\xa6\x30\x99\x7a\xf0\xeb\x68\x3a\x99\x1e\x23\xc6\x1f\x93\x9b\xaf\x97\xb7\x37\xf0\xc7\xe8\xfa\x7a\x74\x71\x33\xf9\x3c\x85\xcb\x6b\x18\x5f\x5e\x9c\x4d\x6e\x26\x97\x17\x78\xf5\x05\x46\x17\x7f\xc2\x6f\x93\x8b\xb3\x63\xa0\xe8\x28\x54\x43\x1f\x52\x69\xf8\x23\x49\x66\x9c\x68\xd6\x0d\x60\x4a\x69\x85\xc0\x4c\x38\x42\x2a\xa5\x11\x9b\xb1\x08\xed\x4a\xe6\x19\x99\x53\x98\x8b\x7b\x2a\x13\x34\x07\x52\x2a\x97\x4c\x99\xa5\x54\x48\x2f\x46\x14\xce\x96\x4c\x13\x6d\xef\x34\x8c\x0a\x8e\x7c\x7f\x78\x74\x34\x58\xe8\x25\x1f\xe2\xe4\xc1\x82\x92\x78\x68\x97\x60\xa0\x99\xe6\x74\x18\x8d\xe2\x7b\xa6\x50\xb3\x0f\x8f\x8f\xc1\x19\x53\x29\x27\xeb\x0b\xb2\xa4\x4f\x4f\x83\xbe\x9b\xe2\xa6\x63\x74\xc2\x39\xd1\x54\x69\x1b\x09\x18\x1b\xb1\x61\x00\x4b\x96\x20\x59\xbc\x18\x4f\xa7\x60\xb4\xd9\xd9\x9c\x25\x77\xb8\x5c\xfc\xd4\x53\x7a\x8d\x6b\xb7\xa0\x54\x7b\xb0\x90\x74\x76\xea\xa1\x9e\x6b\x21\xf4\xd3\x93\x32\xbc\xa3\x7e\x88\x17\xe8\x77\x92\xfa\x1f\x82\x13\xfc\x0f\x11\x83\x48\x29\x6f\x78\xb4\xd5\x7c\x99\x1a\x0b\x09\x37\xc6\x2d\xe9\x6b\xf5\x58\x90\x0e\x6d\x87\x20\x46\x22\x31\xc1\x8e\xb9\xd5\x20\xbc\xd3\x55\xff\x21\xf7\x64\x1a\x49\x86\x89\xb5\xb1\x44\xb9\x6b\x25\xa3\xa6\x9e\xef\x7f\x65\x54\xae\x7d\xa4\xfb\x2e\x78\x6f\x19\x7f\x47\x6d\x83\xbe\x93\x79\x06\x40\x9b\x8b\xbb\x21\xf4\x3a\xa5\xa7\x9e\xa6\x0f\xba\xff\x1d\x99\xba\xbb\x5e\x3b\xf2\xdc\xd6\x27\xff\xbb\x22\x29\xab\x41\xbe\x18\xb3\xe4\xd6\x37\x22\x19\x2d\x88\xd4\x4d\xb4\x41\xbf\xc8\x87\x41\x28\xe2\x75\xae\x20\x66\xf7\x10\x71\xa2\xd4\xa9\xb7\x61\xe2\xe2\xce\x57\x0b\xb1\x8a\x08\x56\x4d\x18\xe6\x75\x6c\x40\xea\xb1\xe1\x6d\x85\xb9\xaf\x96\xfe\xc9\x7b\x0f\x58\x7c\xea\x71\x31\x17\xde\x46\xac\x4f\x36\x5f\x2b\xfa\x0a\x91\xe1\x51\xaf\x3c\x90\x62\x19\xf0\x0d\x59\x2a\xcd\x90\xc9\xe4\x93\x61\x33\x61\xf1\x26\xca\xf5\x51\xd0\xfc\x15\xbc\x10\x0f\x25\x8a\x46\x32\x5b\x86\x4e\xfa\xf1\x51\x62\x6d\xa1\xf0\xf7\x14\x2b\x63\xa2\xc7\x1b\x33\x3f\x9d\x42\x70\x55\xbd\xa7\x9e\x9e\xac\x42\xce\x86\x25\x63\xeb\x92\xc1\x39\xe6\x0d\x1a\x3f\x6c\x19\xba\xc1\x45\x32\xec\x08\x3a\x1f\x51\x1c\x01\x9a\xc4\x06\x78\xd0\x17\x7c\xeb\x14\x4b\xdc\x5d\x3c\x3e\xb2\x19\x04\x13\xe5\x9c\xba\xc7\x57\x90\x7f\x06\x8b\x8f\x5b\x92\x41\xd0\x8f\x45\x74\x67\x3c\x76\x66\xff\xc2\xd6\x26\x47\x06\x67\xb7\xab\x76\xe4\xde\x90\xc8\x02\xa3\x15\x79\x7c\x35\x7f\xe0\x4a\xc4\x87\x13\x98\x66\x61\x54\x5e\x92\xd7\x05\xcf\x87\x61\x05\x0f\x99\x7c\x28\x47\x4e\x49\x98\xe3\x9e\xe8\xcf\xa5\xc8\xd2\x5a\xe8\xa8\x12\x80\x8d\x9b\x3a\xc3\x5e\x25\x3b\x2a\xf3\x8b\x68\x69\x2a\xf1\x99\xa6\x4b\x1b\x45\x95\xf9\xdb\x10\xaa\x45\x4f\xc9\x6b\xdd\x2e\x74\x1e\x74\x41\x30\xc5\x92\x90\xbd\x85\x03\xcf\x24\xc3\x5d\x19\x1c\x5e\xdd\x81\x19\xdf\xeb\x3f\x17\x9b\xca\x8a\x5b\xff\xd5\xf8\xb9\x9c\x73\x30\xd0\xe2\xa2\x81\x4a\xb1\x41\xc9\xb5\x18\x18\x9f\x93\x90\x72\xeb\xbb\x32\x76\xf0\x1b\x5d\x1b\xd7\x99\xe9\x43\xa8\x0f\xfe\x4e\x78\x66\x4b\x47\x3d\x31\xab\x5e\x73\xc6\x6e\xb9\xf5\x5e\x46\x6d\xaa\x85\x44\x5f\x0e\x42\x39\xcc\x09\x19\xa8\x2e\x67\xf5\xb6\xbe\xb2\xea\x1b\xbe\xea\x66\x75\xa8\xbf\x4a\xf8\x4d\x7f\x95\x07\xab\xfe\xea\x6d\xdc\x85\x4b\x9f\x71\x6b\x4d\xe1\xc9\xfc\x46\x57\xb4\xb6\xe5\xb8\xb3\x6a\xb2\x44\x17\xed\x8f\x50\xd8\x7c\xba\x43\x15\x4a\x1f\x13\xb3\x0e\xda\x05\x6b\x69\xa4\xcc\xcb\xa1\x99\x0d\xcb\xc5\x89\xcf\xac\x8c\xd9\x38\x2b\xb3\xcc\x12\xe2\xbf\xa3\x36\x8c\x36\xdb\xf0\xc4\x22\x32\x19\x51\x35\xba\x27\x8c\x9b\xbe\xfd\x0d\x72\x70\xa2\x04\xb7\xbd\x6f\x2d\xff\x9c\xca\x71\x9a\x95\x95\x75\x06\x5a\xc9\x13\x9d\xf1\x03\x24\xd2\x18\x06\x78\x4a\xc8\x35\xfa\xb6\x39\x06\x0c\x12\xca\xdd\x77\x6f\x38\xbe\xba\x75\xcb\xbf\x45\xcc\x8b\x37\xb6\xf4\x86\x0e\xd6\x3d\xec\xd6\x37\x86\xef\x56\xb9\x2b\x8f\xb0\x9f\x31\xeb\x58\xc4\x68\x2a\x59\xa2\xdd\xcd\xa6\x32\xa8\xc0\x64\x09\xdb\xc0\xa8\x32\x4c\x93\x79\x79\x11\x5b\x6c\xf9\x46\x1e\xde\xc8\x1c\x44\x02\x0b\x55\xb3\x68\x2c\xaa\x06\x6d\x35\x76\xdb\x14\x89\x57\x99\xa4\xee\x5e\x6f\x0e\x9e\xd3\xc5\xca\x9c\x88\x44\x73\x91\x8c\x86\x9a\x42\xc0\xff\x47\x0b\xdc\xe6\x26\xc9\x4c\x04\x17\xd9\xd2\xca\x15\x35\xa6\xc9\xbe\x28\x35\x9b\x6b\x67\xc4\x37\xba\x14\x72\xfd\x63\x03\xde\xe9\xdc\x11\xf3\x6e\x42\xe0\x1e\x57\x58\x98\xd7\xbb\xb7\x04\x56\xcf\x00\xf6\x5f\xba\x43\x71\x77\xd0\xe4\xf2\xb7\x78\x6b\x87\xfc\x4b\xa2\x2a\xc7\x79\xa3\x44\x69\x4b\x92\xa6\xd1\x7b\x73\xa4\xd3\xdc\x5c\xf2\x15\x86\x4e\x57\x24\x7d\xab\x22\x87\x50\xad\x65\xa1\x69\x71\x49\xeb\x0b\xac\x2e\x49\xef\xb1\xbc\x9e\x7a\xb9\x75\x95\x2e\xf4\xc5\x9b\xd9\xad\x32\xad\x51\x77\x27\x6e\x33\x2f\xcf\x3f\xb4\x64\x49\xe4\x7a\x47\x1b\x60\x66\x19\x0d\x2c\x99\x37\x1b\x81\xea\xb4\x3c\x99\x2f\xb1\xcb\xb9\x67\x74\xb5\xbb\x3d\x80\x52\x87\x90\x19\xc6\xfe\x9c\x64\x73\xea\x55\x21\xcd\x71\x7a\xd3\x32\xfc\x14\x6b\xae\xa4\xc0\x66\x43\xed\xeb\x76\xca\xe6\xa4\x85\x88\xaf\x45\xfa\x2c\x83\x3a\xfa\x8c\x1f\x68\xa6\x6d\x39\x9e\x63\x60\x8b\x35\x35\x05\x1f\x87\x37\x42\x13\x0e\x45\x1c\x7e\xb4\x91\x59\xf2\x4f\x94\x66\xe8\x19\x9c\xe2\xbb\x85\xb7\x4f\x55\xb6\x4e\xb1\x73\xcd\xb3\x2f\x03\x85\xbc\xe0\x5c\x90\x18\x46\x18\x55\x3b\xf0\x38\xce\xa9\x02\x6d\x1e\x89\x95\x99\x59\x4e\xe6\xe9\xa7\xdd\x54\xbb\xc0\x70\xdc\x37\xfb\x7f\x2b\xbf\x76\xc8\x5f\x25\x25\x77\xb1\x58\x25\x5d\x98\x0e\x2a\x2c\xa6\x75\x82\x36\x43\x63\xef\xee\xfc\x03\xc3\xa4\xd8\xa8\x7f\x50\xa4\x2c\xad\xba\xfd\xcb\x10\xca\x7e\xed\x4e\x89\x80\x14\x2b\x68\x3f\xf0\xec\x5c\xc2\xda\xb4\x66\x39\xfe\xa7\x3d\x5b\x56\x4c\x95\x62\x6e\x1e\xd0\x37\x94\x34\x7c\x92\x4f\xf4\x43\x22\xa1\x7c\xe1\xc7\xe6\xa0\x2a\xbd\xa2\x8e\xb8\x81\x85\xd0\xbe\x73\x45\x2b\x32\x54\xf7\x2a\x25\x7d\x91\x70\x9c\xfa\x55\x68\x28\x16\xcc\x1d\x92\x5b\x24\x9b\xde\x3c\x84\x2e\xc3\x56\xb3\x46\x16\xbd\x13\xbf\x84\xed\x18\xe5\x9e\x4b\xb7\xd7\x6b\xe5\xdd\x7e\xb3\xb9\x72\x1f\xbc\x72\x74\x99\x67\xbf\xb5\xea\x73\x60\x52\x5e\x50\xbd\x12\xf2\xee\xc0\xac\xec\xbd\x3e\x1d\x73\xc5\xf9\x66\x7f\x48\x22\xf6\xea\xa3\xb1\x14\xa9\x09\xfe\x66\x82\x84\x99\xd6\x62\xb3\x5e\xa1\x4e\x00\xff\xf9\x31\x9d\x91\x8c\x6b\x28\xe4\xb0\xa2\xcf\xe7\xc8\x29\x7f\xa0\xee\x84\x9c\x9f\x13\xc7\xd2\x57\x94\xd3\xc8\x1e\x01\x36\xca\x20\x26\x9a\xe4\xa2\x25\x0e\x40\x24\x23\xfe\x82\xa8\x54\xa4\x59\x7a\xea\x69\x99\xd1\xfc\x26\x7d\x40\x3b\x62\x8a\xb0\x33\xc2\x15\x6d\x09\x31\x17\x5e\xed\x8a\x8b\xb5\x6e\x8f\xaf\x4a\x60\x46\x78\xa6\x2d\xcd\xed\x15\x91\xe0\x2c\x6b\x78\x09\x8f\x48\xad\x2a\xbd\xba\x83\x31\x37\x92\xcc\x03\x29\x8c\xc5\xee\xbb\x35\xcc\x76\x97\x9c\xc6\xe1\x7a\xa7\xc7\x9a\x31\x9f\x3f\x1e\xda\x11\xb6\x87\x14\xe4\x05\xb6\xd4\xf3\x45\x9a\xe9\x66\x15\xdc\x94\xe5\x82\x5e\xb8\xd6\xd8\xe4\x34\xb6\xef\x17\xa8\xfd\x2c\xa5\xb0\x8f\x8f\x1b\x5b\x40\xa1\x8b\xda\x19\xdd\xca\x6a\xc6\xd7\x32\xf4\x8b\xfa\x69\x5b\xe6\x17\xc6\xa9\x5a\x2b\x3c\xa3\x3c\xbf\x83\x9c\x6d\x64\xdc\xde\xd7\xda\x44\x76\x23\x75\x94\xa9\x71\xa6\xb4\x58\x7e\xa3\x5a\xb2\xe8\x50\x7f\xec\x29\x56\xbd\x5d\x1e\x18\xb9\x57\xec\x26\x8e\x21\xd7\x5e\xaf\x58\xbb\x62\xa5\xd6\x4b\x59\x23\x30\x89\x2c\xce\xde\x78\xe8\xd5\x0f\x9b\x2d\x6f\x41\x7e\x5a\x68\xb4\xbc\x3b\xd9\x17\x1d\xcf\x6b\xaa\x52\x30\x7d\xb3\x6d\x6b\x3e\xd5\xeb\x05\x4b\x30\xb9\x2b\xad\x6e\xf9\x0d\x89\x1f\xbb\x37\x81\xb8\x8d\x67\x89\xf6\x5a\xf7\xef\xcd\xd6\xdd\x26\x67\xe1\x3b\xe4\xee\xcd\x43\xef\xd3\x93\x77\x35\xca\xdd\x85\xa6\x95\x61\xa5\x1b\xac\x21\xb5\x17\xc0\x17\xfa\xd0\x35\x23\x7b\xdd\x98\xb7\x11\xff\x9f\x9e\xac\xb4\x5a\x4e\x0b\xee\x3c\xbc\xa4\x26\xe4\x22\xba\xab\x7b\xa0\xb9\x3f\xd6\x7b\xf2\x37\x5c\x96\x8e\xd2\xdd\x32\x58\x1e\x2a\x0d\xec\x7e\x97\x5f\x08\x2b\x8d\xd4\xae\x90\xe4\x2f\x8f\x8f\xc1\xe6\x15\xae\x7b\xe5\x7d\x6c\x7e\xb5\x52\x3d\x7f\xdb\x5b\x8d\xe3\x96\xbd\xeb\x5e\xe1\xda\xaf\xc5\xfb\x5c\xfb\xf3\x27\xf3\x89\x25\x59\xb9\xd7\x23\x46\x4d\xf5\x4d\x4c\x3e\xa9\xfa\xd3\x01\xf7\x8b\x01\x5c\x3a\xfb\xdb\x9a\xff\x05\x00\x00\xff\xff\x13\x0e\x97\x14\xbe\x25\x00\x00")

func pagesAssetsHtmlContainersHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pagesAssetsHtmlContainersHtml,
		"pages/assets/html/containers.html",
	)
}

func pagesAssetsHtmlContainersHtml() (*asset, error) {
	bytes, err := pagesAssetsHtmlContainersHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pages/assets/html/containers.html", size: 9662, mode: os.FileMode(420), modTime: time.Unix(1469197804, 0)}
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
	"pages/assets/html/containers.html": pagesAssetsHtmlContainersHtml,
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
	"pages": {nil, map[string]*bintree{
		"assets": {nil, map[string]*bintree{
			"html": {nil, map[string]*bintree{
				"containers.html": {pagesAssetsHtmlContainersHtml, map[string]*bintree{}},
			}},
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
