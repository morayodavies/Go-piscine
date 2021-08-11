package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	divisor := 1000000000000000000
	digit := 0
	i := 10
	if n < 0 {
		if n < -9223372036854775807 {
			z01.PrintRune('-')
			z01.PrintRune('9')
			n = 223372036854775808
		} else {
			z01.PrintRune('-')
			n = n * -1
		}
	}
	for divisor >= 1 {
		if n >= divisor {
			digit = (n / divisor) + 48
			z01.PrintRune(rune(digit))
			n = n - ((n / divisor) * divisor)
			for n < divisor/i {
				z01.PrintRune('0')
				i *= i
			}
			i = 10
		}
		divisor /= 10
	}
}
