package tree

func duichen(root *Tree) bool {
	if root == nil {
		return true
	}
	return isSame(root.Left, root.Right)
}
