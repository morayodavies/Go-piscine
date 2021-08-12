package piscine

func Atoi(s string) int {
	returnvalue := true
	wordlength := 0
	word := []rune(s)
	for index := range word {
		wordlength = index
	}
	i := 10
	n := 0
	firstletter := word[0]
	sign := 1
	if firstletter == 45 {
		sign = -1
	} else {
		sign = 1
	}
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
		if (lastletter > 57 || lastletter < 48) && index > 0 {
			returnvalue = false
		}
	}
	if returnvalue {
		return n * sign
	} else {
		return 0
	}
}
