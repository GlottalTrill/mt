package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
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
			return nil, fmt.Errorf("asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("asset %s not found", name)
}

// get font path for fontname
// searches in common font paths, bindata or absoute path
func GetFont(f string) ([]byte, error) {
	if !strings.HasSuffix(f, ".ttf") {
		f = fmt.Sprintf("%s.ttf", f)
	}
	if strings.Contains(f, "/") && strings.HasSuffix(f, ".ttf") {
		if _, err := os.Stat(f); err == nil {
			log.Infof("using font: %s", f)
			return ioutil.ReadFile(f)
		}
	}
	fdirs := []string{"/Library/Fonts/", "/usr/share/fonts/", "./"}

	for _, dir := range fdirs {
		fpath := filepath.Join(dir, f)
		if _, err := os.Stat(fpath); err == nil {
			log.Infof("using font: %s", fpath)
			return ioutil.ReadFile(fpath)
		}
	}
	log.Info("using font: DroidSans.ttf")
	return Asset("DroidSans.ttf")
}

func readBinaryData(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	copyErr := gz.Close()
	if copyErr != nil {
		return nil, copyErr
	}

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func stripLeftJpgBytes() ([]byte, error) {
	return readBinaryData(
		stripLeftJpg,
		"strip_left.jpg",
	)
}

func strip_left_jpg() (*asset, error) {
	jpgBytes, err := stripLeftJpgBytes()
	if err != nil {
		return nil, err
	}

	info := binaryDataFileInfo{name: "strip_left.jpg", size: 19017, mode: os.FileMode(384), modTime: time.Unix(1439890851, 0)}
	a := &asset{bytes: jpgBytes, info: info}
	return a, nil
}

func stripRightJpgBytes() ([]byte, error) {
	return readBinaryData(
		stripRightJpg,
		"strip_right.jpg",
	)
}

func strip_right_jpg() (*asset, error) {
	jpgBytes, err := stripRightJpgBytes()
	if err != nil {
		return nil, err
	}

	info := binaryDataFileInfo{name: "strip_right.jpg", size: 26871, mode: os.FileMode(384), modTime: time.Unix(1439890913, 0)}
	a := &asset{bytes: jpgBytes, info: info}
	return a, nil
}

func droidsansTtfBytes() ([]byte, error) {
	return readBinaryData(
		droidsansTtf,
		"DroidSans.ttf",
	)
}

func droidsans_ttf() (*asset, error) {
	ttfBytes, err := droidsansTtfBytes()
	if err != nil {
		return nil, err
	}

	info := binaryDataFileInfo{name: "DroidSans.ttf", size: 41028, mode: os.FileMode(493), modTime: time.Unix(1437542528, 0)}
	a := &asset{bytes: ttfBytes, info: info}
	return a, nil
}

func logoPngBytes() ([]byte, error) {
	return readBinaryData(
		logoPng,
		"logo.png",
	)
}

func logo_png() (*asset, error) {
	pngBytes, err := logoPngBytes()
	if err != nil {
		return nil, err
	}

	info := binaryDataFileInfo{name: "logo.png", size: 13698, mode: os.FileMode(420), modTime: time.Unix(1437404635, 0)}
	a := &asset{bytes: pngBytes, info: info}
	return a, nil
}
