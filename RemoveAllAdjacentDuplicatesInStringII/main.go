package RemoveAllAdjacentDuplicatesInStringII

func removeDuplicates(s string, k int) string {
	var changed = true
	var word = []rune(s)
	for changed {
		changed = false
		if len(word) > 0 {
			var subK = 1
			var letter = word[0]
			for index := 1; index < len(word) && !changed; index++ {
				if word[index] == letter {
					subK++
					if subK == k {
						word = append(word[:(index-k)+1], word[index+1:]...)
						changed = true
					}
				} else {
					letter = word[index]
					subK = 1
				}
			}
		}
	}
	return string(word)
}
