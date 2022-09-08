package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

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

var _internal_internal_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func internal_internal_go() ([]byte, error) {
	return bindata_read(
		_internal_internal_go,
		"internal/internal.go",
	)
}

var _internal_internal_gs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xc1\x6e\x13\x31\x10\xbd\xef\x57\x0c\x95\xa8\x12\x35\x4a\xe8\x05\x0e\xd1\x1e\x2b\x21\x54\xb8\xf4\x18\xe5\xb0\xcd\x4e\x12\x13\xd7\x1b\x6c\x67\x21\x8d\x22\x21\x38\x71\xaa\x04\xe2\x03\x90\x38\xc0\x85\x0b\x5c\xb8\xf0\x37\xad\xca\x5f\xa0\xb1\xbd\xb1\xd7\xbb\x11\xf4\x86\x2f\xcd\x8e\xdf\xcc\xbc\xf1\xbc\xb1\x3b\x18\x80\x44\xbd\x92\x02\x94\x96\x4c\xcc\x46\x63\xe0\x28\x66\x7a\x9e\x30\xa1\xe9\x67\x67\x67\xcf\xba\x9b\x6d\x92\x78\x07\x26\x74\x13\x6d\x8d\x31\x74\xca\x8b\xac\x05\x5c\x99\x63\xf8\x79\x51\xf0\x26\xda\x59\x2d\x98\x8c\xf3\x4c\xcd\x1d\x3d\x50\x64\xdd\x19\xe9\x47\xdd\x62\x52\x45\x36\x0a\x68\x4d\xc9\x52\x32\xa1\xb9\xe8\xd0\x47\xa6\x14\x4a\x7d\xf2\x62\x95\x71\xfb\xbd\x5c\xa2\xc8\xcd\xcf\x24\x99\xf0\x4c\x29\x38\x11\x5a\xae\xcf\x4c\xe6\x4d\x02\x00\xee\xf0\x60\x81\xeb\x5e\x99\xf1\x15\x0e\x8d\x35\x80\x81\xc0\x57\xba\x61\xad\xd8\x2f\x7a\x55\x84\xb2\x57\xf7\xea\xda\xf8\xb4\x16\xb8\x4e\x17\xc3\xdd\xa7\xc9\x93\x96\xde\x40\x19\x52\x61\xbf\xb7\xc9\xd6\x51\x7d\x9a\x2d\x43\xa2\x41\xf0\xd1\x18\x74\x76\xce\x11\x52\x18\x1d\x3f\x1c\x6f\xb6\xd6\xd5\x9c\x1d\xbb\xc4\xf4\xc1\x30\x31\x86\xe5\x4a\x77\x82\xfa\x76\x4c\x29\x7f\x40\x6f\x30\x00\x5d\xe4\x05\x4c\x64\xa1\x54\x51\xa2\x7c\xc2\x10\x6e\xde\x7d\xbd\xfe\xf6\x73\x87\x61\xd3\x0e\x55\x91\x1e\x1c\x04\x8e\xb4\x6c\xe3\x7d\x2d\x5b\xef\xe2\xba\x35\x29\x72\x62\x6a\x1a\xb7\xc0\x75\x77\x58\x43\x30\xb7\x65\x50\xf7\x8d\x5c\x4c\x6d\x5d\xf0\xb8\xf0\x5c\x29\x94\x01\x8c\xd8\xd8\x23\x8c\x20\x5e\x4a\xa6\xcd\xb6\xac\xda\x68\x99\x43\x07\xe1\x5e\x0a\x82\xf1\x88\x7b\x18\x57\xe3\xc5\xd2\x7c\x43\x0a\x81\x37\xad\x69\x21\xa1\xe3\xf7\x5d\x28\xa8\xc7\xaa\xea\x99\xac\xa4\x44\xa1\x1f\x67\x6a\x5e\x15\xbd\x73\xed\xd7\xcb\x0f\x19\xd6\xdc\x82\x13\x39\x3c\xa4\xd6\x91\xa9\x1e\xa5\x99\x9c\x96\xc7\x58\x8d\x41\x20\xe9\x78\xb9\xd3\x9a\x66\x5c\xed\x41\xc4\xad\xad\xd6\xb6\x61\x09\x0f\xcf\x73\xf0\x83\xd3\xee\x59\xb5\x11\xd2\xda\x64\x19\xa9\x1a\xe2\x3d\xc0\xe8\xb8\xe8\xa8\x0c\xf1\x96\x03\x20\xe9\x1f\x1d\xc5\x09\xbd\x30\x01\xb9\xc2\xa8\x6b\x83\x01\x5c\x5f\xbd\xb9\x7d\xfd\x16\x9e\xe7\x8b\xe3\xfe\x23\xf8\xfd\xe1\xd7\xed\xa7\x2f\xd7\x9f\x7f\xdc\x5c\xbd\xbf\xf9\xfe\xf1\x5f\xf9\x5a\xba\x24\x8b\x7a\xfe\x98\x93\xe3\xe3\xfe\xb8\x79\x9c\x61\x38\xa6\x41\x65\xff\xf9\xd0\x99\xb9\x40\xd3\xe6\xf6\xf9\xda\x33\x0f\xd8\x32\x07\xd4\x58\x27\x74\xb3\x4d\xc2\xdf\x33\x13\x2d\xad\x77\xaf\x0f\xf6\x5b\xf4\x5e\xd7\x1c\x55\x82\x91\x32\x3d\xa2\x2d\x8e\xeb\x15\xd5\x32\x43\x7d\xc6\x2e\xb1\x13\x30\xa8\x9e\x60\x76\x89\xfe\xfe\x36\x2f\xe2\x29\x13\x0b\xcc\x4f\x99\xd2\xee\x3e\xb7\x86\x67\x45\x8e\xd6\xfd\x2c\xb8\x8d\xad\xaf\x47\x04\x8f\x8e\x37\x56\x22\x29\x83\xfc\xc6\x1b\x52\x88\x5e\x13\x48\xa1\xf1\x9e\x78\x42\x9b\x38\xdb\x94\x49\xa5\x7b\xc0\x33\xa5\xf7\x3c\x25\x59\x9e\xb7\xa5\xa7\xb6\x95\xd4\x9c\x3b\x48\x34\xc8\xcb\x21\x0d\xeb\x2b\xbb\xf5\x6b\xdb\xd0\xa2\xe8\x4d\x69\xb9\x2d\xe0\xf5\x66\x53\x05\x75\xab\x1d\xfb\x4d\x03\xd5\x77\xc7\xf4\xf7\x00\x8d\x81\x86\x96\x89\xce\x57\x17\xcb\x50\x17\x41\x91\x4a\x67\x92\x62\x1a\xca\xd1\xf4\xd8\xbd\xd6\xe1\xa9\xfe\xb1\x31\x10\x2b\xc8\xf8\x72\x71\x81\x2d\xa2\x4d\xd3\x77\xd7\xee\x9f\x00\x00\x00\xff\xff\xaf\x70\x36\x63\x56\x0a\x00\x00")

func internal_internal_gs() ([]byte, error) {
	return bindata_read(
		_internal_internal_gs,
		"internal/internal.gs",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"internal/internal.go": internal_internal_go,
	"internal/internal.gs": internal_internal_gs,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"internal": &_bintree_t{nil, map[string]*_bintree_t{
		"internal.go": &_bintree_t{internal_internal_go, map[string]*_bintree_t{}},
		"internal.gs": &_bintree_t{internal_internal_gs, map[string]*_bintree_t{}},
	}},
}}