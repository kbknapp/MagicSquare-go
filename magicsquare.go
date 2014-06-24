package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
)

const (
	version = "0.1"
	usage   = `magicsquare <size> <prime only>
	-s, --size=<size>	The size of the square (i.e. '3' = 3x3 square)
	-p, --prime[=true]	Use only prime numbers in the square
	-v, --version		Display version information
	-h, --help			Display help information`
)

var vp = flag.Bool("version", false, "Display version information")
var hp = flag.Bool("help", false, "Display help information")
var sp = flag.Int("size", 3, "The size of the square (i.e. '3' = 3x3 square)")
var pp = flag.Bool("prime", false, "Use only prime numbers in the square")

func init() {
	flag.BoolVar(vp, "v", false, "Display version information")
	flag.BoolVar(hp, "h", false, "Display help information")
	flag.IntVar(sp, "s", 3, "The size of the square (i.e. '3' = 3x3 square)")
	flag.BoolVar(pp, "p", false, "Use only prime numbers in the square")
}

func check(err error) {
	_, file, line, _ := runtime.Caller(1)
	if err != nil {
		log.Fatalf("Fatal from <%s:%d>\nError:%s", file, line, err)
	}
}

func main() {
	flag.Parse()

	if *vp {
		fmt.Printf("\nMagicSquare v%s\n\n", version)
		return
	} else if *hp {
		fmt.Printf("\nMagicSquare v%s\n\n%s\n\n", version, usage)
		return
	}

	s := NewMSquare(*sp, *pp)

	if solved := s.Solve(); solved {
		fmt.Printf("Solved:\n")
	} else {
		fmt.Printf("No solution exists.\nFinal square:\n")
	}

	s.Print()
}
