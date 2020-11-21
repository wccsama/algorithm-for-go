package algortm

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	queue := make([]int, 0, k)
	result := make([]int, len(nums)-k+1)
	countR := 0
	for i := 0; i < len(nums); i++ {
		if i-k >= 0 && len(queue) > 0 {
			if nums[i-k] == queue[0] {
				queue = queue[1:]
			}

		}
		for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		if i >= k-1 {
			result[countR] = queue[0]
			countR++
		}
	}
	return result
}
