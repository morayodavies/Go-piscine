package piscine

func LengthIndex(i string) int {
	w := []rune(i)
	length := 0
	for index := range w {
		length = index
	}
	return length
}

func Index(s string, toFind string) int {
	word := []rune(s)
	substring := []rune(toFind)
	c := 0
	if LengthIndex(toFind) > LengthIndex(s) {
		return -1
	} else {
		for i := 0; i <= LengthIndex(s)-LengthIndex(toFind); i++ {
			for j := 0; j <= LengthIndex(toFind); j++ {
				if word[j+i] == substring[j] {
					c += 1
				}
				if c == LengthIndex(toFind)+1 {
					return i + j - LengthIndex(toFind)
				}
			}
		}
		return -1
	}
}
