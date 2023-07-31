package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args

	for _, i := range arg[0][2:] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
