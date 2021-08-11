package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	res := 0
	remainder := 0
	if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	}
	if n >= 100 {
		z01.PrintRune('1')
		n -= 100
		remainder = n % 10
	}
	if n >= 10 {
		res = (n/10 + 48)
		remainder = n % 10
		z01.PrintRune(rune(res))
		n = remainder
	}
	if n >= 0 {
		remainder += 48
		z01.PrintRune(rune(remainder))
	}
}
