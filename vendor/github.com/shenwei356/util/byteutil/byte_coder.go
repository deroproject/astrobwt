package byteutil

import (
	"fmt"
	"sort"
)

var (
	// ErrInvalideLetter means invalid letter
	ErrInvalideLetter = fmt.Errorf("ByteCoder: invalid letter")
	// ErrInvalideCode means invalid code
	ErrInvalideCode = fmt.Errorf("ByteCoder: invalid code")
)

// ByteCoder is used to convert betweeen byte and int
type ByteCoder struct {
	Alphabet           []byte
	alphabetQuerySlice []byte
	bytes2int          []int
	int2bytes          []byte
}

// NewByteCoder Create a ByteCoder type
func NewByteCoder(alphabet []byte) (*ByteCoder, error) {
	if len(alphabet) == 0 {
		return nil, fmt.Errorf("ByteCoder: alphabet should not be empty")
	}

	m := make(map[byte]struct{}, len(alphabet))
	for _, a := range alphabet {
		m[a] = struct{}{}
	}

	max := -1
	var b int
	for a := range m {
		b = int(a)
		if max < b {
			max = b
		}
	}

	alphabet2 := make([]byte, len(m))
	slice := make([]byte, max+1)
	i := 0
	for a := range m {
		slice[a-'\x00'] = a

		alphabet2[i] = a
		i++
	}

	sort.Sort(ByteSlice(alphabet2))

	bytes2int := make([]int, max+1)
	int2bytes := make([]byte, len(m))
	for i, a := range alphabet2 {
		bytes2int[a-'\x00'] = i
		int2bytes[i] = a
	}

	return &ByteCoder{Alphabet: alphabet2, alphabetQuerySlice: slice,
		bytes2int: bytes2int, int2bytes: int2bytes}, nil
}

func (coder *ByteCoder) String() string {
	return fmt.Sprintf(`ByteCoder: alphabet:"%s" num:%d`, coder.Alphabet, len(coder.Alphabet))
}

// Encode converts []byte to []int
func (coder *ByteCoder) Encode(s []byte) ([]int, error) {
	code := make([]int, len(s))
	for i, b := range s {
		if int(b) > len(coder.alphabetQuerySlice) {
			return nil, ErrInvalideLetter
		}
		v := coder.alphabetQuerySlice[b-'\x00']
		if v == 0 {
			return nil, ErrInvalideLetter
		}
		code[i] = coder.bytes2int[v]
	}
	return code, nil
}

// Decode convert []int to []byte
func (coder *ByteCoder) Decode(code []int) ([]byte, error) {
	bytes := make([]byte, len(code))
	for i, b := range code {
		if b >= len(coder.int2bytes) {
			return nil, ErrInvalideCode
		}
		v := coder.int2bytes[b]
		if v == 0 {
			return nil, ErrInvalideCode
		}
		bytes[i] = v
	}
	return bytes, nil
}
