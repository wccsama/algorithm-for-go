package algortm

import "strconv"

func findNthDigit(n int) int {
	dig, by, count := 1, 1, 9
	for count < n {
		n -= count
		dig *= 10
		by += 1
		count = dig * 9 * by

	}
	var num = dig + (n-1)/by
	var res = strconv.Itoa(num)[(n-1)%by] - '0'

	return int(res)
}
