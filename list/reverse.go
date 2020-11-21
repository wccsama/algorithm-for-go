package list

// 反转链表
func reverse() *LinkedList {
	head := NewlinkedList()
	// 为空
	if head == nil {
		return nil
	}
	// 只有一个
	if head.Next == nil {
		return head
	}

	var pre *LinkedList
	current := head

	for {
		if current == nil {
			break
		}

		// 思想每次只更新一个指向
		nextTemp := current.Next
		current.Next = pre
		pre = current
		current = nextTemp
	}

	return pre
}
