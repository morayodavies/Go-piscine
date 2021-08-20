package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := []rune(os.Args[0])
	for i, l := range s {
		if i > 1 {
			z01.PrintRune(l)
		}
	}
	z01.PrintRune('\n')
}
