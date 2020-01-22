package stringutil

import (
	"bytes"

	"github.com/shenwei356/natsort"
)

// String2ByteSlice is for sortint of string-[]byte pairs
type String2ByteSlice struct {
	Key   string
	Value []byte
}

// String2ByteSliceList is list of string2ByteSlice
type String2ByteSliceList []String2ByteSlice

// NaturalOrder is the global variable for sorting String2ByteSlice
var NaturalOrder = false

// IgnoreCase for ignoring case when sorting in natural order
var IgnoreCase = false

func (list String2ByteSliceList) Len() int { return len(list) }
func (list String2ByteSliceList) Less(i, j int) bool {
	if NaturalOrder {
		return natsort.Compare(list[i].Key, list[j].Key, IgnoreCase)
	}
	return list[i].Key < list[j].Key
}
func (list String2ByteSliceList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// ReversedString2ByteSliceList is reversed String2ByteSliceList
type ReversedString2ByteSliceList struct {
	String2ByteSliceList
}

// Less ...
func (list ReversedString2ByteSliceList) Less(i, j int) bool {
	if NaturalOrder {
		return !natsort.Compare(list.String2ByteSliceList[i].Key, list.String2ByteSliceList[j].Key, IgnoreCase)
	}
	return list.String2ByteSliceList[i].Key > list.String2ByteSliceList[j].Key
}

// ByValue ...
type ByValue struct {
	String2ByteSliceList
}

// Less ...
func (list ByValue) Less(i, j int) bool {
	c := bytes.Compare(list.String2ByteSliceList[i].Value, list.String2ByteSliceList[j].Value)
	if c == -1 {
		return true
	}
	return false
}

// ReversedByValue ...
type ReversedByValue struct {
	String2ByteSliceList
}

// Less ...
func (list ReversedByValue) Less(i, j int) bool {
	c := bytes.Compare(list.String2ByteSliceList[j].Value, list.String2ByteSliceList[i].Value)
	if c == -1 {
		return true
	}
	return false
}
