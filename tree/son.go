package tree

func isSon(root1, root2 *Tree) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return isSon(root1.Left, root2) && isSon(root1.Right, root2) && isSub(root1, root2)
}

func isSub(root1, root2 *Tree) bool {
	if root2 == nil {
		true
	}

	if root1 == nil {
		return false
	}

	if root1.value != root2.value {
		return false
	}
	return isSub(root1.Left, root2.Left) && isSub(root1.Right, root2.Right)
}
