package FirstUniqueChar

func firstUniqChar(s string) int {
	var matches = make(map[rune]int)
	for _, letter := range s {
		matches[letter]++
	}
	for index, letter := range s {
		if matches[letter] == 1 {
			return index
		}
	}
	return -1
}
