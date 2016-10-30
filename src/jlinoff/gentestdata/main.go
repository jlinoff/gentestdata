// Program to generate test data to stdout and stderr.
// See the help for more detailed information.
/*
License: The MIT License (MIT)

Copyright (c) 2016 Joe Linoff

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject
to the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR
ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
)

var version = "0.3" // added -d support

func main() {
	opts := getopts()
	line := ""
	dline := getDline(opts)
	for i := 0; i < opts.NumLines; i++ {

		// Get the line data.
		if opts.Deterministic {
			line = dline
		} else {
			line = getRandomString(opts.LineWidth, opts.Alphabet)
		}

		// Insert line numbers if necessary.
		if opts.ShowLineNumbers {
			line = fmt.Sprintf("%6d %s", i+1, line)
		}

		// Trim the line.
		line = line[:opts.LineWidth-1] // make room for the new line

		// Output the line to stdout or stderr.
		if opts.InterleaveStderr == 0 || (i%opts.InterleaveStderr) != (opts.InterleaveStderr-1) {
			fmt.Fprintln(os.Stdout, line)
		} else {
			fmt.Fprintln(os.Stderr, line)
		}
	}
}

// Get the deterministic line.
func getDline(opts options) string {
	// Make sure that the dline is large enough to slice.
	dline := opts.Alphabet // deterministic line
	if opts.Deterministic {
		for len(dline) < opts.LineWidth {
			dline += opts.Alphabet
		}
	}
	return dline[:opts.LineWidth]
}

// Generate a random string of fix length from the alphabet.
func getRandomString(length int, alphabet string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}
