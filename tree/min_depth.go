package tree

func treeMin(root *Tree) int {
	if root == nil {
		return 0
	}

	left := treeMin(root.Left)
	right := treeMin(root.Right)

	if left == 0 || right == 0 {
		return left + right + 1
	}

	if left < right {
		return left + 1
	}
	return right + 1

}
