package algortm

func movingCount(m int, n int, k int) int {
	visited := make([]bool, m*n)
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if max < rs(m, n, k, i, j, visited) {
				max = rs(m, n, k, i, j, visited)
			}
		}
	}
	return max
}

func rs(m int, n int, k int, i, j int, visited []bool) int {
	if i >= 0 && i < m && j >= 0 && j < n && !visited[i*j+i] && (count(i)+count(j) <= k) {
		visited[i*j+i] = true
		return rs(m, n, k, i-1, j, visited) + rs(m, n, k, i+1, j, visited) +
			rs(m, n, k, i, j-1, visited) + rs(m, n, k, i, j+1, visited)
	}
	return 0

}

func count(i int) int {
	sum := 0
	for i >= 0 {
		temp := i % 10
		sum += temp
	}
	return sum
}
