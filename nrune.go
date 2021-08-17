package piscine

func NRune(s string, n int) rune {
	w := []rune(s)
	length := 0
	for i := range w {
		length = i + 1
	}
	if n > length {
		return 0
	} else {
		return w[n]
	}
}
