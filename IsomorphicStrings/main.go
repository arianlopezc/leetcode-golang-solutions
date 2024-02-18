package IsomorphicStrings

func isIsomorphic(s string, t string) bool {
	var index = 0
	var length = len(s)
	var mappedS = make(map[rune]rune)
	var mappedT = make(map[rune]rune)
	for index < length {
		runeS := rune(s[index])
		runeT := rune(t[index])
		if val, ok := mappedS[runeS]; ok {
			if val != runeT {
				return false
			}
		} else {
			if _, ok := mappedT[runeT]; ok {
				return false
			}
			mappedS[runeS] = runeT
			mappedT[runeT] = runeS
		}
		index++
	}
	return true
}
