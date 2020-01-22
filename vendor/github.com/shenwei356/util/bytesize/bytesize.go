// Copyright 2014 Wei Shen (shenwei356@gmail.com). All rights reserved.
// Use of this source code is governed by a MIT-license
// that can be found in the LICENSE file.

// Package bytesize provides a way to show readable values of byte size
// by reediting the code from http://golang.org/doc/effective_go.html.
// It could also parsing byte size text to ByteSize object.
package bytesize

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Division operation is needed, so it uses float64 instead of uint64
type ByteSize float64

// const for bytesize. B is also specified.
const (
	B ByteSize = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// Print readable values of byte size
func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%7.2f YB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%7.2f ZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%7.2f EB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%7.2f PB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%7.2f TB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%7.2f GB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%7.2f MB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%7.2f KB", b/KB)
	}
	return fmt.Sprintf("%7.2f  B", b)
}

// Regexp object for ByteSize Text. The REGEXP is:
//
//     (?i)^\s*([\-?[\d\.]+)\s*([KMGTPEZY]?B|[BKMGTPEZY]|)\s?$
//
// Example:
//
//    data["1234.2 kb"] = 1263820.80    // lower case
//    data["-1234.2 kb"] = -1263820.80  // lower case
//    data[" 1234.2  kb "] = 1263820.80 // space
//    data["1234.2 k"] = 1263820.80     // simple unit
//    data["1234.2 "] = 1234.2          // no unit
//    data[" kb "] = -1                 // illegal value
//    data["- kb"] = -1                 // illegal value
//    data["1234.2 aB"] = -1            // illegal unit
//    data["1234.2 Packages"] = -1      // illegal unit
//
var BytesizeRegexp = regexp.MustCompile(`(?i)^\s*(\-?[\d\.]+)\s*([KMGTPEZY]?B|[BKMGTPEZY]|)\s*$`)

// Error information for Illegal byte size text
var ErrText = "illegal bytesize text"

// Parse ByteSize Text to ByteSize object
//
// Example
//
//     size, err := bytesize.Parse([]byte("1.5 KB"))
//     if err != nil {
//         fmt.Println(err)
//     }
//     fmt.Printf("%.0f bytes\n", size)
//
func Parse(sizeText []byte) (ByteSize, error) {
	if !BytesizeRegexp.Match(sizeText) {
		return 0, errors.New(ErrText)
	}

	// parse value and unit
	subs := BytesizeRegexp.FindSubmatch(sizeText)

	// no need to check ParseFloat error. BytesizeRegexp could ensure this

	size, _ := strconv.ParseFloat(string(subs[1]), 64)
	unit := strings.ToUpper(string(subs[2]))

	switch unit {
	case "B", "":
		size = size * float64(B)
	case "KB", "K":
		size = size * float64(KB)
	case "MB", "M":
		size = size * float64(MB)
	case "GB", "G":
		size = size * float64(GB)
	case "TB", "T":
		size = size * float64(TB)
	case "PB", "P":
		size = size * float64(PB)
	case "EB", "E":
		size = size * float64(EB)
	case "ZB", "Z":
		size = size * float64(ZB)
	case "YB", "Y":
		size = size * float64(YB)
	}

	// fmt.Printf("%s\t=%.2f=\t=%s=\n", sizeText, size, unit)

	return ByteSize(size), nil
}
