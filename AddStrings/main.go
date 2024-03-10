package AddStrings

import "strconv"

func addStrings(num1 string, num2 string) string {
	var numOne, numTwo = len(num1) - 1, len(num2) - 1
	var carryOver int
	var res string
	for numOne > -1 || numTwo > -1 {
		var first int
		if numOne > -1 {
			first, _ = strconv.Atoi(string(num1[numOne]))
		}
		var second int
		if numTwo > -1 {
			second, _ = strconv.Atoi(string(num2[numTwo]))
		}
		var newDigit = first + second + carryOver
		if newDigit >= 10 {
			carryOver = 1
			newDigit = newDigit % 10
		} else {
			carryOver = 0
		}
		res = strconv.Itoa(newDigit) + res
		numOne--
		numTwo--
	}
	if carryOver > 0 {
		res = strconv.Itoa(carryOver) + res
	}
	return res
}
