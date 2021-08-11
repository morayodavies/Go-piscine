package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= 57; i++ {
		for j := '0'; j <= 57; j++ {
			for k := '0'; k <= 57; k++ {
				for l := '0'; l <= 57; l++ {
					if k > i || (k == i && l > j) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						if !(i == '9' && j == '8') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
