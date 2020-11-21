package algortm

//  1 1
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	for i := 0; i < len(nums); i++ {
		if i+1 == nums[i] {
			continue
		}
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]

		}
	}

	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return len(nums) + 1
}
