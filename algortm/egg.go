package algortm

func superEggDrop(K int, N int) int {
	if K == 1 {
		return N
	}

	if N == 0 {
		return 0
	}

	count := 0
	for i := 0; i <= N; i++ {
		count = min(max(superEggDrop(K-1, N-1), superEggDrop(K, N-i))+1, count)
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
