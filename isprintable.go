package piscine

func IsPrintable(s string) bool {
	w := []rune(s)
	for _, l := range w {
		if l >= 7 && l <= 13 {
			return false
		}
	}
	return true
}
