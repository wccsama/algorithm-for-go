package algortm

var result [][]int

// 子集
func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	tmp := make([]int, 0)
	help(nums, 0, tmp)
	return result
}

func help(nums []int, index int, tmp []int) {
	t := make([]int, len(tmp))
	copy(t, tmp)
	result = append(result, t)
	for i := index; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		help(nums, i+1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}

var result [][]int

// 组合
func combine(n int, k int) [][]int {
	result = make([][]int, 0)
	tmp := make([]int, 0)
	helpCombine(n, 1, k, tmp)
	return result
}

func helpCombine(n, index, k int, tmp []int) {
	if len(tmp) == k {
		t := make([]int, len(tmp))
		copy(t, tmp)
		result = append(result, t)
	}
	for i := index; i <= n; i++ {
		tmp = append(tmp, i)
		helpCombine(n, i+1, k, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}

// var result [][]int

// func permute(nums []int) [][]int {
// 	result = make([][]int, 0)
// 	pro(nums, 0)
// 	return result
// }

// func pro(s []int, cur int) {
// 	if cur == len(s) {
// 		tmp := make([]int, len(s))
// 		copy(tmp, s)
// 		result = append(result, tmp)
// 		return
// 	}

// 	for i := cur; i < len(s); i++ {
// 		s[cur], s[i] = s[i], s[cur]
// 		pro(s, cur+1)
// 		s[i], s[cur] = s[cur], s[i]
// 	}
// }

// var data = []byte{'0','1','2','3','4','5','6','7','8','9'}
// var result []byte
// var final []int
// func printNumbers(n int) []int {
//     result = make([]byte, n)
//     final = make([]int, 0)
//     for i := 0; i < 10; i++ {
//         result[0] = data[i]
//         help(n, 0)
//     }
//     return final
// }

// func help(n, index int) {
//     if n  == index + 1 {
//         num := remove0(result)
//         if num != 0 {
//            final = append(final,num )
//         }
//         return
//     }

//     for i := 0; i < len(data); i++ {
//         result[index + 1] = data[i]
//         help(n, index + 1)
//     }
// }

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2

	dp := make([]bool, sum+1)
	dp[0] = true

	for i := 1; i < sum+1; i++ {
		for j := len(nums); j >= 1; j-- {
			if i >= nums[j-1] {
				dp[i] = dp[i] || dp[i-nums[j-1]]
			}
		}
	}
	return dp[sum]
}

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2

	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}

	for i := 1; i < len(nums)+1; i++ {
		for j := 1; j < sum+1; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][sum]
}
func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	count := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[i] == dp[j]+1 {
					count[i] += count[j]
				}
			}
			res = max(res, dp[i])
		}
	}

	resd := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == res {
			resd += count[i]
		}
	}
	return resd
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func smallestK(arr []int, k int) []int {
	result := make([]int, k)
	Tree(arr)
	for i := 0; i < k; i++ {
		result[i] = arr[0]
		arr[0] = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		duiDown(arr, 0)
	}
	return result
}

func Tree(nums []int) {
	for i := (len(nums) - 1) / 2; i >= 0; i-- {
		duiDown(nums, i)
	}
}

func duiDown(nums []int, a, n int) {
	j := 2*a + 1
	for j < n {
		if j+1 < n && nums[j] > nums[j+1] {
			j++
		}
		if nums[a] < nums[j] {
			break
		}
		nums[a], nums[j] = nums[j], nums[a]
		a = j
		j = 2*a + 1
	}
}

func duiUp(nums int[], a int) {
	i := (a-1) / 2
	for i >= 0 {
		if nums[i] > nums[a] {
			nums[a], nums[i] = nums[i], nums[a]
		}
		a = i
		i = (a-1) / 2
	}
}

