package piscine

func Sqrt(nb int) int {
	sqrt := 0
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			sqrt = i
		}
	}
	return sqrt
}
