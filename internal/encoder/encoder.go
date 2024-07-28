package encoder

import (
	"bytes"
)

func Encode(data []byte) []byte {
	data = bytes.ReplaceAll(data, []byte("nUlL"), []byte("IaNuLl"))
	data = bytes.ReplaceAll(data, []byte{0}, []byte("nUlL"))
	return data
}

func Decode(data []byte) ([]byte, error) {
	data = bytes.ReplaceAll(data, []byte("nUlL"), []byte{0})
	data = bytes.ReplaceAll(data, []byte("IaNuLl"), []byte("nUlL"))
	return data, nil
}
