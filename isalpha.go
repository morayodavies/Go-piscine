package piscine

func IsAlpha(s string) bool {
	w := []rune(s)
	for _, l := range w {
		if !((l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') || (l >= '0' && l <= '9')) {
			return false
		}
	}
	return true
}
