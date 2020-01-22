package byteutil

import (
	"testing"
)

func TestByteCoder(t *testing.T) {
	coder, err := NewByteCoder([]byte("acgtryswkmbdhvACGTRYSWKMBDHV"))
	if err != nil {
		t.Error(err)
	}

	dna2int, err := coder.Encode([]byte("Jj"))
	if err != ErrInvalideLetter {
		t.Error(err)
	}

	dna2int, err = coder.Encode([]byte("acTg"))
	if err != nil {
		t.Error(err)
	}
	int2dna, err := coder.Decode(dna2int)
	if err != nil {
		t.Error(err)
	}

	if string(int2dna) != "acTg" {
		t.Errorf("ByteCoder test error")
	}

}
