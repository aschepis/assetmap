package assetmap

import "testing"

func TestAssetMap(t *testing.T) {
	assetMap := NewAssetMap("test/assets.json")
	assets := map[string]string{
		"css/web-ui.css":   "css/web-ui-ac825658e28365c2.css",
		"js/web-ui.js":     "js/web-ui-158a0a39012fb9dd.js",
		"js/web-ui.min.js": "js/web-ui-140a0b39012fa9f4.min.js",
	}
	for asset, expected := range assets {
		p, e := assetMap.AssetPath(asset)
		if e != nil {
			t.Errorf("error getting asset: %v. error=%v", asset, e)
		}
		if p != expected {
			t.Errorf("unexpected asset path for asset: %v. expected=%v, got:%v", asset, expected, p)
		}
	}
}
