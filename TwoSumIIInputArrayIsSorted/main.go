package TwoSumIIInputArrayIsSorted

func twoSum(numbers []int, target int) []int {
	var front, rear = 0, len(numbers) - 1
	for front < rear {
		if numbers[front]+numbers[rear] == target {
			return []int{front + 1, rear + 1}
		} else if numbers[front]+numbers[rear] > target {
			rear--
		} else {
			front++
		}
	}
	return nil
}
