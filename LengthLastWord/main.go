package LengthLastWord

import "unicode"

func lengthOfLastWord(s string) int {
	var separated = []rune(s)
	var endsIn = len(separated) - 1
	var lastWord = -1
	for endsIn >= 0 && lastWord == -1 {
		if unicode.IsLetter(separated[endsIn]) {
			lastWord = endsIn
		}
		endsIn--
	}
	var count int
	for lastWord >= 0 {
		if !unicode.IsLetter(separated[lastWord]) {
			return count
		} else {
			count++
			lastWord--
		}
	}
	return count
}
