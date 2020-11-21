package sort

// 稳定排序
// n2
func bubble(nums []int32) {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}

func quick2(nums []int32, s, e int) {
	i, j := s, e
	pix := nums[i]
	for i < j {
		for i < j && nums[j] > pix {
			j--
		}

		if i < j {
			nums[i] = nums[j]
			i++
			j--
		}

		for i < j && nums[i] <= pix {
			i++
		}

		if i < j {
			nums[j] = nums[i]
			j--
			i++
		}
	}
	nums[i] = pix
	quick2(nums, s, i)
	quick2(nums, i+1, e)
}
