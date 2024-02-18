package MergeStringsAlternatively

func mergeAlternately(word1 string, word2 string) string {
	var lengthA = len(word1)
	var lengthB = len(word2)
	if lengthA == 0 {
		return word2
	}
	if lengthB == 0 {
		return word1
	}
	var result = make([]rune, len(word1)+len(word2))
	var takeA = true
	var indexA = 0
	var indexB = 0
	for index := range result {
		if takeA && indexA < lengthA || !takeA && indexB == lengthB {
			result[index] = rune(word1[indexA])
			indexA++
		} else {
			if indexB < lengthB {
				result[index] = rune(word2[indexB])
				indexB++
			}
		}
		takeA = !takeA
	}
	return string(result)
}
