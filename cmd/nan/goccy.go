package main

import "bytes"

var goccyReplace = []struct {
	src           []byte
	dst           []byte
	checkNoEscape bool
}{
	{src: []byte(`"encoding/json"`), dst: []byte(`json "github.com/goccy/go-json"`)},
	{src: []byte("json.Marshal("), dst: []byte("json.MarshalNoEscape("), checkNoEscape: true},
	{src: []byte("json.Unmarshal("), dst: []byte("json.UnmarshalNoEscape("), checkNoEscape: true},
}

func convertJsonToGoccyJson(str []byte, noEscape bool) []byte {
	for _, r := range goccyReplace {
		if !r.checkNoEscape || (r.checkNoEscape && noEscape) {
			if bytes.Contains(str, r.src) {
				str = bytes.ReplaceAll(str, r.src, r.dst)
			}
		}
	}
	return str
}
