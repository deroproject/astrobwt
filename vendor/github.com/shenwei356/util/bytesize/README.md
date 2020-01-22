bytesize
========

Package for providing a way to show readable values of byte sizes by reediting
the code from http://golang.org/doc/effective_go.html. It could also parsing
byte size text to ByteSize object.

Usage
-------

	fmt.Printf("1024 bytes\t%v\n", bytesize.ByteSize(1024))
	fmt.Printf("13146111 bytes\t%v\n", bytesize.ByteSize(13146111))

    // parsing
	size, err := bytesize.Parse([]byte("1.5 KB"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n%.0f bytes\n", size)


Result:

    1024 bytes         1.00 KB
    13146111 bytes    12.54 MB

    1536 bytes

REGEXP for ByteSize Text
----------------------------

    (?i)^\s*([\-\d\.]+)\s*([KMGTPEZY]?B|[BKMGTPEZY]|)\s*$

Example:

    data["1234.2 kb"] = 1263820.80     lower case
    data["-1234.2 kb"] = -1263820.80   lower case
    data[" 1234.2  kb "] = 1263820.80  space
    data["1234.2 k"] = 1263820.80      simple unit
    data["1234.2 "] = 1234.2           no unit
    data[" kb "] = -1                  illegal value
    data["- kb"] = -1                  illegal value
    data["1234.2 aB"] = -1             illegal unit
    data["1234.2 Packages"] = -1       illegal unit
