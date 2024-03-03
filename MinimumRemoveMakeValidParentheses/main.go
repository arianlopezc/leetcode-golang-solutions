package MinimumRemoveMakeValidParentheses

import "strings"

func minRemoveToMakeValid(s string) string {
	var pTrack []int
	var splited = []rune(s)
	for index := 0; index < len(splited); index++ {
		if splited[index] == '(' {
			pTrack = append(pTrack, index)
		} else if splited[index] == ')' {
			if len(pTrack) == 0 {
				splited[index] = '_'
			} else {
				pTrack = pTrack[:len(pTrack)-1]
			}
		}
	}
	for index := len(pTrack) - 1; index >= 0; index-- {
		splited[pTrack[index]] = '_'
	}
	builder := strings.Builder{}
	for _, val := range splited {
		if val != '_' {
			builder.WriteRune(val)
		}
	}
	return builder.String()
}
