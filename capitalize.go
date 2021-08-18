package piscine

func Capitalize(s string) string {
	w := []rune(s)
	for index, letter := range w {
		if index == 0 {
			if letter >= 'a' && letter <= 'z' {
				w[index] = letter - 32
			}
		} else {
			if !((w[index-1] >= 'a' && w[index-1] <= 'z') || (w[index-1] >= 'A' && w[index-1] <= 'Z') || (w[index-1] >= '0' && w[index-1] <= '9')) {
				if letter >= 'a' && letter <= 'z' {
					w[index] = letter - 32
				}
			} else {
				if letter >= 'A' && letter <= 'Z' {
					w[index] = letter + 32
				}
			}
		}
	}
	return string(w)
}
