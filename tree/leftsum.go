package tree

func leftSum(root *Tree) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.value
	}
	return sum + leftSum(root.Left) + leftSum(root.Right)
}
