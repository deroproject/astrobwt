exec
========

Package exec runs external commands. It wraps os.exec to
allow using full command string as arguments, and provides functions
of providing (stdin, stdout, stderr) channel for (stdin, stdout, stderr) pipe.

***Attention, this package is experimental***.

This package is imported by [crun of go edition](https://github.com/shenwei356/crun/blob/master/go/crun.go)
