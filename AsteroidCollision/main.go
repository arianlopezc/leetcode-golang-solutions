package AsteroidCollision

func asteroidCollision(asteroids []int) []int {
	var front, rear, length = 1, 0, len(asteroids)
	var destroyedAsteroids int
	for front < length {
		if asteroids[front] < 0 {
			rear = front - 1
			for rear > 0 && asteroids[rear] <= 0 {
				rear--
			}
			if asteroids[rear] > 0 {
				if absFront := asteroids[front] * -1; asteroids[rear] > absFront {
					asteroids[front] = 0
					front++
					destroyedAsteroids++
				} else if asteroids[rear] < absFront {
					asteroids[rear] = 0
					destroyedAsteroids++
				} else {
					asteroids[front] = 0
					asteroids[rear] = 0
					front++
					destroyedAsteroids += 2
				}
			} else {
				front++
			}
		} else if asteroids[front] >= 0 {
			front++
		}
	}
	remainingAsteroids := make([]int, length-destroyedAsteroids)
	index := 0
	for _, val := range asteroids {
		if val != 0 {
			remainingAsteroids[index] = val
			index++
		}
	}
	return remainingAsteroids
}
