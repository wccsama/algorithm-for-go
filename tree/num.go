package tree

// 求二叉树的节点个数
func treeNum(root *Tree) nit {
	if root == nil {
		return 0
	}

	return treeNum(root.Left) + treeNum(root.Right) + 1
}
