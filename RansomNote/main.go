package RansomNote

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int)
	for _, letter := range magazine {
		letters[letter]++
	}
	for _, letter := range ransomNote {
		letters[letter]--
		if total := letters[letter]; total < 0 {
			return false
		}
	}
	return true
}
