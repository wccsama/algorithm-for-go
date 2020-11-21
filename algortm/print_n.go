package algortm

import "fmt"

// n个数字的全排列，从1到n
// n个数字的全排列，从1到n
func printN(n int) {
	if n < 0 {
		return
	}

	a := make([]int, n)
	//for i := 0; i < 10; i++ {
	//a[0] = i
	recurN(0, n, a)
	//}
}

func recurN(index, n int, a []int) {
	if index == n {
		print2(a)
		return
	}

	for i := 0; i < 10; i++ {
		a[index] = i
		recurN(index+1, n, a)
	}
}

func print(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

// abc 三个字符的全排列
func printAll(a []int) {

}

func printAll2(a []int, n, index int) {
	if index == len(a) {
		// print
		return
	}
	for i := index + 1; i < len(a); i++ {
		a[i], a[index] = a[index], a[i]
		printAll2(a, n, index+1)
		a[i], a[index] = a[index], a[i]
	}
}
