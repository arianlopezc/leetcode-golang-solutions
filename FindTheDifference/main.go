package FindTheDifference

func findTheDifference(s string, t string) byte {
	rs := []rune(s)
	rt := []rune(t)
	theRune := rt[len(s)]

	for i, _ := range rs {
		theRune -= rs[i]
		theRune += rt[i]
	}

	return byte(theRune)
}
