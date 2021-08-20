package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := range os.Args {
		if i < len(os.Args)-1 {
			for _, letter := range os.Args[len(os.Args)-i-1] {
				z01.PrintRune(letter)
			}
			z01.PrintRune('\n')
		}
	}
}
