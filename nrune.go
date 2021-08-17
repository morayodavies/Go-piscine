package piscine

func NRune(s string, n int) rune {
	w := []rune(s)
	length := 0
	if s != "" {
		for i := range w {
			length = i
		}
	}
	if n-1 > length || n < 0 {
		return 0
	} else {
		return w[n-1]
	}
}
