package algortm

func countDigitOne(n int) int {
	cur, hight, low, digit, res := n%10, n/10, 0, 1, 0
	for cur != 0 || hight != 0 {
		if cur == 0 {
			res += hight * digit
		} else if cur == 1 {
			res += hight*digit + 1 + low
		} else {
			res += (hight + 1) * digit
		}
		low += cur * digit
		cur = hight % 10
		hight /= 10
		digit *= 10
	}
	return res
}
