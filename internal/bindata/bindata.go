package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type binaryDataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi binaryDataFileInfo) Name() string {
	return fi.name
}
func (fi binaryDataFileInfo) Size() int64 {
	return fi.size
}
func (fi binaryDataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi binaryDataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi binaryDataFileInfo) IsDir() bool {
	return false
}
func (fi binaryDataFileInfo) Sys() interface{} {
	return nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"DroidSans.ttf":   &_bintree_t{droidsans_ttf, map[string]*_bintree_t{}},
	"logo.png":        &_bintree_t{logo_png, map[string]*_bintree_t{}},
	"strip_left.jpg":  &_bintree_t{strip_left_jpg, map[string]*_bintree_t{}},
	"strip_right.jpg": &_bintree_t{strip_right_jpg, map[string]*_bintree_t{}},
}}

// binaryData is a table, holding each asset generator, mapped to its name.
var binaryData = map[string]func() (*asset, error){
	"strip_left.jpg":  strip_left_jpg,
	"strip_right.jpg": strip_right_jpg,
	"DroidSans.ttf":   droidsans_ttf,
	"logo.png":        logo_png,
}

// Asset loads and returns the assets for the given name.
// It returns an error if the assets could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := binaryData[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

// AssetInfo loads and returns the assets info for the given name.
// It returns an error if the assets could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := binaryData[canonicalName]; ok {
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
	names := make([]string, 0, len(binaryData))
	for name := range binaryData {
		names = append(names, name)
	}
	return names
}

// Restore an assets under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func strip_left_jpg_bytes() ([]byte, error) {
	return bindata_read(
		stripLeftJpg,
		"strip_left.jpg",
	)
}

func strip_left_jpg() (*asset, error) {
	jpgBytes, err := strip_left_jpg_bytes()
	if err != nil {
		return nil, err
	}

	info := binaryDataFileInfo{name: "strip_left.jpg", size: 19017, mode: os.FileMode(384), modTime: time.Unix(1439890851, 0)}
	a := &asset{bytes: jpgBytes, info: info}
	return a, nil
}

func strip_right_jpg_bytes() ([]byte, error) {
	return bindata_read(
		stripRightJpg,
		"strip_right.jpg",
	)
}

func strip_right_jpg() (*asset, error) {
	jpgBytes, err := strip_right_jpg_bytes()
	if err != nil {
		return nil, err
	}

	info := binaryDataFileInfo{name: "strip_right.jpg", size: 26871, mode: os.FileMode(384), modTime: time.Unix(1439890913, 0)}
	a := &asset{bytes: jpgBytes, info: info}
	return a, nil
}

func droidsans_ttf_bytes() ([]byte, error) {
	return bindata_read(
		droidsansTtf,
		"DroidSans.ttf",
	)
}

func droidsans_ttf() (*asset, error) {
	ttfBytes, err := droidsans_ttf_bytes()
	if err != nil {
		return nil, err
	}

	info := binaryDataFileInfo{name: "DroidSans.ttf", size: 41028, mode: os.FileMode(493), modTime: time.Unix(1437542528, 0)}
	a := &asset{bytes: ttfBytes, info: info}
	return a, nil
}

func logo_png_bytes() ([]byte, error) {
	return bindata_read(
		logoPng,
		"logo.png",
	)
}

func logo_png() (*asset, error) {
	pngBytes, err := logo_png_bytes()
	if err != nil {
		return nil, err
	}

	info := binaryDataFileInfo{name: "logo.png", size: 13698, mode: os.FileMode(420), modTime: time.Unix(1437404635, 0)}
	a := &asset{bytes: pngBytes, info: info}
	return a, nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
