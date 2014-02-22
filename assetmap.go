// simple package for getting actual asset path from a json
// asset map as produced by https://www.npmjs.org/package/grunt-hashmap
package assetmap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type assets map[string]string

// AssetMap is the exposed type to
type AssetMap struct {
	path           string
	stat           os.FileInfo
	reloadOnChange bool
	assets         assets
}

// Create a new AssetMap that reads from the file located
// at the path argument.
func NewAssetMap(path string, reloadOnChange bool) *AssetMap {
	stat, _ := os.Stat(path)
	m := &AssetMap{
		path:           path,
		reloadOnChange: reloadOnChange,
		stat:           stat,
	}
	m.Load()
	return m
}

// retrieve the path to the asset map file
func (m *AssetMap) Path() string {
	return m.path
}

// load/reload asset map from file
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

// get the real path to an asset based on its original path
func (m *AssetMap) AssetPath(name string) (string, error) {
	if m.reloadOnChange {
		stat, err := os.Stat(m.path)
		if err == nil && (stat.Size() != m.stat.Size() || stat.ModTime() != m.stat.ModTime()) {
			m.stat = stat
			m.Load()
		}
	}

	if hash, ok := m.assets[name]; ok {
		re := regexp.MustCompile("^([^\\.]+)(.+)$")
		return re.ReplaceAllString(name, fmt.Sprintf("$1-%v$2", hash)), nil
	}

	return "", fmt.Errorf("unknown resource: %v", name)
}
