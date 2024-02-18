package ReverseWords

import (
	"unicode"
)

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func reverseWords(s string) string {
	var currentWord []rune
	var fullWord = []rune(s)
	var reversedWord []rune
	var nonLetters []rune
	for i := len(fullWord) - 1; i >= 0; i-- {
		if unicode.IsLetter(fullWord[i]) || unicode.IsNumber(fullWord[i]) {
			if len(nonLetters) > 0 {
				if len(reversedWord) > 0 {
					reversedWord = append(reversedWord, nonLetters...)
				}
				nonLetters = make([]rune, 0)
			}
			currentWord = append(currentWord, fullWord[i])
		} else {
			if len(nonLetters) == 0 {
				nonLetters = append(nonLetters, fullWord[i])
			}
			currentWord = reverse(currentWord)
			reversedWord = append(reversedWord, currentWord...)
			currentWord = make([]rune, 0)
		}
	}
	if len(currentWord) > 0 {
		currentWord = reverse(currentWord)
		reversedWord = append(reversedWord, currentWord...)
	}
	return string(reversedWord)
}
