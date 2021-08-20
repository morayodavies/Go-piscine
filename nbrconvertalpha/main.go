package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	if os.Args[1] == "--upper" {
		for _, letter := range os.Args[2:] {
			if piscine.TrimAtoi(letter) >= 1 && piscine.TrimAtoi(letter) <= 26 {
				z01.PrintRune(rune(piscine.TrimAtoi(letter) + 64))
			} else {
				z01.PrintRune(' ')
			}
		}
	} else {
		for _, letter := range os.Args[1:] {
			if piscine.TrimAtoi(letter) >= 1 && piscine.TrimAtoi(letter) <= 26 {
				z01.PrintRune(rune(piscine.TrimAtoi(letter) + 96))
			} else {
				z01.PrintRune(' ')
			}
		}
	}
}
