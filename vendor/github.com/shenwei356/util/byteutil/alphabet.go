package byteutil

import "sort"

// Alphabet returns the alphabet of a byte slice
func Alphabet(s []byte) []byte {
	count := CountOfByte(s)
	letters := make([]byte, len(count))
	i := 0
	for b := range count {
		letters[i] = b
		i++
	}
	sort.Sort(ByteSlice(letters))
	return letters
}

// AlphabetFromCountOfByte returns the alphabet of a byte slice from count
func AlphabetFromCountOfByte(count map[byte]int) []byte {
	letters := make([]byte, len(count))
	i := 0
	for b := range count {
		letters[i] = b
		i++
	}
	sort.Sort(ByteSlice(letters))
	return letters
}
