package BackspaceStringCompare

func backspaceCompare(s string, t string) bool {
	var sIndex, tIndex = len(s) - 1, len(t) - 1
	for sIndex >= 0 && tIndex >= 0 {
		var sSkip, tSkip int
		for sIndex >= 0 && (s[sIndex] == '#' || sSkip > 0) {
			if s[sIndex] == '#' {
				sSkip++
			} else {
				sSkip--
			}
			sIndex--
		}
		for tIndex >= 0 && (t[tIndex] == '#' || tSkip > 0) {
			if t[tIndex] == '#' {
				tSkip++
			} else {
				tSkip--
			}
			tIndex--
		}
		if sIndex < 0 && tIndex < 0 {
			return true
		}
		if s[sIndex] != t[tIndex] {
			return false
		}
		sIndex--
		tIndex--
	}
	return true
}
