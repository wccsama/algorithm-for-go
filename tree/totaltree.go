package tree

// 完全二叉树
func totalTree(root *Tree) {
	if root == nil {
		return
	}
	queue := Queue{}
	queue.enqueue(root)

	flag = false
	for queue.len() != 0 {
		root := queue.dequeue()
		if root != nil {
			if flag {
				return false
			}
			queue.enqueue(root.Left)
			queue.enqueue(root.Right)
		} else {
			flag = true
		}

	}

	return true
}
