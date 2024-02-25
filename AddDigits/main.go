package AddDigits

import (
	"strconv"
	"strings"
)

func addDigits(num int) int {
	var currVal = num
	for currVal > 9 {
		var temp int
		var splitted = strings.Split(strconv.Itoa(currVal), "")
		for _, val := range splitted {
			if v, err := strconv.Atoi(val); err == nil {
				temp += v
			}
		}
		currVal = temp
	}
	return currVal
}
