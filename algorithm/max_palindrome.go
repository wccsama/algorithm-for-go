package algortm

func Max_Pal(a string) string {
	var s string
	for i := 0; i < len(a); i++ {
		s = append(s, "#")
		s = append(s, a[i])
	}
	s = append(s, "#")

	max := -1
	start := -1
	end := -1
	for i := 0; i < len(s); i++ {
		j := i - 1
		k := i + 1
		for j > 0 && k < len(s) {
			if s[j] == s[k] {
				if max < k-j {
					max = k - j
					start = j
					end = k
				}
			} else if {
				break
			}
		}
	}
}

// 思想同上，但是新增一个每个点的回文串的半径，以及回文边境
func manacher(a string) string {
	var s string
	for i := 0; i < len(a); i++ {
		s = append(s, "#")
		s = append(s, a[i])
	}
	s = append(s, "#")

	sl := make([]int, len(s))
	edge:=0 // 右侧边界
	pos := 0 // 右侧边界对应中心地址
	max := 0

	for i :=0 ;i < len(s); i++ {
		if i < edge {
			sl[i] = math.Min(a[s*pos - i ], a[edge- i])
		} else {
			sl[i] = 1	
		}

		for i - sl[i] >0 && i + sl[i] < len(s) && s[i - sl[i]] == s[i + sl[i]] {
			sl[i]++
		}

		if i-sl[i]+1 > edge {
			edge = i-sl[i]+1
			pos = i
		}

		if max < sl[i]{
			max = sl[i]
		}
	}

}
