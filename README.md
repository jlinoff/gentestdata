# gentestdata
Go program to interleave randomly generated test data between stdout and stderr.

I use it to test whether code that spawns and manages subprocesses
properly handles large amounts of output and whether it correctly interleaves
stdout and stderr.

The simple example below shows a case where it is generating 16 lines of output
where each line has 32 characters, every 5th line is output to stderr, all others
are output to stdout and line numbers are added. The options used are briefly
described below. More detail can be found from the help (-h).

| Option | Function |
| ------ | -------- |
| -i NUM | Stderr interleave factor, output to stderr every NUM lines. |
| -l     | Prefix the line with a line count starting at 1. |
| -n NUM | The number of lines. |
| -w NUM | The number of characters on each line, includes the new line. |

```bash
$ gentestdata -w 32 -n 16 -i 5 -l
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
    13 zMd1JTT3ZGR5mEuJOaJCo9AZ
    14 mMTu3yTV0p7opMMsnA87D6TS
    15 TAXY5NACNjbsUfPoHYixe6pj
    16 0dHuKlxQyyNenUNQDSWPtW4u
```

When I redirect stdout to /dev/null so you only see the stderr output,
you only see 3 lines:

```bash
$ bin/gentestdata -w 32 -n 16 -l -i 5  1>/dev/null
     5 eDZ0tHBxFq6Pysq3N267L1vq
    10 j77Ax7Tlfj84Qyu6uRn8CTEC
    15 TAXY5NACNjbsUfPoHYixe6pj
```

When I redirect stderr to /dev/null so you only see the stdout output,
you only see 12 lines:

```bash
$ bin/gentestdata -w 32 -n 16 -l -i 5  2>/dev/null 
     1 bPlNFGdSC2wd8f2QnFhk5A84
     2 JJjKWZdKH9H2FHFuvUs9Jz8U
     3 vBHv3Vc5awx39ivuwsp2nChC
     4 IwVQztA2n95rXrtzhwuSAd6h
     6 kgnBsUje9FqBZonjaaWDcXMm
     7 8biABkerSuHpnMmMDF2EsjYy
     8 TQWCfIuilZxV2FCniRwo7StO
     9 fGOILa0u1wXnEw1GDGuvdSew
    11 WzT5s4ZJHd0TxrtMKykqOn91
    12 fMwNqsk2Wrc5uhk2kQaTXJp2
    13 zMd1JTT3ZGR5mEuJOaJCo9AZ
    14 mMTu3yTV0p7opMMsnA87D6TS
    16 0dHuKlxQyyNenUNQDSWPtW4u
```

One other thing, the number of bytes output is exact because the width (-w)
includes the newline character so you have full control over the number
of characters generated.

The example below demonstrates this by generating 1024 lines
each with 1024 characters each for a total 1048576.

```bash
$ bin/gentestdata -w 1024 -n 1024 -l | wc
    1024    2048 1048576
$ bin/gentestdata -w 1024 -n 1024 | wc
    1024    1024 1048576
$ bin/gentestdata -w 1024 -n 1024 -l | head -4 | cut -c -50
     1 bPlNFGdSC2wd8f2QnFhk5A84JJjKWZdKH9H2FHFuvUs
     2 nrIYoyi0Z7DC2VVQfBezSxCIrTF1uLSMty7al9Urols
     3 VkrxhmMfQHd10I4ok71ouNnVxSrYU2VuxUe36JKsBLl
     4 yyGlfngastMaGOuGh1k3RJTI7kAYvlujUbpH2eGctFb
```


## Building it
```bash
$ go version
$ git clone https://github.com/jlinoff/gentestdata.git
$ cd gentestdata
$ make
$ bin/gentestdata --version
```

I have only built it on Mac OS X (10.11.6) and CentOS 7.2 so you may have to modify the Makefile for your platform.

## Help
Help is available from the program. Just use the -h or --help option.

This is the output.

```bash
$ gendata -h

USAGE
    gentestdata [OPTIONS]

DESCRIPTION
    Program that generates text data for testing purposes.

    You can control the alphabet used, the number of lines, the line width and
    whether to output some or all of the lines to stderr.

    It is very useful to testing stderr handling in client/server systems.

OPTIONS
    -a STRING, --alphabet STRING
                       Specify an alternative alphabet to use for generating
                       the data.
                       The default is: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789".

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
    $ gentestdata -h

    # Example 2. short output
    $ gentestdata -n 4 -w 16
    bPlNFGdSC2wd8f2
    QnFhk5A84JJjKWZ
    dKH9H2FHFuvUs9J
    z8UvBHv3Vc5awx3

    # Example 3. short output with line numbers
    $ gentestdata -n 4 -w 16 -l
         1 bPlNFGdS
         2 C2wd8f2Q
         3 nFhk5A84
         4 JJjKWZdK

    # Example 4. interleave stderr so that every 4th line is printed to stderr
    $ # All output (stdout and stderr)
    $ gentestdata -n 12 -w 32 -l -i 4
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
    $ gentestdata -n 12 -w 32 -l -i 4 1>/dev/null
         4 IwVQztA2n95rXrtzhwuSAd6h
         8 TQWCfIuilZxV2FCniRwo7StO
        12 fMwNqsk2Wrc5uhk2kQaTXJp2

    $ # Only view the stdout output.
    $ gentestdata -n 12 -w 32 -l -i 4 2>/dev/null
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
    $ gentestdata -n 4 -w 16 -a 0123456789
    177918506041298
    415765688777805
    187196715630433
    784937199058835

    # Example 6. show the output size
    $ gentestdata -n 32 -w 32 -a 0123456789abcdef | wc
        32      32    1024

VERSION
    v0.2
```

