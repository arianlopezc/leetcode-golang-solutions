package IsPalindrome

import "unicode"

func isPalindrome(s string) bool {
	var letters []rune
	for _, letter := range s {
		if unicode.IsLetter(letter) || unicode.IsNumber(letter) {
			letters = append(letters, unicode.ToLower(letter))
		}
	}
	var total int = len(letters)
	for index, _ := range letters {
		if letters[index] != letters[total-1-index] {
			return false
		}
	}
	return true
}
