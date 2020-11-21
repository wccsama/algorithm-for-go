package sort

func mergeSort(num []int) []int {
	if len(num) < 2 {
		return num
	}
	middle := len(num) / 2
	left := num[0:middle]
	right := num[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(a, b []int) []int {
	result := make([]int, 0)
	for len(a) != 0 && len(b) != 0 {
		if a[0] < b[0] {
			result = append(result, a[0])
			a = a[1:]
		} else {
			result = append(result, b[0])
			b = b[1:]
		}
	}

	for len(a) != 0 {
		result = append(result, a[0])
		a = a[1:]
	}

	for len(b) != 0 {
		result = append(result, b[0])
		a = b[1:]
	}
	return result
}

// 多路归并
func multiple(a [][]int, s, e int) []int {
	if s >= e {
		return a[s]
	}

	middle := (s + e) / 2
	left := multiple(a, s, middle)
	right := multiple(a, middle+1, e)
	return merge(left, right)
}

func mergeSort2(a []int) []int {
	if len(a) == 0 {
		return nil
	}

	left := 0
	right := len(a)
	middle := (right-left)/2 + left
	return merge(mergeSort2(a[left:middle]), mergeSort2(a[middle:right]))
}
