package LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	var separated = []rune(s)
	var longest, currentLongest int
	var mapped = make(map[rune]bool)
	var front, rear = 0, 0
	var length = len(s)
	for {
		if rear == length {
			if currentLongest > longest {
				longest = currentLongest
			}
			return longest
		}
		if !mapped[separated[rear]] {
			mapped[separated[rear]] = true
			currentLongest++
			if currentLongest > longest {
				longest = currentLongest
			}
			rear++
		} else {
			mapped[separated[front]] = false
			front++
			currentLongest--
		}
	}
}
