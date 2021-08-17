package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 || nb > 25 {
		return 0
	} else {
		result := 1
		for i := 0; i+1 <= nb; i++ {
			result *= (i + 1)
		}
		return result
	}
}
