package piscine

func checkprimes(nb int, s []int) bool {
	result := 0
	for _, n := range s {
		if nb%n == 0 {
			result += 1
		} else {
			result += 0
		}
	}
	if result > 0 && nb != 2 {
		return false
	} else {
		return true
	}
}

func IsPrime(nb int) bool {
	s := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	if nb > 200000 || nb <= 1 {
		return false
	} else if nb <= s[len(s)-1]*s[len(s)-1] {
		return checkprimes(nb, s)
	} else {
		for i := s[len(s)-1] + 1; (i < nb) && i < s[len(s)-1]*s[len(s)-1]; i++ {
			if checkprimes(i, s) {
				s = append(s, i)
			}
		}
		return checkprimes(nb, s)
	}
}
