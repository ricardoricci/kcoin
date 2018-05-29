// Code generated by go-bindata.
// sources:
// index.html
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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x6d\x6f\xdb\xbe\x11\x7f\xed\x7c\x8a\xab\xba\x22\x36\x6a\x49\x94\x1f\x12\x47\xb1\x5c\x74\x69\x31\x74\x40\xd7\x02\xcd\xb0\xee\x25\x2d\x9d\x24\x26\x14\xa9\x92\x67\xbb\x5e\x91\xef\x3e\x50\x92\x1f\x92\xa6\xed\xd6\x16\xf8\xff\xdf\x24\xbc\x87\xdf\xdd\x8f\xe2\xdd\x91\x9e\x3f\x79\xf5\xee\xea\xfa\xdf\xef\x5f\x43\x49\x95\x5c\x9c\xcc\xdd\x3f\x90\x5c\x15\x89\x87\xca\x5b\x9c\xf4\xe6\x25\xf2\x6c\x71\xd2\xeb\xcd\x2b\x24\x0e\x69\xc9\x8d\x45\x4a\xbc\x15\xe5\xfe\xcc\x3b\x18\x4a\xa2\xda\xc7\x4f\x2b\xb1\x4e\xbc\x8f\xfe\x3f\x5f\xfa\x57\xba\xaa\x39\x89\xa5\x44\x0f\x52\xad\x08\x15\x25\xde\x9b\xd7\x09\x66\x05\x1e\xe1\x14\xaf\x30\xf1\xd6\x02\x37\xb5\x36\x74\xe4\xba\x11\x19\x95\x49\x86\x6b\x91\xa2\xdf\x08\x43\x10\x4a\x90\xe0\xd2\xb7\x29\x97\x98\x44\xde\xe2\xc4\xc5\x21\x41\x12\x17\xb7\xa9\x16\x0a\xae\xb4\x22\xa3\x25\xbc\xe7\x0a\xe5\x3c\x6c\x4d\x8d\x97\x14\xea\x16\x4a\x83\x79\xe2\x39\xae\x36\x0e\xc3\x34\x53\x37\x36\x48\xa5\x5e\x65\xb9\xe4\x06\x83\x54\x57\x21\xbf\xe1\x9f\x43\x29\x96\x36\xa4\x8d\x20\x42\xe3\x2f\xb5\x26\x4b\x86\xd7\xe1\x38\x18\x07\xe7\x61\x6a\x6d\xb8\xd7\x05\x95\x50\x41\x6a\xad\x07\x06\x65\xe2\x59\xda\x4a\xb4\x25\x22\x79\x10\x2e\x7e\x2e\x6f\xae\x15\xf9\x7c\x83\x56\x57\x18\x4e\x82\xf3\x80\x35\x29\x8f\xd5\x3f\x93\xd5\xe1\x6d\x50\x68\x5d\x48\xe4\xb5\xb0\x4d\xd6\xd4\xda\x17\x39\xaf\x84\xdc\x26\x6f\x9d\x1d\x8d\xe1\x14\x8f\x18\x1b\x8e\x19\x1b\x4e\x18\x1b\x4e\x19\x1b\x9e\x31\xf6\x78\x26\x97\xca\xa6\x46\xd4\x04\xd6\xa4\xff\xf3\x0e\x6f\x3e\xad\xd0\x6c\xc3\x71\x10\x05\x51\x27\x34\x3b\xba\xb1\xde\x62\x1e\xb6\x01\x17\xbf\x14\xdb\x57\x9a\xb6\xe1\x28\x98\x04\x51\x58\xf3\xf4\x96\x17\x98\xed\x32\x39\x53\xb0\x53\xfe\xb6\xbc\xdf\xaa\x96\x9b\x87\xc5\xf2\x3b\x92\x55\xba\x42\x45\xc1\x8d\x0d\x47\x41\x34\x0b\xd8\x4e\xf1\x75\xfc\x26\x81\x3b\x34\x97\xaa\x57\x46\x43\x28\x47\x43\x28\xc7\xf0\xc5\xc9\xbd\xa6\xa8\xda\x02\x88\xe1\x50\x01\x97\x07\xe3\x06\x45\x51\x52\x0c\x23\xc6\x1a\xed\x9d\xfb\xc3\x5b\x74\xaa\xa5\x36\x31\x3c\x8d\x72\x36\x1e\x67\x8f\x80\x26\xf7\x40\x71\xa9\xd7\x68\xee\x43\xa7\x51\x94\xcd\xa2\x1f\x41\x83\x35\x1a\x12\x29\x97\x7e\x8a\x8a\xd0\x74\xec\x2b\xa1\xfc\xb2\xf3\x8f\x18\x7b\x76\xf9\x98\x76\x5d\xb6\xea\x4c\xd8\x5a\xf2\x6d\x0c\xb9\xc4\xcf\xad\x8a\x4b\x51\x28\x5f\x10\x56\x36\x86\x36\x72\x6b\x58\xf2\xf4\xb6\x30\x7a\xa5\x32\xff\xb1\x3d\xee\x74\x79\x9e\x5f\xfe\xe4\x67\x3c\x4a\xe1\xbe\xc2\x38\x4a\x67\x53\x78\x22\x2a\x37\x05\xb9\x6a\x91\x00\xc7\x4e\x7e\xa5\xff\xe3\x4b\xa1\x90\x1b\xbf\x30\x3c\x13\xa8\xa8\x4f\xba\x1e\xee\xe1\xec\x99\x5b\x9f\x47\xcb\xd9\x04\xa2\x89\x13\xce\x26\xd1\xc5\x8c\xc1\xb8\xb1\x9c\xf3\x68\x7a\xce\x61\x72\xe6\x84\x8b\x59\xc4\xce\x23\x38\x1b\x39\x61\x89\x8c\x9f\x4d\xe1\xfc\xc2\x09\x98\xb1\xd1\x74\x06\x17\x8d\x5b\x7e\xc1\xd8\x74\xd2\x7c\xdc\xc1\x0f\xe8\x6d\x70\x79\x2b\xe8\x4f\xcc\xf0\x6b\x66\xb0\xd4\x44\xba\xfa\x03\xf9\xe5\x42\x12\x9a\x18\x6a\xa3\x0b\x91\xc5\xaf\x3e\xbe\xa9\x78\x81\xd7\x86\x2b\x9b\x6b\x53\x05\x6f\x45\x6a\xb4\xd5\x39\x05\x7b\xda\x60\x89\x1b\xba\x72\x25\x68\xc9\x24\xa7\x1d\xf7\xd3\x21\xa0\xca\x8e\xd4\x6d\xe2\xd3\xe1\xdf\x3a\xe0\xf5\xb6\xc6\x84\xc1\xe0\xa8\xab\x5c\x56\x83\xd6\x76\xed\x54\x6b\x2b\x48\x68\x15\xbb\x59\xcf\x49\xac\xf1\x31\x5f\x5b\x73\xf5\x15\x80\x2f\xad\x96\x2b\xc2\x07\xad\xb6\x94\x3a\xbd\x6d\x75\xcd\x15\x7e\xdc\xa6\x5d\x13\x6d\x4a\xd1\xc1\xa0\x49\x04\xb5\xc1\x2e\x3c\xd4\x3c\xcb\x84\x2a\x62\x38\xab\xbb\x8e\x85\x8a\x9b\x42\xa8\x18\xd8\x01\x32\x0f\x77\x03\x6e\x1e\xb6\xaf\x95\x93\xde\x7c\xa9\xb3\x6d\x33\x5d\x33\xb1\x86\x54\x72\x6b\x13\xef\xc1\x10\x69\x5e\x21\xf7\x1c\xdc\xe3\x83\x0b\xb5\x33\xdd\xb3\x19\xbd\xf1\xa0\x49\x94\x78\x2d\x09\xbf\xad\x9e\x18\x22\x47\xaf\x83\x3c\x88\x27\x7d\x59\xf8\xd1\x68\x67\xec\xcd\xcb\x68\x17\x84\xf0\x33\xf9\xcd\x04\xda\xcf\x1e\x37\xb8\xc5\xe3\x2f\x99\x32\xda\xc5\x0f\x33\xb1\xee\xe8\x1d\x2d\xbf\xcf\x94\x74\x1d\xc3\x78\xf4\x1d\x9a\xb6\xf2\xc7\xde\x62\x2e\x76\xba\x9c\x43\xce\xfd\x8c\x13\x5f\x72\x8b\x1e\x70\x23\xb8\x5f\x8a\x2c\x43\x95\x78\x64\x56\xd8\x72\x85\x79\x53\x0e\x22\x4b\xbc\xe6\xac\x9b\xab\xa7\xe6\x6a\xd1\x1e\xbd\x7d\x94\xed\x7e\xd5\x2d\x0e\x97\x60\x63\x5e\x73\x03\x16\xcd\xda\x0d\x63\x27\x87\x21\xbc\xc2\x5c\x28\x04\x0e\x15\x52\xa9\x33\x20\x0d\x06\x53\xad\x14\xa6\x04\xab\x5a\xab\x0e\x00\x52\x5b\xbb\x8b\x71\xf0\x48\x20\x5f\xa9\xd4\x55\x6a\x7f\xd0\xd5\x56\xe7\x9f\x80\xc2\x0d\xfc\x0b\x97\x1f\x74\x7a\x8b\xd4\xef\xf7\x37\x42\x65\x7a\x13\x48\x9d\x72\x07\x70\x75\x4f\x3a\xd5\x12\x92\x24\x81\xee\x7e\xf6\x06\xf0\x02\xbc\x8d\x75\x37\xb5\x07\xb1\x5b\xba\xd5\x00\x9e\xc3\x43\x78\xa9\x2d\xc1\x73\xf0\x42\x5e\x0b\x6f\xd0\x6e\xa8\x4b\x1e\x68\x55\xa1\xb5\xbc\xc0\x63\x82\xb8\x46\x45\x3b\x96\xcd\x3e\x2a\x5b\x40\x02\x7f\xff\xf0\xee\x1f\x41\xed\x5e\xdf\xad\x4b\xe0\xce\xa6\x6d\xe6\x5e\x4f\xe4\xd0\x6f\xdc\x92\x04\xd4\x4a\xca\x3d\xbe\x67\x90\x56\x46\x75\x6e\x77\x27\xf7\xdc\x83\xe6\x8c\xe0\x49\x92\xc0\x4a\x65\xcd\x27\xce\x0e\xc8\xbf\xf4\xbd\xa7\xed\x99\x0e\x02\x57\xab\x07\xc4\x60\x1f\xee\x5e\x34\x34\x46\x9b\x6f\x45\x73\x8f\xae\xfe\x17\xc9\xb7\x7a\x45\x31\x9c\x92\xae\xaf\x9a\x9a\x3f\x1d\x82\x0b\x1e\xc3\x3e\xc2\x10\x68\x5b\x63\x0c\xa7\x8d\xe4\xec\xa2\xc2\x06\x35\x65\x8c\x0d\x61\x37\x88\xfe\xca\x4d\x0c\xae\x12\xef\xbe\xc1\xc7\xae\xd2\xd4\x0d\xac\x5f\x61\xd4\xc5\xd8\x73\xea\xe4\xff\x9f\xd5\xdd\xfd\x83\x4f\xa5\xb6\xf8\xa0\x2e\xc1\x22\x5d\xb7\x51\xfb\xfb\xda\x1d\xc2\x98\x31\x36\xb8\x84\xbb\xc3\x2c\x0e\x43\x78\x6d\x89\x2f\xa5\xb0\x25\x70\xd8\xe0\xd2\x36\xd5\x0b\x1d\x46\x68\xe5\x3a\x84\x4a\x84\x97\xef\xdf\x74\xad\x71\xd2\x14\x43\xe7\xd1\x6f\xc8\x1d\xbd\x42\xe7\x61\x3b\x30\xe7\x61\xfb\x43\xf0\xbf\x01\x00\x00\xff\xff\x6e\xb5\x9f\x64\x19\x0e\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"index.html": indexHtml,
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
	"index.html": {indexHtml, map[string]*bintree{}},
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