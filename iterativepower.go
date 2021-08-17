package piscine

func IterativePower(nb int, power int) int {
	if power > 0 {
		result := nb
		for i := 1; i < power; i++ {
			result *= nb
		}
		return result
	} else if power == 0 {
		return 1
	} else {
		return 0
	}
}
