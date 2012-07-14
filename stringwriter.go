// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package main

import (
	"fmt"
	"io"
)

var line = []byte("\"+\n\"")

type StringWriter struct {
	io.Writer
	c int
}

func (w *StringWriter) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}

	for n = range p {
		if w.c%16 == 0 {
			w.Writer.Write(line)
			w.c = 0
		}

		fmt.Fprintf(w.Writer, "\\x%02x", p[n])
		w.c++
	}

	n++

	return
}
