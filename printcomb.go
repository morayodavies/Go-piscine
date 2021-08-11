package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var i rune = 48
	var j rune = 48
	var k rune = 48
	for i <= 57 {
		for j <= 57 {
			for k <= 57 {
				if i < j && j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i < 55 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				if k == 57 {
					k = j
					j++
				}
				if j == 57 {
					j = i
					i++
				}
				k++
			}
		}
	}
}
