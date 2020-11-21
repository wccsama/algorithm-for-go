package sort

// 顺序数组，查找指定元素
func middleSearch(num []int, a int) int {
	left := 0
	right := len(num)

	for left <= right {
		middle := left + (right-left)/2
		if num[middle] == a {
			return middle
		}
		if num[middle] > a {
			right = middle - 1
		} else if num[middle] < a {
			left = middle + 1
		}
	}
	return -1
}
