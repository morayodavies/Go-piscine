package piscine

func BasicAtoi(s string) int {
	wordlength := 0
	word := []rune(s)
	for index := range word {
		wordlength = index
	}
	i := 10
	n := 0
	for index := range word {
		if index > 0 {
			lastletter := word[wordlength-index]
			e := int(lastletter) - 48
			n = n + (e * i)
			i *= 10
		} else {
			n = int(word[wordlength]) - 48
		}
	}
	return n
}
