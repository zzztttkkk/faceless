package h2tp

import (
	"testing"
)

func TestHeaders_Add(t *testing.T) {
	var headers Headers

	headers.Add([]byte("hello"), []byte("world"))
}
