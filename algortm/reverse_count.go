package algortm

func recount(a []int, low, high int) {
	if low >= high {
		return
	}
	mer(a, middle, high)

}

func mer(a []int, s, e int) int {
	if s > e {
		return 0
	}

	middle := (e-s)/2 + s
	left := mer(a, s, middle-1)
	right := mer(a, middle, e)

	i := middle - 1
	j := end
	temp := make([]int, e-s)
	lenTemp := len(temp)
	count := 0

	// 赋值
	for i >= s && j >= middle {
		if a[i] > a[j] {
			count += j - middle
			temp[lenTemp--] = a[i--]
		}else {
			temp[lenTemp--] = a[j--]
		}
	} 

	 // 通常赋值

	return left + right + count
}
