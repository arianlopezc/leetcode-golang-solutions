package IsSubsequence

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	var indexS int
	var indexT int
	var sRunes []rune = []rune(s)
	for _, letter := range []rune(t) {
		if letter == sRunes[indexS] {
			indexS++
			if indexS == len(s) {
				return true
			}
		}
		indexT++
	}
	return false
}
