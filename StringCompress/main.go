package StringCompress

import "strconv"

func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}
	var totalBytesNeeded = 0
	var overIndex = 0
	var index = 0
	for index < len(chars) {
		totalBytesNeeded++
		var currChart = chars[index]
		var currTotal = 0
		for index < len(chars) && chars[index] == currChart {
			index++
			currTotal++
		}
		chars[overIndex] = currChart
		overIndex++
		if currTotal > 1 {
			var converted = strconv.Itoa(currTotal)
			for _, digit := range converted {
				totalBytesNeeded++
				chars[overIndex] = byte(digit)
				overIndex++
			}
			currTotal = 0
		}
	}
	return totalBytesNeeded
}
