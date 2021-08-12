package piscine

func BasicAtoi2(s string) int {
	returnvalue := true
	wordlength := 0
	word := []rune(s)
	for index := range word {
		wordlength = index
	}
	i := 10
	n := 0
	lastletter := rune(0)
	for index := range word {
		if index > 0 {
			lastletter = word[wordlength-index]
			e := int(lastletter) - 48
			n = n + (e * i)
			i *= 10
		} else {
			lastletter = word[wordlength]
			n = int(word[wordlength]) - 48
		}
		if lastletter > 57 || lastletter < 48 {
			returnvalue = false
		}
	}
	if returnvalue {
		return n
	} else {
		return 0
	}
}
