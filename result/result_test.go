package result_test

import (
	"github.com/zzztttkkk/faceless/result"
	"os"
	"testing"
)

func TestFsOpen(t *testing.T) {
	name := "a.txt"
	fr := result.Call(os.Open, name)
	f := fr.Or(func() *os.File {
		fr := result.Call3(os.OpenFile, name, os.O_WRONLY|os.O_CREATE, 0666)
		return fr.Must()
	})
	defer f.Close()
}
