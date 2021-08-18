package piscine

func ToLower(s string) string {
	w := []rune(s)
	for index, letter := range w {
		if letter >= 'A' && letter <= 'Z' {
			w[index] = letter + 32
		}
	}
	return string(w)
}
