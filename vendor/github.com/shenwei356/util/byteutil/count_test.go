package byteutil

import "testing"

func TestCountOfByteAndAlphabet(t *testing.T) {
	s := []byte("abcdefadfwefasdga")
	count := CountOfByte(s)
	alphabet := Alphabet(s)
	sum := 0
	for _, letter := range alphabet {
		sum += count[letter]
	}
	if sum != len(s) {
		t.Error("Test failed: TestCountOfByteAndAlphabet")
	}
}

func TestSortCountOfByte(t *testing.T) {
	s := []byte("cccaaadd")
	countList := SortCountOfByte(CountOfByte(s), true)
	// fmt.Println(countList)
	// if !(countList[0].Count == 3 && (countList[0].Key == 'a' || countList[0].Key == 'c')) {
	if !(countList[0].Count == 3 && countList[0].Key == 'a') {
		t.Error("Test failed: TestSortCountOfByte")
	}

	countList = SortCountOfByte(CountOfByte(s), false)
	// fmt.Println(countList)
	if !(countList[0].Key == 'd' && countList[0].Count == 2) {
		t.Error("Test failed: TestSortCountOfByte")
	}
}
