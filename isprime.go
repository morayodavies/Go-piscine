package piscine

func checkprimes(nb int, s []int) bool {
	result := 0
	for _, n := range s {
		if nb%n == 0 && nb != n {
			result += 1
		} else {
			result += 0
		}
	}
	if result > 0 {
		return false
	} else {
		return true
	}
}
func Sqroot(nb int) int {
	sqroot := 0
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			sqroot = i
		}
	}
	return sqroot
}

func IsPrime(nb int) bool {
	s := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	sqrtnb := Sqroot(nb)
	if nb > 2000000000000 || nb <= 1 {
		return false
	} else if s[len(s)-1] >= sqrtnb {
		return checkprimes(nb, s)
	} else {
		for i := s[len(s)-1] + 1; (i < sqrtnb) && i < 3037000500; i++ {
			if checkprimes(i, s) {
				s = append(s, i)
			}
		}
		return checkprimes(nb, s)
	}
}
