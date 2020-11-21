package tree

func zhi(root *Tree) {
	queue := Queue{}
	if root == nil {
		return
	}
	queue.enqueue(root)
	flag := true
	for queue.len() != 0 {
		temp := make([]*Tree, 0)
		for i := 0; i < queue.len(); i++ {
			root := queue.dequeue()
			temp = append(temp, root)
			if root.Left != nil {
				queue.enqueue(root.Left)
			}
			if root.Right != nil {
				queue.enqueue(root.Right)
			}
		}
		if flag {

		} else {

		}

	}
}
