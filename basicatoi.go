package piscine

func BasicAtoi(s string) int {
	wordlength := 0
	word := []rune(s)
	for index := range word {
		wordlength = index
	}
	n := int(word[wordlength]) - 48
	i := 10
	for index := range word {
		if index > 0 {
			lastletter := word[wordlength-index]
			e := int(lastletter) - 48
			n = n + (e * i)
			i *= 10
		}
	}
	return n
}
