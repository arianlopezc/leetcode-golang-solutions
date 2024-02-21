package FibonacciNumber

func fib(n int) int {
	if n == 0 {
		return 0
	}
	var front, rear = 1, 0
	for n > 1 {
		front, rear = front+rear, front
		n--
	}
	return front
}
