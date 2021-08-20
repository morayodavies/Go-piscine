package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	minword := os.Args[1]
	length := len(os.Args[1:])
	c := 0
	d := 0
	var runes []rune

	for _, word := range os.Args[1:] {
		if word <= minword {
			minword = word
		}
	}
	for _, letter := range minword {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
	w := []string{minword}
	for len(w) <= len(os.Args[1:])-1 {
		for _, word := range os.Args[1:] {
			for _, minwords := range w {
				if word >= minwords {
					c += 1
				}
			}
			for _, comwords := range os.Args[1:] {
				if word <= comwords {
					d += 1
				}
			}
			if c == len(w) && d == length-len(w) {
				runes = []rune(word)
				for _, rune := range runes {
					z01.PrintRune(rune)
				}
				z01.PrintRune('\n')
				w = append(w, word)
			}
			c = 0
			d = 0
		}
	}
}
