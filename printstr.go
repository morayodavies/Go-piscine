package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	word := []rune(s)
	for _, letter := range word {
		z01.PrintRune(letter)
	}
}
