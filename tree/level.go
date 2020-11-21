package tree

func level(root *Tree) {
	if root == nil {
		return
	}
	queue := Queue{}
	queue.enqueue(root)

	for queue.len() != 0 {
		root := queue.dequeue()
		// print
		if root.Left != nil {
			queue.enqueue(root.Left)
		}
		if root.Right != nil {
			queue.enqueue(root.Right)
		}
	}

}
