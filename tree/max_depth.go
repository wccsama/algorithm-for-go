package tree

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return treeMax(root.Left) + treeMax(root.Right)
}

var maxN int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	help(root)
	return maxN
}

func help(root *TreeNode) {
	if root == nil {
		return 0
	}
	left := help(root.Left)
	right := help(root.Right)
	maxN = max(max, left+right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

var head, pre *Tree

func generate(root *Tree) *Tree {
	if root == nil {
		return nil
	}

	in(root)
	head.Left = pre
	pre.Right = head
	return head
}

func in(root *Tree) {
	if root == nil {
		return
	}
	in(root.Left)
	if pre == nil {
		head = pre
	} else {
		pre.Right = root
	}
	root.Left = pre
	pre = root
	in(root.Right)
}
