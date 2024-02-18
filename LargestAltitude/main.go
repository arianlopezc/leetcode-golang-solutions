package LargestAltitude

func largestAltitude(gain []int) int {
	var highest int = 0
	var currentAltitude int = 0
	for _, change := range gain {
		currentAltitude += change
		if currentAltitude > highest {
			highest = currentAltitude
		}
	}
	return highest
}
