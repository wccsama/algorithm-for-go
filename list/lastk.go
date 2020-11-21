package list

// 双指针法
func lastK(k int) int {
	head := NewlinkedList()
	if k == 0 {
		return head.Value
	}

	headTwo := NewlinkedList()

	for {
		if k == 0 {
			break
		}
		if head.Next == nil {
			return -1
		}
		headTwo = headTwo.Next
		k--
	}

	for {
		if headTwo.Next == nil {
			return head.Value
		}

		headTwo = headTwo.Next
		head = head.Next
	}

	return -1
}
