package ValidWordAbbreviation

import (
	"strconv"
	"strings"
	"unicode"
)

func validWordAbbreviation(word string, abbr string) bool {
	abbrLength := len(abbr)
	wordLength := len(word)
	var wIndex, aIndex = 0, 0
	for wIndex < wordLength && aIndex < abbrLength {
		if abbr[aIndex] == '0' {
			return false
		} else if unicode.IsDigit(rune(abbr[aIndex])) {
			var val strings.Builder
			for aIndex < abbrLength && unicode.IsDigit(rune(abbr[aIndex])) {
				val.WriteByte(abbr[aIndex])
				aIndex++
			}
			if n, err := strconv.Atoi(val.String()); err != nil {
				return false
			} else {
				wIndex += n
			}
		} else if abbr[aIndex] == word[wIndex] {
			wIndex++
			aIndex++
		} else {
			return false
		}
	}
	return wIndex == wordLength && aIndex == abbrLength
}
