// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package i18n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
		"es_MX": &dictionary{index: es_MXIndex, data: es_MXData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"Hello %s":   0,
	"Welcome!\n": 1,
}

var en_USIndex = []uint32{ // 3 elements
	0x00000000, 0x0000000c, 0x0000001a,
} // Size: 36 bytes

const en_USData string = "\x02Hello %[1]s\x04\x00\x01\n\t\x02Welcome!"

var es_MXIndex = []uint32{ // 3 elements
	0x00000000, 0x0000000b, 0x0000001c,
} // Size: 36 bytes

const es_MXData string = "\x02Hola %[1]s\x04\x00\x01\n\f\x02Bienvenido!"

// Total table size 126 bytes (0KiB); checksum: 1EEF4A86
