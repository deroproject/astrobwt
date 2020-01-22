package byteutil

import "testing"

func TestSubSlice(t *testing.T) {
	s := []byte("0123456789")
	if true &&
		string(SubSlice(s, 0, 0)) == "0123456789" &&
		string(SubSlice(s, 0, 1)) == "0" &&
		string(SubSlice(s, 1, 2)) == "1" &&
		string(SubSlice(s, -2, -1)) == "8" &&
		string(SubSlice(s, len(s)-1, len(s))) == "9" &&
		string(SubSlice(s, -1, 0)) == "9" && // different from python
		string(SubSlice(s, 7, -1)) == "78" &&
		true {
	} else {
		t.Error("SubSlice error")
	}
}
