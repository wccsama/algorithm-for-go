package algortm

func min(a []int, k) {
	index := p(a, 0, len(a))
	for index != k-1{
		if index < k -1 {
			low = index + 1
		}else {
			high = index -1
		}
		index = p(a, low, high)

	}

	return
}

func p(a []int, s, end int) int{
	start := a[s]
	for s < end {
		for s < end && a[s] <= start {
			s++
		}
		if a[s] > start {
			a[end]= a[s]
			end--
		}

		for s < end && a[end] > start {
			end--
		}
		if a[end] < start {
			a[s] = a[end]
			s++
		}
	}
	a[s] = start
	return start
}