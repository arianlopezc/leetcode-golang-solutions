package CanPermutatePalyndrome

func canPermutePalindrome(s string) bool {
	var mapped = make(map[rune]int)
	for _, letter := range s {
		mapped[letter]++
	}
	var oddUsed = false
	for _, total := range mapped {
		if total%2 != 0 {
			if !oddUsed {
				oddUsed = true
			} else {
				return false
			}
		}
	}
	return true
}
