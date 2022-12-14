package leven

// https://github.com/m1ome/leven/blob/master/leven.go

// Distance determines the Levenshtein distance between two strings
func Distance(str1, str2 string) int {
	var cost, lastdiag, olddiag, p int

	s1 := []rune(str1) // массив символов 1 строки
	s2 := []rune(str2) // 2й строки

	len1 := len(s1) // длина 1 строки
	len2 := len(s2) // 2й строки

	for ; p < len1 && p < len2; p++ { // p свдигается до индекса первого различного символа
		if s2[p] != s1[p] {
			break
		}
	}
	s1, s2 = s1[p:], s2[p:] // 
	len1 -= p               // длина различающихся символов слова 1
	len2 -= p               // длина различающихся символов слова 2

	for 0 < len1 && 0 < len2 {
		if s1[len1-1] != s2[len2-1] {
			s1, s2 = s1[:len1], s2[:len2]
			break
		}
		len1--
		len2--
	}

	if len1 == 0 {
		return len2
	}

	if len2 == 0 {
		return len1
	}

	column := make([]int, len1+1)

	for y := 1; y <= len1; y++ {
		column[y] = y
	}

	for x := 1; x <= len2; x++ {
		column[0] = x
		lastdiag = x - 1
		for y := 1; y <= len1; y++ {
			olddiag = column[y]
			cost = 0
			if s1[y-1] != s2[x-1] {
				cost = 1
			}
			column[y] = min(
				column[y]+1,
				column[y-1]+1,
				lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return column[len1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
