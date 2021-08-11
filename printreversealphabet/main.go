package main

import "github.com/01-edu/z01"

func main() {
	i := 'z'

	for i >= 97 {
		z01.PrintRune(i)
		i--
	}
	z01.PrintRune('\n')
}
