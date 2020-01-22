package cliutil

import (
	"strings"

	"github.com/shenwei356/breader"
)

// ReadKVs parse two-column (key\tvalue) tab-delimited file
func ReadKVs(file string, ignoreCase bool) (map[string]string, error) {
	type KV [2]string
	fn := func(line string) (interface{}, bool, error) {
		line = strings.TrimRight(line, "\r\n")
		if len(line) == 0 {
			return nil, false, nil
		}
		items := strings.Split(line, "\t")
		if len(items) < 2 {
			return nil, false, nil
		}
		if ignoreCase {
			return KV([2]string{strings.ToLower(items[0]), items[1]}), true, nil
		}
		return KV([2]string{items[0], items[1]}), true, nil
	}
	kvs := make(map[string]string)
	reader, err := breader.NewBufferedReader(file, 2, 10, fn)
	if err != nil {
		return kvs, err
	}
	var items KV
	for chunk := range reader.Ch {
		if chunk.Err != nil {
			return kvs, err
		}
		for _, data := range chunk.Data {
			items = data.(KV)
			kvs[items[0]] = items[1]
		}
	}
	return kvs, nil
}
