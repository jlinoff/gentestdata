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
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getProgramName() string {
	x, _ := filepath.Abs(os.Args[0])
	return filepath.Base(x)
}

type options struct {
	Deterministic    bool
	LineWidth        int
	NumLines         int
	InterleaveStderr int
	ShowLineNumbers  bool
	Alphabet         string
}

func getopts() (opts options) {
	// lambda to get the next argument on the command line.
	nextArg := func(idx *int, o string) (arg string) {
		*idx++
		if *idx < len(os.Args) {
			arg = os.Args[*idx]
		} else {
			log.Fatalf("ERROR: missing argumnent for option '%s'", o)
		}
		return
	}

	// lambda to get a range in an interval
	nextArgInt := func(idx *int, o string, min int, max int) (arg int) {
		a := nextArg(idx, o)
		arg = 0
		if v, e := strconv.Atoi(a); e == nil {
			if v < min {
				log.Fatalf("ERROR: '%v' too small, minimum accepted value is %v", o, min)
			} else if v > max {
				log.Fatalf("ERROR: '%v' too large, maximum value accepted is %v", o, max)
			}
			arg = v
		} else {
			log.Fatalf("ERROR: '%v' expected a number in the range [%v..%v]", o, min, max)
		}
		return
	}

	opts = options{
		Alphabet:        "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
		NumLines:        16,
		LineWidth:       16,
		ShowLineNumbers: false,
	}
	for i := 1; i < len(os.Args); i++ {
		opt := os.Args[i]
		switch opt {
		case "-a", "--alphabet":
			opts.Alphabet = nextArg(&i, opt)
		case "-d", "--deterministic":
			opts.Deterministic = true
		case "-h", "--help":
			help()
		case "-i", "--interleave":
			opts.InterleaveStderr = nextArgInt(&i, opt, 0, 1000000)
		case "-l", "--line-numbers":
			opts.ShowLineNumbers = true
		case "-n", "--num-lines":
			opts.NumLines = nextArgInt(&i, opt, 1, 1000000000)
		case "-V", "--version":
			fmt.Printf("%v v%v\n", getProgramName(), version)
			os.Exit(0)
		case "-w", "--width":
			opts.LineWidth = nextArgInt(&i, opt, 16, 1000000)
		default:
			log.Fatalf("ERROR: unrecognized option '%v'", opt)
		}
	}
	return
}

func help() {
	f := `
USAGE
    %[1]v [OPTIONS]

DESCRIPTION
    Program that generates text data to stdout and stderr for testing output
    handling.

    You can control the alphabet used, the number of lines, the line width and
    whether to output some or all of the lines to stderr.

    You can also control whether the output is randomly generated. Using
    deterministic output is useful for building unit tests.

    It is very useful to testing stderr handling in client/server systems.

OPTIONS
    -a STRING, --alphabet STRING
                       Specify an alternative alphabet to use for generating
                       the data.
                       The default is: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789".

    -d, --deterministic
                       Do not create random strings. Instead use the alphabet
                       unchanged. This is useful for creating unit tests.

    -h, --help         This help message.

    -i NUMBER, --interleave NUMBER
                       Interleave stderr output even NUMBER lines.
                       For example, to output every 5th line to stderr, specify
                       -i 5.
                       The default is to output everything to stdout.

    -l, --line-numbers Print line numbers in the first 7 characters.

    -n NUMBER, --num-lines NUMBER
                       The number of lines to print.
                       The default is 16.

    -V, --version      Print the program version and exit.

    -w NUMBER, --line-width NUMBER
                       The total line width, including the new line.
                       The default is 16.

EXAMPLES
    # Example 1. help
    $ %[1]v -h

    # Example 2. short output
    $ %[1]v -n 4 -w 16
    bPlNFGdSC2wd8f2
    QnFhk5A84JJjKWZ
    dKH9H2FHFuvUs9J
    z8UvBHv3Vc5awx3

    # Example 3. short output with line numbers
    $ %[1]v -n 4 -w 16 -l
         1 bPlNFGdS
         2 C2wd8f2Q
         3 nFhk5A84
         4 JJjKWZdK

    # Example 4. interleave stderr so that every 4th line is printed to stderr
    $ # All output (stdout and stderr)
    $ %[1]v -n 12 -w 32 -l -i 4
         1 bPlNFGdSC2wd8f2QnFhk5A84
         2 JJjKWZdKH9H2FHFuvUs9Jz8U
         3 vBHv3Vc5awx39ivuwsp2nChC
         4 IwVQztA2n95rXrtzhwuSAd6h
         5 eDZ0tHBxFq6Pysq3N267L1vq
         6 kgnBsUje9FqBZonjaaWDcXMm
         7 8biABkerSuHpnMmMDF2EsjYy
         8 TQWCfIuilZxV2FCniRwo7StO
         9 fGOILa0u1wXnEw1GDGuvdSew
        10 j77Ax7Tlfj84Qyu6uRn8CTEC
        11 WzT5s4ZJHd0TxrtMKykqOn91
        12 fMwNqsk2Wrc5uhk2kQaTXJp2

    $ # Only view the stderr output.
    $ %[1]v -n 12 -w 32 -l -i 4 1>/dev/null
         4 IwVQztA2n95rXrtzhwuSAd6h
         8 TQWCfIuilZxV2FCniRwo7StO
        12 fMwNqsk2Wrc5uhk2kQaTXJp2

    $ # Only view the stdout output.
    $ %[1]v -n 12 -w 32 -l -i 4 2>/dev/null
         1 bPlNFGdSC2wd8f2QnFhk5A84
         2 JJjKWZdKH9H2FHFuvUs9Jz8U
         3 vBHv3Vc5awx39ivuwsp2nChC
         5 eDZ0tHBxFq6Pysq3N267L1vq
         6 kgnBsUje9FqBZonjaaWDcXMm
         7 8biABkerSuHpnMmMDF2EsjYy
         9 fGOILa0u1wXnEw1GDGuvdSew
        10 j77Ax7Tlfj84Qyu6uRn8CTEC
        11 WzT5s4ZJHd0TxrtMKykqOn91

    # Example 5. define a different alphabet
    $ %[1]v -n 4 -w 16 -a 0123456789
    177918506041298
    415765688777805
    187196715630433
    784937199058835

    # Example 6. show the output size
    $ %[1]v -n 32 -w 32 -a 0123456789abcdef | wc
        32      32    1024

    # Example 7. generate deterministic output for unit tests
    $ %[1]v -d -l -n 12 -w 32 -a 'Lorem ipsum dolor sit amet, consectetur adipiscing elit'
         1 Lorem ipsum dolor sit ame
         2 Lorem ipsum dolor sit ame
         3 Lorem ipsum dolor sit ame
         4 Lorem ipsum dolor sit ame
         5 Lorem ipsum dolor sit ame
         6 Lorem ipsum dolor sit ame
         7 Lorem ipsum dolor sit ame
         8 Lorem ipsum dolor sit ame
         9 Lorem ipsum dolor sit ame
        10 Lorem ipsum dolor sit ame
        11 Lorem ipsum dolor sit ame
        12 Lorem ipsum dolor sit ame

VERSION
    v%[2]v
	`
	f = "\n" + strings.TrimSpace(f) + "\n\n"
	fmt.Printf(f, getProgramName(), version)
	os.Exit(0)
}
