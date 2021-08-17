package piscine

func checkprimes1(nb int, s []int) bool {
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

func Sqroot1(nb int) int {
	sqroot := 0
	for i := 1; i <= nb; i++ {
		if i*i >= nb {
			sqroot = i
			return sqroot
		}
	}
	return sqroot
}

func IsPrime1(nb int) bool {
	s := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	sqrtnb := Sqroot1(nb)
	if nb > 2000000000000 || nb <= 1 {
		return false
	} else if s[len(s)-1] >= sqrtnb {
		return checkprimes1(nb, s)
	} else {
		for i := s[len(s)-1] + 1; (i < sqrtnb) && i < 3037000500; i++ {
			if checkprimes1(i, s) {
				s = append(s, i)
			}
		}
		return checkprimes1(nb, s)
	}
}

func FindNextPrime(nb int) int {
	for i := nb; i < 9223372036854775807; i++ {
		if IsPrime1(i) {
			return i
		}
	}
	return 0
}
