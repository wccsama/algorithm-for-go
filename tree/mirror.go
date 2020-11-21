package tree

func mirror(root *Tree) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	mirror(root.Left)
	mirror(root.Right)
}

// 前序遍历
