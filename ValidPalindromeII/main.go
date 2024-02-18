package ValidPalindromeII

func palidrome(runes []rune) bool {
	var fI, rI = 0, len(runes) - 1
	for fI < rI {
		if runes[fI] != runes[rI] {
			return false
		}
		fI++
		rI--
	}
	return true
}

func validPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	var fI, rI = 0, len(s) - 1
	var switched = false
	var runes = []rune(s)
	for fI < rI {
		if runes[fI] != runes[rI] {
			if switched {
				return false
			} else {
				return palidrome(runes[fI+1:rI+1]) || palidrome(runes[fI:rI])
			}
		}
		fI++
		rI--
	}
	return true
}
