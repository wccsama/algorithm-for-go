package algortm

// 一串首尾相连的珠子M个，有N个颜色取出其中一段包含所有的N个颜色的长度最短
func mn(a []int, n int) (int, int) {
	// check a n

	start := 0
	end := len(a) - 1
	min := len(a)
	colorMap := make(map[int]int)

	for i := 0; i < len(a); i++ {
		if a[i] == a[start] {
			start++
		}

		if len(colorMap) < n {
			if , ok := colorMap[a[i]]; !ok {
				colorMap[a[i]= i
				end = i
			}
		}

		if len(colorMap) == n {
			if min < end-start {
				min = end - start
				end = i
			}
		}

	}

	return end - start
}

func mn2(a []int, n int) (int, int) {
	start := 0
	end := len(a) - 1
	min := len(a)
	colorMap := make(map[int]int)

	for i := 0; i < len(a); i++ {
		if len(colorMap) < n {
			if , ok := colorMap[a[i]]; !ok {
				colorMap[a[i]= i
				end = i
			}
		}

		if len(colorMap) == n {
			// move start
			if _, ok := colorMap; ok {
				start++
			}

			// update  start/ end min

		}

	}
}
