package piscine

func IsLower(s string) bool {

	for l := range s {
		if l < 'a' || l > 'z' {
			return false
		}
	}
	return true
}
