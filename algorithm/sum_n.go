package algortm

// 输入两个整数n和sum，从数列1，2，3.......n 中随意取几个数，使其和等于sum，要求将其中所有的可能组合列出来。
// 注意到取n，和不取n个区别即可，考虑是否取第n个数的策略，可以转化为一个只和前n-1个数相关的问题。
// 如果取第n个数，那么问题就转化为“取前n-1个数使得它们的和为sum-n”，对应的代码语句就是sumOfkNumber(sum - n, n - 1)；
// 如果不取第n个数，那么问题就转化为“取前n-1个数使得他们的和为sum”，对应的代码语句为sumOfkNumber(sum, n - 1)。
var a := make([]int, 0)
func sumN(sum int , n int) {
	if sum <= 0 || n <= 0 {
		return
	}

	if sum == n {
		// print
		fmt.Print(a)
	}

	a = append(a, n)
	sumN(sum - n , n - 1)
	a = a[1:]
	sumN(sum, n - 1)
}
