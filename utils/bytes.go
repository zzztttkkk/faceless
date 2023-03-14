package utils

import "unsafe"

func B2s(bs []byte) string {
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}

func S2b(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
