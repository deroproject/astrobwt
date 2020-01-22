package byteutil

import "bytes"

// SliceOfByteSlice is [][]byte
type SliceOfByteSlice [][]byte

func (s SliceOfByteSlice) Len() int { return len(s) }
func (s SliceOfByteSlice) Less(i, j int) bool {
	c := bytes.Compare(s[i], s[j])
	if c == -1 {
		return true
	}
	return false
}
func (s SliceOfByteSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByteSlice is []byte
type ByteSlice []byte

func (bs ByteSlice) Len() int           { return len(bs) }
func (bs ByteSlice) Less(i, j int) bool { return bs[i] < bs[j] }
func (bs ByteSlice) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
