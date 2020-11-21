package sort

// 稳定
func insert(nums []int) {
	for i := 1; i < len(nums); i++ {
		current := nums[i]
		j := i - 1
		for j > 0 && nums[j] > current {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = current
	}
}


