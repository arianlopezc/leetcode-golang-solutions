package CanPlaceFlowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	var added int = 0
	var length int = len(flowerbed)
	for index, spot := range flowerbed {
		if spot == 0 {
			if (index-1 < 0 || flowerbed[index-1] == 0) && (index+1 >= length || flowerbed[index+1] == 0) {
				flowerbed[index] = 1
				added++
				if added == n {
					return true
				}
			}
		}
	}
	return false
}
