package ReverseVowels

import "unicode"

func isVowel(letter rune) bool {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func reverseVowels(s string) string {
	var spread []rune = []rune(s)
	var indexFront int = 0
	var indexRear int = len(s) - 1
	for indexFront < indexRear {
		if !unicode.IsLetter(spread[indexFront]) || !isVowel(spread[indexFront]) {
			indexFront++
		} else if !unicode.IsLetter(spread[indexRear]) || !isVowel(spread[indexRear]) {
			indexRear--
		} else {
			spread[indexFront], spread[indexRear] = spread[indexRear], spread[indexFront]
			indexFront++
			indexRear--
		}
	}
	return string(spread)
}
