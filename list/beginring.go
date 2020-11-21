package list

func beginRing() *LinkedList {
	head := NewlinkedList()

	cur1 := head
	cur2 := head
	if head == nil {
		return nil
	}

	for {
		if cur2.Next == nil {
			return nil
		}

		if cur2.Next.Next == nil {
			return nil
		}

		cur1 = cur1.Next
		cur2 = cur2.Next

		if cur1 == cur2 {
			return cur1
		}
	}
}
