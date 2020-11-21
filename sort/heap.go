package sort

func buildHeap(a []int, n int) {
	for i := (n - 1) / 2; i >= 0; i-- {
		MinHeapFixdown(a, i, n)
	}
}

func MinHeapFixdown(a []int, i, n int) {
	j := 2*i + 1
	for j < n {
		if j+1 < n && a[j+1] < a[j] {
			j++
		}

		if a[i] <= a[j] {
			break
		}
		a[i], a[j] = a[i], a[j]
		i = j
		j = 2*i + 1
	}
}
