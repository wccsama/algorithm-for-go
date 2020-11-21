package tree

func searchAnce(root, t1, t2 *Tree) {
	if root == nil || root == t1 || root == t2 {
		return root
	}

	if root.value > t1 && root.value > t2 {
		return searchAnce(root.Left, t1, t2)
	}

	if root.value < t1 && root.value < t2 {
		return searchAnce(root.Right, t1, t2)
	}

	return root
}
