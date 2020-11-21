package algortm

const gMax = 6
func twoSum(n int) []float64 {
	total := math.Pow(float(gMax), float64(n))
	flag = 0
	result := make([][]int, 2)

	for i:=1; i= < gMax; i++ {
		result[flag][i] = 1
	}

	for k = 2; k <= n; i++ {
		result[1-flag][i] = 0
		for i := k; i <= gMax*n; i++ {
			result[1-flag][i] = 0
			for j := 1; i -j >=0 && j < gMax; j++ {
				result[1-flag][i] += result[1-flag][i-j]
			}
		}
		flag = 1- flag
	}


}