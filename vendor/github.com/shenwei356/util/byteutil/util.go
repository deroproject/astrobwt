package byteutil

import (
	"bytes"
	// "fmt"
	"unsafe"

	"github.com/shenwei356/bpool"
)

// ReverseByteSlice reverses a byte slice
func ReverseByteSlice(s []byte) []byte {
	// make a copy of s
	l := len(s)
	t := make([]byte, l)
	for i := 0; i < l; i++ {
		t[i] = s[i]
	}

	// reverse
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
	return t
}

// ReverseByteSliceInplace reverses a byte slice
func ReverseByteSliceInplace(s []byte) {
	// reverse
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// WrapByteSlice wraps byte slice
func WrapByteSlice(s []byte, width int) []byte {
	if width < 1 {
		return s
	}
	l := len(s)
	if l == 0 {
		return s
	}
	var lines int
	if l%width == 0 {
		lines = l/width - 1
	} else {
		lines = int(l / width)
	}
	// var buffer bytes.Buffer
	buffer := bytes.NewBuffer(make([]byte, 0, l+lines))
	var start, end int
	for i := 0; i <= lines; i++ {
		start = i * width
		end = (i + 1) * width
		if end > l {
			end = l
		}

		buffer.Write(s[start:end])
		if i < lines {
			buffer.WriteString("\n")
		}
	}
	return buffer.Bytes()
}

// BufferedByteSliceWrapper is used to wrap byte slice,
// using a buffer of bytes.Buffer to reduce GC
type BufferedByteSliceWrapper struct {
	pool *bpool.SizedBufferPool
}

// NewBufferedByteSliceWrapper create a new BufferedByteSliceWrapper
func NewBufferedByteSliceWrapper(size, alloc int) *BufferedByteSliceWrapper {
	if size < 1 {
		panic("buffer number should be > 0")
	}
	if alloc < 1 {
		panic("buffer size should be > 0")
	}
	return &BufferedByteSliceWrapper{bpool.NewSizedBufferPool(size, alloc)}
}

// NewBufferedByteSliceWrapper2 could pre-alloc space according to length of slice and width
func NewBufferedByteSliceWrapper2(size int, length, width int) *BufferedByteSliceWrapper {
	if size < 1 {
		// panic("buffer number should be > 0")
		size = 1
	}
	if length < 1 {
		// panic("buffer size should be > 0")
		length = 1
	}
	if width <= 0 {
		return NewBufferedByteSliceWrapper(size, length)
	}
	var lines int
	if length%width == 0 {
		lines = length/width - 1
	} else {
		lines = int(length / width)
	}
	return &BufferedByteSliceWrapper{bpool.NewSizedBufferPool(size, length+lines)}
}

// Recycle a buffer
func (w *BufferedByteSliceWrapper) Recycle(b *bytes.Buffer) {
	w.pool.Put(b)
}

// Wrap a byte slice. DO NOT FORGET call Recycle() with the returned buffer
func (w *BufferedByteSliceWrapper) Wrap(s []byte, width int) ([]byte, *bytes.Buffer) {
	if width < 1 {
		return s, nil
	}
	l := len(s)
	if l == 0 {
		return s, nil
	}
	var lines int
	if l%width == 0 {
		lines = l/width - 1
	} else {
		lines = int(l / width)
	}
	// var buffer bytes.Buffer
	// buffer := bytes.NewBuffer(make([]byte, 0, l+lines))

	buffer := w.pool.Get()

	var start, end int
	for i := 0; i <= lines; i++ {
		start = i * width
		end = (i + 1) * width
		if end > l {
			end = l
		}

		buffer.Write(s[start:end])
		if i < lines {
			buffer.WriteString("\n")
		}
	}
	return buffer.Bytes(), buffer
}

// WrapByteSliceInplace wraps byte slice in place.
// Sadly, it's too slow. Never use this!
func WrapByteSliceInplace(s []byte, width int) []byte {
	if width < 1 {
		return s
	}
	var l, lines int

	l = len(s)
	if l%width == 0 {
		lines = l/width - 1
	} else {
		lines = int(l / width)
	}

	var end int
	j := 0
	for i := 0; i <= lines; i++ {
		end = (i+1)*width + j
		if end >= l {
			break
		}
		// fmt.Printf("len:%d, lines:%d, i:%d, j:%d, end:%d\n", l, lines, i, j, end)
		if i < lines {

			// https://github.com/golang/go/wiki/SliceTricks
			// Sadly, it's too slow
			// s = append(s, []byte(" ")[0])
			// copy(s[end+1:], s[end:])
			// s[end] = []byte("\n")[0]

			// slow too
			s = append(s[:end], append([]byte("\n"), s[end:]...)...)

			l = len(s)
			if l%width == 0 {
				lines = l/width - 1
			} else {
				lines = int(l / width)
			}

			j++
		}
	}
	return s
}

// SubSlice provides similar slice indexing as python with one exception
// that end could be equal to 0.
// So we could get the last element by SubSlice(s, -1, 0)
// or get the whole element by SubSlice(s, 0, 0)
func SubSlice(slice []byte, start int, end int) []byte {
	if start == 0 && end == 0 {
		return slice
	}
	if start == end || (start < 0 && end > 0) {
		return []byte{}
	}
	l := len(slice)
	s, e := start, end

	if s < 0 {
		s = l + s
		if s < 1 {
			s = 0
		}
	}
	if e < 0 {
		e = l + e
		if e < 0 {
			e = 0
		}
	}
	if e == 0 || e > l {
		e = l
	}
	return slice[s:e]
}

// ByteToLower lowers a byte
func ByteToLower(b byte) byte {
	if b <= '\u007F' {
		if 'A' <= b && b <= 'Z' {
			b += 'a' - 'A'
		}
		return b
	}
	return b
}

// ByteToUpper upper a byte
func ByteToUpper(b byte) byte {
	if b <= '\u007F' {
		if 'a' <= b && b <= 'z' {
			b -= 'a' - 'A'
		}
		return b
	}
	return b
}

// MakeQuerySlice is used to replace map.
// see: http://blog.shenwei.me/map-is-not-the-fastest-in-go/
func MakeQuerySlice(letters []byte) []byte {
	max := -1
	for i := 0; i < len(letters); i++ {
		j := int(letters[i])
		if max < j {
			max = j
		}
	}
	querySlice := make([]byte, max+1)
	for i := 0; i < len(letters); i++ {
		querySlice[int(letters[i])] = letters[i]
	}
	return querySlice
}

// Split splits a byte slice by giveen letters.
// It's much faster than regexp.Split
func Split(slice []byte, letters []byte) [][]byte {
	querySlice := MakeQuerySlice(letters)
	results := [][]byte{}
	tmp := []byte{}

	var j int
	var value byte
	var sliceSize = len(querySlice)
	for _, b := range slice {
		j = int(b)
		if j >= sliceSize { // not delimiter byte
			tmp = append(tmp, b)
			continue
		}
		value = querySlice[j]
		if value == 0 { // not delimiter byte
			tmp = append(tmp, b)
			continue
		} else {
			if len(tmp) > 0 {
				results = append(results, tmp)
				tmp = []byte{}
			}
		}
	}
	if len(tmp) > 0 {
		results = append(results, tmp)
	}
	return results
}

// Bytes2Str convert byte slice to string without GC. Warning: it's unsafe!!!
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// CountBytes counts given ASCII characters in a byte slice
func CountBytes(seq, letters []byte) int {
	if len(letters) == 0 || len(seq) == 0 {
		return 0
	}

	// do not use map
	querySlice := make([]byte, 256)
	for i := 0; i < len(letters); i++ {
		querySlice[int(letters[i])] = letters[i]
	}

	var g byte
	var n int
	for i := 0; i < len(seq); i++ {
		g = querySlice[int(seq[i])]
		if g > 0 { // not gap
			n++
		}
	}
	return n
}
