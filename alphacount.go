package piscine

func AlphaCount(s string) int {
	w := []rune(s)
	counter := 0
	for _, l := range w {
		if (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') {
			counter += 1
		}
	}
	return counter
}
