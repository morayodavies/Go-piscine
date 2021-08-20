package main

import (
	"os"

	"github.com/01-edu/z01"
)

func TrimAtoi(s string) int {
	w := []rune(s)
	negindex := -1
	firstnumbindex := -1
	negfound := false
	numfound := false
	nbr := 0
	var digits []rune
	for index, character := range w {
		if character == '-' && !negfound {
			negindex = index
			negfound = true
		}
		if (character >= '0' && character <= '9') && !numfound {
			firstnumbindex = index
			numfound = true
		}
		if character >= '0' && character <= '9' {
			digits = append(digits, character)
		}
	}
	Lenw := 10
	for i := range digits {
		if i == 0 {
			Lenw = 1
		} else if i == 1 {
			Lenw = 10
		} else {
			Lenw *= 10
		}
	}
	for _, digit := range digits {
		nbr += (int(digit) - 48) * Lenw
		Lenw /= 10
	}
	if negindex < firstnumbindex && negindex >= 0 {
		return nbr * -1
	} else {
		return nbr
	}
}

func main() {
	if os.Args[1] == "--upper" {
		for _, letter := range os.Args[2:] {
			if TrimAtoi(letter) >= 1 && TrimAtoi(letter) <= 26 {
				z01.PrintRune(rune(TrimAtoi(letter) + 64))
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	} else {
		for _, letter := range os.Args[1:] {
			if TrimAtoi(letter) >= 1 && TrimAtoi(letter) <= 26 {
				z01.PrintRune(rune(TrimAtoi(letter) + 96))
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
