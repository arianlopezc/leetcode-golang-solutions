package MaximumNumberVowelsSubstringGivenLength

func isVowel(letter rune) bool {
	if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
		return true
	}
	return false
}

func maxVowels(s string, k int) int {
	var conv = []rune(s)
	var totalVowels = 0
	var maxVowels = 0
	var added = k
	var innerIndex int
	var lastLetter = conv[innerIndex]
	for _, letter := range conv {
		if added > 0 {
			if isVowel(letter) {
				totalVowels++
			}
			added--
			if totalVowels > maxVowels {
				maxVowels = totalVowels
			}
		} else {
			if isVowel(lastLetter) {
				totalVowels--
			}
			innerIndex++
			lastLetter = conv[innerIndex]
			if isVowel(letter) {
				totalVowels++
			}
			if totalVowels > maxVowels {
				maxVowels = totalVowels
				if maxVowels == k {
					return maxVowels
				}
			}
		}
	}
	return maxVowels
}
