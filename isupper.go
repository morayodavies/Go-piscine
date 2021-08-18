package piscine

func IsUpper(s string) bool {
	w := []rune(s)
	for _, l := range w {
		if l < 'A' || l > 'Z' {
			return false
		}
	}
	return true
}
