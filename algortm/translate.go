package algortm

import (
	"strconv"
)

// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
// 示例 1:

// 输入: 12258
// 输出: 5
// 解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

func translateNum(num int) int {
	s := strconv.Itoa(num)
	if len(s) == 1 {
		return 1
	}
	if len(s) == 2 && s >= "10" && s <= "25" {
		return 2
	}

	result := make([]int, len(s))
	result[0] = 1
	if s[0:2] >= "10" && s[0:2] <= "25" {
		result[1] = 2
	} else {
		result[1] = 1
	}

	for i := 1; i < len(s); i++ {
		result[i] = result[i-1]
		if s[i-1:i+1] >= "10" && s[i-1:i+1] <= "25" {
			if i-2 >= 0 {
				result[i] += result[i-2]
			}
		}
	}
	return result[len(s)-1]
}
