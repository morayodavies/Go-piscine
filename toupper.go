package piscine

func ToUpper(s string) string {
	w := []rune(s)
	for index, letter := range w {
		if letter >= 'a' && letter <= 'z' {
			w[index] = letter - 32
		}
	}
	return string(w)
}
