package slices

import (
	"fmt"
	"testing"
)

func TestSlicePopRight(t *testing.T) {
	sv := []int{1, 2, 3, 4}
	fmt.Println(Pop(&sv))
	fmt.Println(PopLeft(&sv))
	fmt.Println(sv)
}
