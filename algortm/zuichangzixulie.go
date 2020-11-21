package algortm

func findLengthOfLCIS(nums []int) int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				result[i] = max(result[i], result[j]+1)
			}
		}
	}

	max := 0
	for i := 0; i < len(nums); i++ {
		if max < result[i] {
			max = result[i]
		}
	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func less(a, b string) {

}

var res int

func dp(n, k int) int {
	if n == 1 {
		return 1
	}

	if k == 1 {
		return n
	}

	for i := 0; i < n; i++ {
		res = min(res, max(dp(n-i, k), dp(i-1, k-1))+1)
	}

	return res
}

func rob(a []int, start int) int{
	if start == len(a) {
		return 0
	}

	return max(rob(a, start+1), nums[]start + rob(a, start+2))
}

func rob(root *Tree) int {
	if root == nil {
		return 0
	}

	do := root.Val
	if root.Left != nil {
		 do += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
	     do += rob(root.Right.Left) + rob(root.Right.Right)
	}

	notdo := ob(root.Right) + rob(root.Left)
	return max(do, notdo)
}
