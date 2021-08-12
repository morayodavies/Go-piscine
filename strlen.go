package piscine

func StrLen(s string) int {
	word := []rune(s)
	length := 0
	for index := range word {
		length = index + 1
	}
	return length
}
