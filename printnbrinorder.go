package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var digits []rune
	remainder := 0
	for number := n; number >= 10; number /= 10 {
		remainder = number % 10
		digits = append(digits, rune(remainder+48))
		n /= 10
	}
	digits = append(digits, rune(n+48))

	for i := '0'; i <= '9'; i++ {
		for _, letter := range digits {
			if letter == i {
				z01.PrintRune(letter)
			}
		}
	}
}
