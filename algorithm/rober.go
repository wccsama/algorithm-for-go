package algortm

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	result := make([]int, len(nums))
	result[0] = nums[0]
	result[1] = max(nums[0], nums[1])
	maxN := result[1]

	for i := 2; i < len(nums); i++ {
		result[i] = max(result[i-1], result[i-2]+nums[i])
		maxN = max(maxN, result[i])
	}

	return maxN
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
