package GuessNumberHigherLower

func guess(p int) int {
	if p == 2 {
		return 0
	} else if p > 2 {
		return -1
	} else {
		return 1
	}
}

func guessNumber(n int) int {
	l := 0
	r := n
	for l <= r {
		m := (l + r) / 2
		if check := guess(m); check == 1 {
			l = m + 1
		} else if check == -1 {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
