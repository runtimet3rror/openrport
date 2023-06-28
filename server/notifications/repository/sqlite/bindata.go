// Code generated by go-bindata.
// sources:
// 001_init.down.sql
// 001_init.up.sql
// DO NOT EDIT!

package sqlite

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

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\x8e\xcf\xc9\x4f\xb7\x06\x04\x00\x00\xff\xff\x5d\x68\xf4\x86\x1d\x00\x00\x00")

func _001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initDownSql,
		"001_init.down.sql",
	)
}

func _001_initDownSql() (*asset, error) {
	bytes, err := _001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.down.sql", size: 29, mode: os.FileMode(436), modTime: time.Unix(1687534397, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x4d\x4f\xdb\x40\x10\xbd\xfb\x57\x4c\x73\x81\x48\x4e\x85\x90\xe0\x82\x7a\x70\x83\x2b\x50\x43\x40\xc1\x54\xdc\xac\xb5\x77\x1c\x6f\xb1\x77\xdd\xd9\x31\xe0\xfe\xfa\x6a\xd7\x4e\x42\x42\x08\x97\xfa\x38\x3b\xef\x3d\xbf\xf9\x9a\x2e\xe2\x28\x89\x21\x89\xbe\xcf\x62\xd0\x86\x55\xa1\x72\xc1\xca\x68\x9b\x56\x66\x09\xc7\x01\x00\x6c\xc5\x53\x25\x61\x7a\x15\x2d\x8e\x4f\xcf\xc7\x30\xbf\x4d\x60\xfe\x30\x9b\xc1\xf4\x2a\x9e\xfe\x84\xe3\xdd\xc4\x2f\xdf\xe0\xe8\x68\x1c\xc2\x64\xb2\xcb\x11\x82\x6e\xeb\x0c\x09\x4c\x01\x5c\x22\x38\x31\xd4\x4c\x5d\x08\x6d\xa5\xa4\x97\x65\x55\xa3\x65\x51\x37\x70\x19\x25\x71\x72\x7d\x13\x6f\x04\x2f\xe3\x1f\xd1\xc3\x2c\x81\xe9\xc3\x62\x11\xcf\x93\xd4\xbd\xde\x27\xd1\xcd\x5d\xe8\x90\x93\xc9\x06\x1c\x82\x14\x8c\x20\xb4\xf4\x31\x27\x58\xa3\xb5\x62\x89\x20\x95\x6d\x04\xe7\xa5\xd2\x4b\x2f\x68\x48\x2d\x95\x86\x5f\xd1\xa2\x77\x78\x32\x7e\x2f\x38\x1a\x79\x05\x18\x64\x7a\x48\x08\x5a\xf4\xd4\xce\x8b\x6d\x33\xdb\x59\xc6\x1a\xb8\x14\x0c\xa5\xb0\x90\x13\x0a\x46\xe9\x9f\xdf\x56\xc2\xcb\x16\x6d\x55\xdd\xee\x48\x9f\x9d\x8d\x61\xaf\xb8\x47\xe4\x46\x33\x6a\x4e\xba\x06\xd7\x90\xb3\x93\x43\x08\xc2\x02\x09\x75\x8e\xae\x2d\x6b\x95\xf3\x0f\x20\xbd\x37\xc2\x62\x05\x09\xa1\x21\x93\x55\x58\x83\x92\x50\x0a\x2d\x51\x82\x79\x46\x82\xac\xf3\x9e\x44\xc5\x48\x4a\x2f\xc1\x22\x3d\xab\x1c\xbf\xf6\x0d\x24\xa1\x6d\x63\x88\x21\x89\x1f\x93\xc3\xb5\xfc\xfc\x73\x4d\x5d\x11\x86\x80\x8a\x4b\x24\x18\xd9\x9a\x9b\x11\x18\xda\xee\x40\x4e\xaa\xe1\xc1\x78\xae\x1a\x85\x9a\xed\x7f\xf8\x09\x5f\x94\x15\x5f\xe8\xcc\x2a\x51\xa9\xbf\x28\x41\x10\x89\x2e\xf0\x8a\x96\xdd\xbc\xed\x1d\xa2\x61\x4d\xfa\x8c\x37\xcb\x61\xff\x54\x8a\x11\xa4\x41\xab\x8f\x18\x2a\xf5\x84\x80\xba\xad\xed\xc0\xd8\x66\xbf\x31\x3f\x50\x45\x9f\x95\x19\xd9\x1d\x48\xe9\x27\xbc\xfd\x98\xc5\xfd\xc8\x4b\x29\x18\x5d\x5f\x91\xc8\x10\x28\x0b\x84\xdc\x92\x76\xed\xd6\x55\x07\x2f\x25\xea\xf7\x6f\xc1\xf8\x22\x98\x4c\x82\x60\x38\x26\xd7\xf3\xcb\xf8\x11\x94\x7c\x4d\xb7\x0f\x8a\x92\xc1\xed\x7c\xdf\x91\xd9\xb9\x0d\xe3\x8b\x4f\xa9\xd6\xdb\xed\x5d\xed\x67\x5d\xe7\x38\x3e\x57\x64\x16\xdc\xda\x34\x37\x12\x43\x70\x63\x33\x44\xc0\x45\x36\x3e\x87\x89\xbe\xbf\x49\xee\xfc\x34\xbb\x1b\x45\x3e\x84\xaf\x8a\xfb\xe4\xed\x31\x73\xdc\x35\x37\xa9\xaf\x8b\x0d\x87\xfa\xac\x8e\xcc\x9a\xb8\x20\x53\xef\x52\x7b\xac\x67\x49\x09\x6d\x63\xb4\xc5\x10\x0a\x45\x96\xe1\xf4\xc9\xc9\x58\x96\xae\x67\xee\x7a\x59\x96\x48\xb4\xad\xfd\x2f\x00\x00\xff\xff\xae\x26\x81\x94\xbf\x05\x00\x00")

func _001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initUpSql,
		"001_init.up.sql",
	)
}

func _001_initUpSql() (*asset, error) {
	bytes, err := _001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.up.sql", size: 1471, mode: os.FileMode(436), modTime: time.Unix(1687595215, 0)}
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
	"001_init.down.sql": _001_initDownSql,
	"001_init.up.sql":   _001_initUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"001_init.down.sql": &bintree{_001_initDownSql, map[string]*bintree{}},
	"001_init.up.sql":   &bintree{_001_initUpSql, map[string]*bintree{}},
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
