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

var version = "0.2"

func main() {
	opts := getopts()
	m := ""
	for i := 0; i < opts.NumLines; i++ {
		if opts.ShowLineNumbers {
			s := getRandomString(opts.LineWidth-8, opts.Alphabet)
			m = fmt.Sprintf("%6d %s", i+1, s)
		} else {
			m = getRandomString(opts.LineWidth-1, opts.Alphabet) // leave room for newline
		}
		if opts.InterleaveStderr == 0 || (i%opts.InterleaveStderr) != (opts.InterleaveStderr-1) {
			fmt.Fprintln(os.Stdout, m)
		} else {
			fmt.Fprintln(os.Stderr, m)
		}
	}
}

func getRandomString(width int, alphabet string) string {
	b := make([]byte, width)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}
