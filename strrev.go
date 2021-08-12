package piscine

func StrRev(s string) string {
	length := 0
	for index := range s {
		length = index
	}
	word := []byte(s)
	for i := range s {
		word[i] = s[length-i]
	}
	stringword := string(word)
	return stringword
}
