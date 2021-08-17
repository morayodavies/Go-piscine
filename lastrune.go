package piscine

func LastRune(s string) rune {
	w := []rune(s)
	lastrune := ''
	for index := range w {
		lastrune = w[index]
	}
	return lastrune
}