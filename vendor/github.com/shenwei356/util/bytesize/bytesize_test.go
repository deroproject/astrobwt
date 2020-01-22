package bytesize

import (
	// "fmt"
	"testing"
)

// test Right cases
func Test1(t *testing.T) {
	bytes := []ByteSize{3, 1000, 1024, 1203, 132434, 41234134, 132413241324, 13e15}

	data := make(map[ByteSize]string)
	for _, b := range bytes {
		// 132434 : "129.33 KB"
		data[b] = b.String()
	}
	for _, s := range data {
		// size == 132433.92
		size, err := Parse([]byte(s))

		if err != nil && err.Error() != ErrText {
			t.Error("UNKNOWN ERROR TYPE")
		}

		// size == "129.33 KB"
		if s != size.String() {
			t.Error("FAILED")
		}
	}
}

// test more
func Test2(t *testing.T) {

	data := make(map[string]ByteSize)
	data["1234.2 kb"] = 1263820.80     // lower case
	data["-1234.2 kb"] = -1263820.80   // lower case
	data[" 1234.2  kb  "] = 1263820.80 // space
	data["1234.2 k"] = 1263820.80      // simple unit
	data["1234.2 "] = 1234.2           // no unit
	data[" kb "] = -1                  // illegal value
	data["- kb"] = -1                  // illegal value
	data["1234.2 aB"] = -1             // illegal unit
	data["1234.2 Packages"] = -1       // illegal unit
	data["1234.2 P."] = -1             // illegal unit

	for s, info := range data {
		size, err := Parse([]byte(s))
		if err != nil {
			// fmt.Printf("%s\t%s\n", s, err)
			if err.Error() != ErrText || info != -1 {
				t.Error("unknown error type or test sample error")
			}
		} else { // check value
			if size != info {
				t.Error("value error")
			}
		}
		// fmt.Printf("%s\t%.2f\t%.2f\n", s, info, size)
	}
}
