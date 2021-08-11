package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	divisor := 1000000000000000000
	res := 0
	if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	}
	for divisor > 1 {
		if n >= divisor {
			res = (n / divisor) + 48
			z01.PrintRune(rune(res))
			n = n - ((n / divisor) * divisor)
		}
		divisor /= 10
	}
	z01.PrintRune(rune(n + 48))
}
