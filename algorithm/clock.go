package algortm

// 顺时针打印，打印一圈，然后再打印一圈
func printTwo(a [][]int) {
	startrow := 0
	startl := 0
	row := len(a)
	l := len(a[0])
	for row >= 0 && l >= 0 {
		for i := startrow; i < row; i++ {
			// print
		}

		for i := startl + 1; i < l; i++ {

		}

		for i := l - 1; i > startl; i-- {

		}

		for i := row - 1; i > startrow; i-- {

		}
		startrow++
		startl++
		row--
		l--
	}
}
