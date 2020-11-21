package tree

func isSame(root1, root2 *Tree) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.value == root2.value {
		return true
	}

	if isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right) {
		return true
	}

	return false
}
