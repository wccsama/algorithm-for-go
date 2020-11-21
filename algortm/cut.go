package algortm

func max(a []int) int {
	result := make([]int, len(a))
	if len(a) == 0 {
		return 0
	}

	result[1] = a[1]
	result[2] = a[2]
	result[3] = a[1] * a[2]

	for i := 4; i < len(a); i++ {
		for j := 1; j < i/2; j++ {
			result[i] = max(result[j] * result[i-j])
		}
	}

}
