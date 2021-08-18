package piscine

func IsNumeric(s string) bool {
	w := []rune(s)
	for _, l := range w {
		if !(l >= '0' && l <= '9') {
			return false
		}
	}
	return true
}
