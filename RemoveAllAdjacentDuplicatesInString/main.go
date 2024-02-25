package RemoveAllAdjacentDuplicatesInString

func removeDuplicates(s string) string {
	if len(s) < 2 {
		return s
	}
	var removed = make(map[int]bool)
	var front, rear = 1, 0
	for front < len(s) {
		if s[front] == s[rear] {
			removed[front] = true
			removed[rear] = true
			for removed[front] && front < len(s) {
				front++
			}
			for removed[rear] && rear >= 0 {
				rear--
			}
			if rear < 0 {
				rear, front = front, front+1
			}
		} else {
			for removed[front] && front < len(s) {
				front++
			}
			front++
			rear = front - 1
		}
	}
	var result []byte
	for index := 0; index < len(s); index++ {
		if !removed[index] {
			result = append(result, s[index])
		}
	}
	return string(result)
}
