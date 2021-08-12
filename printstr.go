package piscine

import "fmt"

func PrintStr(s string) {
	word := []rune(s)
	for _, letter := range word {
		fmt.Printf("%v", string(letter))
	}
}
