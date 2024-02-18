package WordPattern

import "strings"

func wordPattern(pattern string, s string) bool {
	split := strings.Split(s, " ")
	if len(split) != len(pattern) {
		return false
	}
	mapped := make(map[rune]string)
	mappedWord := make(map[string]rune)
	for index, letter := range pattern {
		if val, ok := mapped[letter]; ok {
			if val != split[index] || mappedWord[val] != letter {
				return false
			}
		} else if _, okW := mappedWord[split[index]]; okW {
			return false
		} else {
			mapped[letter] = split[index]
			mappedWord[split[index]] = letter
		}
	}
	return true
}
