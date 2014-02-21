// simple package for getting actual asset path from a json
// asset map as produced by https://www.npmjs.org/package/grunt-hashmap
package assetmap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
)

type assets map[string]string

// AssetMap is the exposed type to
type AssetMap struct {
	path   string
	assets assets
}

func NewAssetMap(path string) *AssetMap {
	m := &AssetMap{
		path: path,
	}
	m.Load()
	return m
}

func (m *AssetMap) Path() string {
	return m.path
}

func (m *AssetMap) Load() (err error) {
	buf, err := ioutil.ReadFile(m.path)
	if err != nil {
		panic(fmt.Sprintf("failed to read file(%v): %v", err, m.path))
	}
	err = json.Unmarshal(buf, &m.assets)
	if err != nil {
		panic(fmt.Sprintf("failed to decode json(%v): %v", err, string(buf)))
	}
	return
}

func (m *AssetMap) AssetPath(name string) (string, error) {
	if hash, ok := m.assets[name]; ok {
		re := regexp.MustCompile("^([^\\.]+)(.+)$")
		return re.ReplaceAllString(name, fmt.Sprintf("$1-%v$2", hash)), nil
	}

	return "", fmt.Errorf("unknown resource: %v", name)
}
