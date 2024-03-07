package RemoveOutermostParentheses

func removeOuterParentheses(s string) string {
	var opening int
	var start = -1
	var res string
	for index := 0; index < len(s); index++ {
		if s[index] == '(' {
			opening++
			if start == -1 {
				start = index
			}
		} else {
			opening--
			if opening == 0 {
				res += s[start+1 : index]
				start = -1
			}
		}
	}
	return res
}
