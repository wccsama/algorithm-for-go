package tree

func getLevelNum(root *Tree, k int) int {
	if root == nil || k <= 0 {
		return 0
	}

	if root != nil && k == 1 {
		return 1
	}

	return getLevelNum(root.Left, k-1) + getLevelNum(root.Right, k-1)
}
