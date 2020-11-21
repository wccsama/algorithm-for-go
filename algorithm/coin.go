package algortm

func waysToChange(n int) int {
	coins := []int{0, 1, 5, 10, 25}
	result := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		result[i] = make([]int, n+1)
	}
	result[0][0] = 1

	for i := 1; i < len(coins); i++ {
		for j := 0; j <= n; j++ {
			if j-coins[i] >= 0 {
				result[i][j] = (result[i-1][j] + result[i][j-coins[i]]) % 1000000007
			} else {
				result[i][j] = result[i-1][j] % 1000000007
			}
		}
	}
	return result[4][n]
}
