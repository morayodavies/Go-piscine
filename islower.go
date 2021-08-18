package piscine

func IsLower(s string) bool {
	w := []rune(s)
	for _, l := range w {
		if l < 'a' || l > 'z' {
			return false
		}
	}
	return true
}
