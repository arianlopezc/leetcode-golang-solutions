package QueueReconstructionByHeight

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	var res [][]int
	for i := 0; i < len(people); i++ {
		res = append(res, people[i])
		if people[i][1] < len(res) {
			copy(res[people[i][1]+1:len(res)], res[people[i][1]:len(res)-1])
			res[people[i][1]] = people[i]
		}
	}
	return res
}
