package list

// 有序链表中删除重复的元素,重复元素不保留
func deleteDup() *LinkedList {
	head := NewlinkedList()

	if head == nil {
		return nil
	}

	// head 减去重复的
	var pre *LinkedList

	for head != nil && head.Next != nil {
		if head.Value == head.Next.Value {
			for head.Value == head.Next.Value {
				head = head.Next
			}
			pre.Next = head

		} else {
			pre = head
			head = head.Next
		}
	}

	return head

}
