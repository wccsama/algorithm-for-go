package sort

// 不稳定
// n2
func selectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if min < nums[j] {
				min = nums[j]
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
