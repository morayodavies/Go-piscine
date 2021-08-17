package piscine

func NRune(s string, n int) rune {
	position := n - 1
	w := []rune(s)
	length := 0
	if s != "" {
		for i := range w {
			length = i
		}
	}
	if position > length || n <= 0 {
		return 0
	} else {
		return w[position]
	}
}
