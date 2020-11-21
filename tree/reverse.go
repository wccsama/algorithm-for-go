package tree

// 反转、镜像
func reverse(root *Tree) *Tree {
	if root == nil {
		return nil
	}

	temp = root.Right
	root.Right = reverse(root.Left)
	root.Left = reverse(temp)
	return root
}
