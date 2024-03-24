package DotProductOfTwoSparseVectors

type SparseVector struct {
	Nums []int
}

func Constructor(nums []int) SparseVector {
	return SparseVector{Nums: nums}
}

func (this *SparseVector) dotProduct(vec SparseVector) int {
	var vecNums = vec.Nums
	var total, index int
	for index < len(this.Nums) {
		var curr = vecNums[index] * this.Nums[index]
		total += curr
		index++
	}
	return total
}
