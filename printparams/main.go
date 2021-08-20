package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i, word := range os.Args {
		if i >= 1 {
			for _, letter := range word {
				z01.PrintRune(letter)
			}
			z01.PrintRune('\n')
		}
	}
}
