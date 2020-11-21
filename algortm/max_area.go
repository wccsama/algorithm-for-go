package algortm

import "math"

var num int

// 岛屿中的陆地都是1
func maxLand(a [][]int) int {
	max := 0
	flag := make([][]bool, 0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if flag[i][j] {
				continue
			}
			dfs(a, flag, i, j)
			max = math.Max(num, max)
		}
	}
	return max
}

func dfs(a [][]int, flag [][]bool, i, j int) {
	if i < 0 || i > len(a) || j < 0 || j > len(a[0]) || a[i][j] == 0 {
		return
	}
	num++
	flag[i][j] = true
	dfs(a, flag, i+1, j)
	dfs(a, flag, i-1, j)
	dfs(a, flag, i, j+1)
	dfs(a, flag, i, j-1)
}
