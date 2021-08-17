package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 0; i+1 <= nb; i++ {
		result *= (i + 1)
	}
	return result
}
