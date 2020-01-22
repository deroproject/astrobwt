package sa

import (
	"reflect"
	"testing"
)

func TestSuffixArray(t *testing.T) {
	s := []byte("banana$")
	sa := SuffixArray(s)
	if !reflect.DeepEqual(sa, []int{6, 5, 3, 1, 0, 4, 2}) {
		t.Error("Test failed: TestSuffixArray")
	}
}
