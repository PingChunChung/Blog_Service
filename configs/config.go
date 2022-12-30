// Code generated for package configs by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.yaml
package configs

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x52\x4d\x6f\xdb\x30\x0c\xbd\x17\xe8\x7f\x20\xb0\xf3\x1c\x27\xcd\x47\xa1\xd3\xfa\x91\x62\x2d\x9a\xc1\x98\x5d\xf4\x38\x30\x0e\xa3\x78\x93\x43\x41\xa2\x13\x77\xbf\x7e\x90\xdc\xc6\x5e\xea\x93\xf5\xc8\x47\xbd\xf7\xa8\x9c\xdc\x81\x9c\xba\xbc\x00\xf8\xd9\xec\x57\xbc\x21\x05\x1b\x5a\x37\x3a\x20\xdf\x45\x6c\xc6\x4e\x14\x5c\xa7\x69\x1a\x7b\x08\x37\x45\x55\x13\x37\xa2\x60\x1e\xa1\x57\x57\x09\xfd\x87\x5d\x5e\xdc\x58\x1b\x47\xde\xd3\x16\x1b\x23\x19\x6a\xca\xab\xbf\xa4\x60\x1c\x29\x2b\x6c\x87\x50\xc4\x9e\x59\xe7\x78\xa0\x0c\x65\xa7\xc0\x0b\x3b\xd4\x34\x32\xac\xfd\x7b\xf1\xa1\x32\xf4\x03\x6b\x52\x80\xd6\x0e\xb0\x65\x2b\x0a\x12\xc3\x51\xf0\x8b\x35\x8c\x9b\xcf\x73\x9a\x88\xfb\x41\x4b\xb4\xfd\xe2\x8c\x82\x9d\x88\x55\xa3\xd1\x78\xb2\x48\xd2\x24\x4d\xc6\x2a\x78\x1d\x79\x41\xa9\xca\x9e\xf0\x58\xa3\xa6\x15\xb6\x9d\xe6\x19\x7c\x81\xd5\xed\x59\xf5\xc6\x18\x3e\x2e\x5b\xf1\xd1\x3a\xc0\x57\x48\x7e\x5b\x3d\xf8\xa7\xfe\x60\xf7\x7a\x90\xcf\x1d\xef\x85\x5a\x39\x0b\xf1\x1e\x05\xd7\xe8\xa9\x4b\xf2\xb6\x78\xb3\xa4\xc0\xb2\x17\xed\xa8\xb3\xe2\xc9\xed\x63\x24\x43\x34\x43\xef\x8f\xec\x36\x0a\xc6\x93\xab\xe9\x6c\xbe\xb8\x8e\xab\x64\x2f\x0a\x0c\x97\x68\x76\xec\xa5\x1b\xd9\xe5\xb9\x36\xac\x7f\x79\x72\x87\xaa\xa4\x80\x17\xb8\x36\x94\x39\xda\x56\xed\x7b\x31\xa0\x77\x3b\x74\x9e\x44\x41\x23\xdb\x38\x31\x73\x3e\xae\x5d\x41\xe1\x9a\x48\xec\x9e\xca\x6c\x7a\x35\x0e\xf2\x9f\x5e\x8b\xa8\x3c\xa7\xd2\x05\xde\x13\x96\x7f\xc2\xf9\xd1\xfb\x86\xdc\xe7\x6b\x97\xad\xad\x1c\x29\x58\x4c\xd2\x68\x7f\x59\x63\x65\x54\xaf\xdd\xd7\x62\x13\xcd\xac\x0d\x25\x25\xd7\xfd\x85\xd3\xf9\xec\x23\x8d\xce\x50\x1b\xbf\x6f\x3a\x0c\x38\xb5\x9e\x52\x09\xb5\x4e\x47\x9e\x3f\xf7\xe2\x1f\x1c\xd7\x1f\xd4\x73\x72\xc1\xa7\x95\x9e\x1a\xde\x70\xc7\x1c\x1a\x12\x39\xfe\x0b\x00\x00\xff\xff\x32\x8e\x18\xe1\x45\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 837, mode: os.FileMode(438), modTime: time.Unix(1672378653, 0)}
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
	"configs/config.yaml": configsConfigYaml,
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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
