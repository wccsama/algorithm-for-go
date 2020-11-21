package sort

func quickSort(nums []int, left, right int) {
	i := left
	j := right
	key := nums[i]

	for i < j {
		for i < j && nums[i] < key {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}

		for i < j && nums[j] >= key {
			j--
		}
		if i < j {
			nums[i] = nums[j]
			i++
		}
	}
	nums[i] = key
	// quickSort
	//
}

func quick(num []int, l, r int) {
	if r >= l {
		return
	}
	i := l
	j := r
	key := num[i]
	for i < j {
		for i < j && num[j] >= key {
			j--
		}
		if i < j {
			num[i] = num[j]
			i++
		}

		for i < j && num[i] < key {
			i++
		}
		if i < j {
			num[j] = num[i]
			j--
		}
	}
	num[i] = key

	//
}

func quick(a []int, l, r int) {
	if r >= l {
		return
	}
	pit := a[l]
	i := l
	j := r
	for i < j {
		for i < j && a[j] >= pit {
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}

		for i < j && a[i] < pit {
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	a[i] = pit
	quick(a, l, i-1)
	quick(a, i+1, r)
}

func quick2(a []int, s, e int) {
	if s >= e {
		return
	}

	i := s
	j := e
	pix := a[s]
	for i < j {
		for i < j && a[j] >= pix {
			j--
		}
		if j > i {
			a[i] = a[j]
			i++
		}

		for i < j && a[i] < pix {
			i++
		}
		if j > i {
			a[j] = a[i]
			j--
		}
	}
	a[i] = pix
	quick2(a, s, i-1)
	quick2(a, i+1, e)
}
